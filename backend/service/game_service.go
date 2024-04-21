package service

import (
	"context"
	"e-wallet/apperror"
	"e-wallet/constants"
	"e-wallet/entity"
	"e-wallet/helper"
	"e-wallet/repository"
	"errors"
	"fmt"
	"strconv"

	"github.com/shopspring/decimal"
)

var (
	numToDecreaseGachaTrial = decimal.NewFromInt(-1)
)

type GameService interface {
	AttemptGacha(ctx context.Context, selectionStr string) (*entity.GachaBox, error)
}

type gameServiceIpl struct {
	prizeRepository  repository.PrizeRepository
	gachaRepository  repository.GachaRepository
	walletRepository repository.WalletRepository
}

func NewGameServiceIpl(prizeRepository repository.PrizeRepository, gachaRepository repository.GachaRepository, walletRepository repository.WalletRepository) gameServiceIpl {
	return gameServiceIpl{
		prizeRepository:  prizeRepository,
		gachaRepository:  gachaRepository,
		walletRepository: walletRepository,
	}
}

var (
	companyWalletId = constants.CompanyWalletId
)

func (s *gameServiceIpl) AttemptGacha(ctx context.Context, selectionStr string) (*entity.GachaBox, error) {
	userId := ctx.Value(constants.UserId).(int)
	if userId == 0 {
		return nil, apperror.StatusUnauthorized()
	}

	wallet, err := s.walletRepository.GetWalletByUserId(ctx, userId)
	if err != nil {
		return nil, apperror.InternalServerErr("internal error while checking user wallet")
	}
	if wallet == nil {
		return nil, apperror.StatusWalletNotFound()
	}

	if wallet.GachaTrial < 1 {
		return nil, apperror.StatusGachaTrialInsufficient()
	}

	selection, err := helper.CheckForGacha(selectionStr)
	if err != nil {
		return nil, err
	}

	gachaBoxes, err := s.prizeRepository.GetAllBox(ctx)
	if err != nil {
		return nil, apperror.InternalServerErr("internal error while fetching boxes")
	}

	selectedBox := helper.SelectBox(gachaBoxes, selection)

	if !selectedBox.Prize.IsZero() {
		errWithTx := s.walletRepository.WithTx(ctx, func(repo repository.WalletRepository) error {
			err = repo.UpdateBalance(ctx, wallet.Id, selectedBox.Prize)
			if err != nil {
				return apperror.InternalServerErr("internal error while updating user balance")
			}

			topupReq := entity.Transaction{
				RecepientWalletNum: wallet.Number,
				SenderWalletNum: fmt.Sprintf("100000000000%s", strconv.Itoa(companyWalletId)),
				SourceOfFunds: constants.GachaStr,
				Description:   "Top Up From Reward",
				Amount:        selectedBox.Prize,
			}

			err = repo.PostOneTransaction(ctx, topupReq)
			if err != nil {
				return apperror.InternalServerErr("internal error while posting transaction")
			}

			return nil
		})

		if errWithTx != nil {
			if errors.As(errWithTx, &apperror.AppError{}) {
				return nil, errWithTx
			}

			return nil, apperror.InternalServerErr("internal error while executing transaction")
		}
	}

	err = s.gachaRepository.PostOneUserGacha(ctx, selectedBox.PrizeId, wallet.Id)
	if err != nil {
		return nil, apperror.InternalServerErr("internal error while posting user result")
	}

	err = s.walletRepository.UpdateGachaTrial(ctx, userId, numToDecreaseGachaTrial)
	if err != nil {
		return nil, apperror.InternalServerErr("internal error while updating user wallet")
	}

	return &selectedBox, nil
}
