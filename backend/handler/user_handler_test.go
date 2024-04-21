package handler_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"e-wallet/apperror"
	"e-wallet/constants"
	"e-wallet/dto"
	"e-wallet/entity"
	"e-wallet/handler"
	"e-wallet/mocks"
	"e-wallet/server"

	"github.com/go-playground/assert/v2"
)

func TestHandler_RegisterUser(t *testing.T) {
	var (
		invalidReq = "{\"abc\":\"123\",}"

		path = "/users"

		userReq = dto.UserDTO{
			Name:     "abc",
			Email:    "abc",
			Password: "abc",
		}

		user = entity.User{
			Name:     "abc",
			Email:    "abc",
			Password: "abc",
		}

		errTest = apperror.InternalServerErr("test error")

		ctx = context.Background()
	)

	t.Run("should return error response when req body is invalid", func(t *testing.T) {
		// given
		mockUserService := new(mocks.UserService)
		userHandler := handler.NewUserHandler(mockUserService)
		router := server.Setup(server.RouterOpt{UserHandler: userHandler})
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, path, strings.NewReader(invalidReq))
		expectedErrRes := dto.ErrResponse{
			Message: "json syntax error",
			Code:    http.StatusBadRequest,
		}
		expectedErrResBytes, _ := json.Marshal(expectedErrRes)

		// when
		router.ServeHTTP(w, req)

		// then
		assert.Equal(t, string(expectedErrResBytes), w.Body.String())
	})
	t.Run("should return error response when service layer return an error", func(t *testing.T) {
		// given
		mockUserService := new(mocks.UserService)
		mockUserService.On("RegisterUser", ctx, user).Return(nil, errTest)
		userHandler := handler.NewUserHandler(mockUserService)
		router := server.Setup(server.RouterOpt{UserHandler: userHandler})
		w := httptest.NewRecorder()
		reqBodyBytes, _ := json.Marshal(userReq)
		req, _ := http.NewRequest(http.MethodPost, path, bytes.NewReader(reqBodyBytes))
		expectedErrRes := dto.ErrResponse{
			Message: errTest.Message,
			Code:    errTest.StatusCode,
		}
		expectedErrResBytes, _ := json.Marshal(expectedErrRes)

		// when
		router.ServeHTTP(w, req)

		// then
		assert.Equal(t, string(expectedErrResBytes), w.Body.String())
	})
	t.Run("should return message response", func(t *testing.T) {
		// given
		mockUserService := new(mocks.UserService)
		mockUserService.On("RegisterUser", ctx, user).Return(&user, nil)
		userHandler := handler.NewUserHandler(mockUserService)
		router := server.Setup(server.RouterOpt{UserHandler: userHandler})
		w := httptest.NewRecorder()
		reqBodyBytes, _ := json.Marshal(userReq)
		req, _ := http.NewRequest(http.MethodPost, path, bytes.NewReader(reqBodyBytes))
		expectedRes := dto.MessageResponse{
			Message: constants.MsgResUserCreated,
			Data:    dto.ToUserDTO(user),
		}
		expectedResBytes, _ := json.Marshal(expectedRes)

		// when
		router.ServeHTTP(w, req)

		// then
		assert.Equal(t, string(expectedResBytes), w.Body.String())
	})
}
