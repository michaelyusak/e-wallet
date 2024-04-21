package service

import (
	"context"
	"e-wallet/apperror"
	"e-wallet/constants"
	"e-wallet/entity"
	"e-wallet/helper"
	"e-wallet/repository"
	"time"
)

type RPTService interface {
	ResetPassword(ctx context.Context, resetPwdReq entity.ResetPasswordToken) error
	RequestToken(ctx context.Context, rptReq entity.ResetPasswordToken) (*entity.ResetPasswordToken, error)
}

type rptServiceIpl struct {
	rptRepository  repository.RPTRepository
	userRepository repository.UserRepository
	tokenHelper    helper.TokenHelperIntf
	hashHelper     helper.HashHelperIntf
}

func NewRPTService(rptRepository repository.RPTRepository, userRepository repository.UserRepository, tokenHelper helper.TokenHelperIntf, hashHelper helper.HashHelperIntf) rptServiceIpl {
	return rptServiceIpl{
		rptRepository:  rptRepository,
		userRepository: userRepository,
		tokenHelper:    tokenHelper,
		hashHelper:     hashHelper,
	}
}

func (s *rptServiceIpl) RequestToken(ctx context.Context, rptReq entity.ResetPasswordToken) (*entity.ResetPasswordToken, error) {
	err := helper.CheckEmail(rptReq.Email)
	if err != nil {
		return nil, apperror.StatusInvalidEmail()
	}

	user, err := s.userRepository.GetUserByEmail(ctx, rptReq.Email)
	if err != nil {
		return nil, apperror.InternalServerErr("internal error while checking user email")
	}
	if user == nil {
		return nil, apperror.StatusAccountUnregistered()
	}

	err = s.rptRepository.DeleteExisting(ctx, user.Id)
	if err != nil {
		return nil, apperror.InternalServerErr("internal error while invalidating user previous token")
	}

	token, err := s.tokenHelper.CreateAndSign(user.Id, constants.ResetTokenPurpose)
	if err != nil {
		return nil, apperror.InternalServerErr("internal error while generating code")
	}

	rptReq.Token = token
	rptReq.ExpiredAt = time.Now().Add(constants.ResetPasswordDuration)

	err = s.rptRepository.PostOneRPT(ctx, rptReq, user.Id)
	if err != nil {
		return nil, apperror.InternalServerErr("internal error while storing user token")
	}

	return &rptReq, nil
}

func (s *rptServiceIpl) ResetPassword(ctx context.Context, resetPwdReq entity.ResetPasswordToken) error {
	user, err := s.userRepository.GetUserByEmail(ctx, resetPwdReq.Email)
	if err != nil {
		return apperror.InternalServerErr("internal error while checking user email")
	}
	if user == nil {
		return apperror.StatusAccountUnregistered()
	}

	rpt, err := s.rptRepository.GetTokenExpiredAt(ctx, resetPwdReq.Token)
	if err != nil {
		return apperror.InternalServerErr("internal error while validating user token")
	}
	if rpt == nil {
		return apperror.StatusTokenInvalid()
	}
	if rpt.ExpiredAt.Before(time.Now()) {
		return apperror.StatusTokenInvalid()
	}

	claims, err := helper.ParseAndVerify(resetPwdReq.Token)
	if err != nil {
		return apperror.InternalServerErr("internal error while parsing user token")
	}

	if int(claims["id"].(float64)) != user.Id {
		return apperror.StatusUnauthorized()
	}

	err = helper.CheckPassword(resetPwdReq.NewPassword)
	if err != nil {
		return err
	}

	hashPwd, err := s.hashHelper.HashPassword(resetPwdReq.NewPassword)
	if err != nil {
		return apperror.InternalServerErr("internal error while preparing user new password")
	}

	err = s.userRepository.ResetPassword(ctx, user.Id, hashPwd)
	if err != nil {
		return apperror.InternalServerErr("internal error while updating user password")
	}

	err = s.rptRepository.DeleteExisting(ctx, user.Id)
	if err != nil {
		return apperror.InternalServerErr("internal error while invalidation user token")
	}

	return nil
}
