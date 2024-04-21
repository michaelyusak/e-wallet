package helper

import (
	"errors"
	"os"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

type HashHelperIntf interface {
	HashPassword(pwd string) ([]byte, error)
	CheckPassword(pwd string, hash []byte) (bool, error)
}

type HashHelperImpl struct{}

func (h *HashHelperImpl) HashPassword(pwd string) ([]byte, error) {
	costStr := os.Getenv("HASH_COST")

	cost, err := strconv.Atoi(costStr)
	if err != nil {
		return nil, errors.New("error while reading env")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), cost)
	if err != nil {
		return nil, err
	}
	return hash, nil
}

func (h *HashHelperImpl) CheckPassword(pwd string, hash []byte) (bool, error) {
	err := bcrypt.CompareHashAndPassword(hash, []byte(pwd))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
