package helper

import (
	"fmt"
	"os"
	"time"

	"e-wallet/constants"

	"github.com/golang-jwt/jwt/v5"
)

type TokenHelperIntf interface {
	CreateAndSign(userId int, purpose constants.Purpose) (string, error)
}

type TokenHelperImpl struct{}

type vars struct {
	secretKey string
	issuer    string
}

func (h *TokenHelperImpl) CreateAndSign(userId int, purpose constants.Purpose) (string, error) {
	vars := getVars()

	duration := constants.DefaultDuration

	if purpose == constants.ResetTokenPurpose {
		duration = constants.ResetPasswordDuration
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  userId,
		"iat": time.Now(),
		"iss": vars.issuer,
		"exp": time.Now().Add(duration).Unix(),
	})
	signed, err := token.SignedString([]byte(vars.secretKey))
	if err != nil {
		return "", err
	}
	return signed, nil
}

func ParseAndVerify(signed string) (jwt.MapClaims, error) {
	vars := getVars()

	token, err := jwt.Parse(signed, func(token *jwt.Token) (interface{}, error) {
		return []byte(vars.secretKey), nil
	}, jwt.WithIssuer(vars.issuer),
		jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Name}),
		jwt.WithExpirationRequired(),
	)
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return claims, nil
	} else {
		return nil, fmt.Errorf("unknown claims")
	}
}

func getVars() vars {
	secretKey := os.Getenv("JWT_SECRET_KEY")
	issuer := os.Getenv("ISSUER")

	return vars{
		secretKey: secretKey,
		issuer:    issuer,
	}
}
