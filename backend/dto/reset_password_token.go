package dto

import (
	"e-wallet/entity"
	"time"
)

type ResetPasswordTokenDTO struct {
	Email       string `json:"email,omitempty" binding:"required"`
	Token       string `json:"token"`
	ExpiredAt   string `json:"expired_at"`
	NewPassword string `json:"new_password,omitempty"`
}

func ToResetPasswordToken(rptDTO ResetPasswordTokenDTO) entity.ResetPasswordToken {
	return entity.ResetPasswordToken{
		Email:       rptDTO.Email,
		Token:       rptDTO.Token,
		NewPassword: rptDTO.NewPassword,
	}
}

func ToResetPasswordTokenDTO(rpt entity.ResetPasswordToken) ResetPasswordTokenDTO {
	return ResetPasswordTokenDTO{
		Email:     rpt.Email,
		Token:     rpt.Token,
		ExpiredAt: rpt.ExpiredAt.Format(time.RFC3339),
	}
}
