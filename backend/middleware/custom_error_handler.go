package middleware

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"e-wallet/apperror"
	"e-wallet/dto"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func CustomMiddlewareError(ctx *gin.Context) {
	ctx.Next()

	if len(ctx.Errors) > 0 {
		firstError := ctx.Errors[0].Err
		errResponse := CheckError(firstError)
		ctx.AbortWithStatusJSON(errResponse.Code, errResponse)
	}
}

func CheckError(err error) dto.ErrResponse {
	var ve validator.ValidationErrors

	var appErr apperror.AppError

	var unmarshalErr json.UnmarshalTypeError
	unmarchalErrType := &unmarshalErr

	var jse json.SyntaxError
	jseErrType := &jse

	if errors.As(err, &ve) {
		details := GenerateValidationErrs(ve)
		return dto.ErrResponse{Code: http.StatusBadRequest, Message: "validation error", Details: details}

	} else if errors.As(err, &appErr) {
		return dto.ErrResponse{Code: appErr.StatusCode, Message: appErr.Message}

	} else if errors.As(err, &unmarchalErrType) {
		return dto.ErrResponse{Code: http.StatusBadRequest, Message: "unmarshal error"}

	} else if errors.As(err, &jseErrType) {
		return dto.ErrResponse{Code: http.StatusBadRequest, Message: "json syntax error"}
	}

	return dto.ErrResponse{Code: http.StatusInternalServerError, Message: "internal error"}
}

func GenerateValidationErrs(ve validator.ValidationErrors) []dto.ValidationErrorMessage {
	details := make([]dto.ValidationErrorMessage, len(ve))

	for i, fe := range ve {
		details[i] = dto.ValidationErrorMessage{Field: fe.Field(), Message: getErrorMsg(fe)}
	}

	return details
}

func getErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "this field is required"

	case "max":
		return fmt.Sprintf("should be less than %s character", fe.Param())
	}

	return "unknown error"
}
