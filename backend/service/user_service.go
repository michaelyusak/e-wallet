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
	"io"
	"log"
	"mime/multipart"
	"os"
	"strings"

	"github.com/google/uuid"
)

type UserService interface {
	GetDetail(ctx context.Context) (*entity.User, error)
	Login(ctx context.Context, userReq entity.User) (string, error)
	RegisterUser(ctx context.Context, userReq entity.User) (*entity.User, string, error)
	UpdateProfilePicture(ctx context.Context, uploadedImage multipart.File, fileHeader *multipart.FileHeader) error
	UpdateUserData(ctx context.Context, userReq entity.User) error
}

type userServiceIpl struct {
	userRepository   repository.UserRepository
	walletRepository repository.WalletRepository
	hashHelper       helper.HashHelperIntf
	tokenHelper      helper.TokenHelperIntf
}

func NewUserServiceIpl(userRepository repository.UserRepository, walletRepository repository.WalletRepository, hashHelper helper.HashHelperIntf, tokenHelper helper.TokenHelperIntf) userServiceIpl {
	return userServiceIpl{
		userRepository:   userRepository,
		walletRepository: walletRepository,
		hashHelper:       hashHelper,
		tokenHelper:      tokenHelper,
	}
}

func (s *userServiceIpl) UpdateProfilePicture(ctx context.Context, uploadedImage multipart.File, fileHeader *multipart.FileHeader) error {
	userId := ctx.Value(constants.UserId)
	if userId == nil {
		return apperror.StatusUnauthorized()
	}

	user, err := s.userRepository.GetUserById(ctx, userId.(int))
	if err != nil {
		return apperror.InternalServerErr("internal error while fetching user data")
	}
	if user == nil {
		return apperror.StatusUnauthorized()
	}

	fileName := fileHeader.Filename
	format := strings.Split(fileName, ".")[1]

	if format != "png" && format != "jpg" {
		return apperror.BadRequest("wrong type of file")
	}

	imageName := fmt.Sprintf("%s.%s", uuid.New().String(), format)
	imagePath := fmt.Sprintf("./profile_pictures/%s", imageName)

	targetFile, err := os.OpenFile(imagePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		if !os.IsNotExist(err) {
			return apperror.InternalServerErr("internal error while retrieving uploaded image")
		}

		err := os.MkdirAll(imagePath, 0755)
		if err != nil {
			return apperror.InternalServerErr("internal error while preparing uploaded image")
		}
	}
	defer targetFile.Close()

	_, err = io.Copy(targetFile, uploadedImage)
	if err != nil {
		return apperror.InternalServerErr("internal error while saving user image")
	}

	err = s.userRepository.UpdateProfilePicture(ctx, user.Id, imageName)
	if err != nil {
		return apperror.InternalServerErr("internal error while saving user image path")
	}

	return nil
}

func (s *userServiceIpl) RegisterUser(ctx context.Context, userReq entity.User) (*entity.User, string, error) {
	err := helper.CheckRegisterRequest(userReq)
	if err != nil {
		return nil, "", err
	}

	userReq.HashPassword, err = s.hashHelper.HashPassword(userReq.Password)
	if err != nil {
		return nil, "", apperror.InternalServerErr("internal error while hashing password")
	}

	user, errWithTx := s.userRepository.WithTx(ctx, func(r repository.UserRepository) (*entity.User, error) {
		user, err := s.userRepository.GetUserByEmail(ctx, userReq.Email)
		if err != nil {
			return nil, apperror.InternalServerErr("internal error while checking your email")
		}
		if user != nil {
			return nil, apperror.StatusEmailTaken()
		}

		user, err = s.userRepository.PostOneUser(ctx, userReq)
		if err != nil {
			return nil, apperror.InternalServerErr("internal error while creating your account")
		}

		return user, nil
	})

	if errWithTx != nil {
		if errors.As(errWithTx, &apperror.AppError{}) {
			return nil, "", errWithTx
		}

		return nil, "", apperror.InternalServerErr("internal error while executing transaction")
	}

	wallet, err := s.walletRepository.PostOneWallet(ctx, user.Id)
	if err != nil {
		return nil, "", apperror.InternalServerErr("internal error while creating your wallet")
	}

	token, err := s.tokenHelper.CreateAndSign(user.Id, constants.DefaultPurpose)
	if err != nil {
		return nil, "", apperror.InternalServerErr("internal error while creating token")
	}

	user.Wallet = *wallet

	return user, token, nil
}

func (s *userServiceIpl) Login(ctx context.Context, userReq entity.User) (string, error) {
	user, err := s.userRepository.GetUserByEmail(ctx, userReq.Email)
	if err != nil {
		return "", apperror.InternalServerErr("internal error while checking user email")
	}
	if user == nil {
		return "", apperror.StatusAccountUnregistered()
	}

	isMatched, err := s.hashHelper.CheckPassword(userReq.Password, user.HashPassword)
	if err != nil {
		return "", apperror.InternalServerErr("internal error while comparing user password")
	}
	if !isMatched {
		return "", apperror.StatusUnmatchedPwd()
	}

	token, err := s.tokenHelper.CreateAndSign(user.Id, constants.DefaultPurpose)
	if err != nil {
		return "", apperror.InternalServerErr("internal error while generating token")
	}

	return token, nil
}

func (s *userServiceIpl) GetDetail(ctx context.Context) (*entity.User, error) {
	userId := ctx.Value(constants.UserId)
	if userId == nil {
		return nil, apperror.StatusUnauthorized()
	}

	user, err := s.userRepository.GetUserById(ctx, userId.(int))
	if err != nil {
		return nil, apperror.InternalServerErr("internal error while fetching user data")

	}
	if user == nil {
		return nil, apperror.StatusAccountUnregistered()
	}

	wallet, err := s.walletRepository.GetWalletByUserId(ctx, userId.(int))
	if err != nil {
		return nil, apperror.InternalServerErr("internal error while fetching user wallet")
	}
	if wallet == nil {
		log.Println(userId)
		return nil, apperror.StatusWalletNotFound()
	}

	err = s.walletRepository.GetCashFlow(ctx, wallet)
	if err != nil {
		return nil, apperror.InternalServerErr("internal error while fetching user cash flow")
	}

	user.Wallet = *wallet

	return user, nil
}

func (s *userServiceIpl) UpdateUserData(ctx context.Context, userReq entity.User) error {
	userId := ctx.Value(constants.UserId).(int)
	if userId == 0 {
		return apperror.StatusUnauthorized()
	}

	userReq.Id = userId

	_, errWithTx := s.userRepository.WithTx(ctx, func(repo repository.UserRepository) (*entity.User, error) {
		err := repo.UpdateUserEmail(ctx, userReq)
		if err != nil {
			return nil, apperror.InternalServerErr("error while updating user email")
		}

		err = repo.UpdateUserName(ctx, userReq)
		if err != nil {
			return nil, apperror.InternalServerErr("error while updating user name")
		}

		return nil, nil
	})

	if errWithTx != nil {
		return apperror.InternalServerErr("error while updating user data")
	}

	return nil
}
