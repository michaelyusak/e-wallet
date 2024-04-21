package helper

import (
	"strings"

	"e-wallet/apperror"
	"e-wallet/constants"
	"e-wallet/entity"
)

func CheckRegisterRequest(userReq entity.User) error {
	err := CheckEmail(userReq.Email)
	if err != nil {
		return err
	}

	if userReq.Name == "" {
		return apperror.StatusNameEmpty()
	}

	err = CheckPassword(userReq.Password)
	if err != nil {
		return err
	}

	return nil
}

func CheckEmail(email string) error {
	if !strings.Contains(email, "@") {
		return apperror.StatusInvalidEmail()
	}

	emailSplit := strings.Split(email, "@")
	if len(emailSplit) != 2 {
		return apperror.StatusInvalidEmail()
	}

	prefix := emailSplit[0]
	domain := emailSplit[1]

	if prefix != "" && domain != "" {
		if !strings.Contains(domain, ".") {
			return apperror.StatusInvalidEmail()
		}

		domainSplit := strings.Split(domain, ".")
		if domainSplit[0] != "" && domainSplit[1] != "" {
			return nil
		}
	}

	return apperror.StatusInvalidEmail()
}

func CheckPassword(password string) error {
	if len(password) < constants.MinPasswordLen {
		return apperror.StatusPasswordInvalid()
	}

	containsNum := false
	containsAlphaLow := false
	containsAlphaUpp := false

	for _, char := range password {
		if char >= 48 && char <= 57 {
			containsNum = true
		}
		if char >= 97 && char <= 122 {
			containsAlphaLow = true
		}
		if char >= 65 && char <= 90 {
			containsAlphaUpp = true
		}
	}

	if !containsNum || !containsAlphaLow || !containsAlphaUpp {
		return apperror.StatusPasswordInvalid()
	}

	return nil
}
