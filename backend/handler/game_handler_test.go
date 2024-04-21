package handler_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"e-wallet/apperror"
	"e-wallet/constants"
	"e-wallet/dto"
	"e-wallet/handler"
	"e-wallet/mocks"
	"e-wallet/server"

	"github.com/go-playground/assert/v2"
	"github.com/joho/godotenv"
)

func TestHandler_AttemptGacha(t *testing.T) {
	var (
		selectionStr = "2"

		path = fmt.Sprintf("/games/gacha?selectio=%s", selectionStr)

		errTest = apperror.InternalServerErr("test error")

		ctx = context.WithValue(context.Background(), constants.UserId, 4)
	)

	t.Run("should return error response", func(t *testing.T) {
		// given
		mockGameService := new(mocks.GameService)
		gameHandler := handler.NewGameHandler(mockGameService)
		router := server.Setup(server.RouterOpt{GameHandler: gameHandler})
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, path, nil)
		expectedErr := apperror.StatusUnauthorized()
		expectedErrRes := dto.ErrResponse{
			Message: expectedErr.Message,
			Code:    expectedErr.StatusCode,
		}
		expectedErrResBytes, _ := json.Marshal(expectedErrRes)

		// when
		router.ServeHTTP(w, req)

		// then
		assert.Equal(t, string(expectedErrResBytes), w.Body.String())
	})
	t.Run("should return error response", func(t *testing.T) {
		// given
		godotenv.Load(".test.env")
		testToken := os.Getenv("TEST_TOKEN")
		mockGameService := new(mocks.GameService)
		mockGameService.On("AttemptGacha", ctx, selectionStr).Return(nil, errTest)
		gameHandler := handler.NewGameHandler(mockGameService)
		router := server.Setup(server.RouterOpt{GameHandler: gameHandler})
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, path, nil)
		w.Header().Add("Authorization", testToken)
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
}
