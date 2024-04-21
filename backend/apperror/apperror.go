package apperror

import (
	"fmt"
	"net/http"

	"github.com/shopspring/decimal"
)

type AppError struct {
	StatusCode int
	Message    string
}

func (e AppError) Error() string {
	return fmt.Sprintf(e.Message)
}

func InternalServerErr(message string) AppError {
	return AppError{
		StatusCode: http.StatusInternalServerError,
		Message:    message,
	}
}

func BadRequest(message string) AppError {
	return AppError{
		StatusCode: http.StatusBadRequest,
		Message:    message,
	}
}

func StatusEmailTaken() AppError {
	return AppError{
		StatusCode: http.StatusConflict,
		Message:    "other user has registered using this email, please try to register with a different email",
	}
}

func StatusInvalidEmail() AppError {
	return AppError{
		StatusCode: http.StatusBadRequest,
		Message:    "your email is invalid",
	}
}

func StatusPasswordInvalid() AppError {
	return AppError{
		StatusCode: http.StatusBadRequest,
		Message:    "your password is invalid, password should at least 6 character long and contains number and lowercase & uppercase alphabet",
	}
}

func StatusNameEmpty() AppError {
	return AppError{
		StatusCode: http.StatusBadRequest,
		Message:    "field name cannot be empty",
	}
}

func StatusAccountUnregistered() AppError {
	return AppError{
		StatusCode: http.StatusUnauthorized,
		Message:    "cannot do the action because your account is unregistered",
	}
}

func StatusUnmatchedPwd() AppError {
	return AppError{
		StatusCode: http.StatusUnauthorized,
		Message:    "either your email or password is wrong, please check your request",
	}
}

func StatusAmountInvalid(min, max decimal.Decimal) AppError {
	return AppError{
		StatusCode: http.StatusBadRequest,
		Message:    fmt.Sprintf("amount must be in between %s and %s", min.String(), max.String()),
	}
}

func StatusSourceOfFundsEmpty() AppError {
	return AppError{
		StatusCode: http.StatusBadRequest,
		Message:    "please select one of the available source of funds",
	}
}

func StatusSelfTransfer() AppError {
	return AppError{
		StatusCode: http.StatusBadRequest,
		Message:    "transfer to the same wallet is not allowed",
	}
}

func StatusRecepientNotFound() AppError {
	return AppError{
		StatusCode: http.StatusConflict,
		Message:    "recepient account not found",
	}
}

func StatusUnauthorized() AppError {
	return AppError{
		StatusCode: http.StatusUnauthorized,
		Message:    "user are not unauthorized to perform this action",
	}
}

func StatusUnsufficientBalance() AppError {
	return AppError{
		StatusCode: http.StatusBadRequest,
		Message:    "user balance is unsufficient to perform this action",
	}
}

func StatusRequestUnavailable() AppError {
	return AppError{
		StatusCode: http.StatusNotImplemented,
		Message:    "user request cannot be processed in current version, wait for update",
	}
}

func StatusTokenInvalid() AppError {
	return AppError{
		StatusCode: http.StatusUnauthorized,
		Message:    "user token is unregistered or has been expired",
	}
}

func StatusWalletNotFound() AppError {
	return AppError{
		StatusCode: http.StatusConflict,
		Message:    "cannot found user wallet",
	}
}

func StatusGachaTrialInsufficient() AppError {
	return AppError{
		StatusCode: http.StatusPaymentRequired,
		Message:    "you do not have a chance",
	}
}
