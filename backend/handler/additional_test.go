package handler_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"e-wallet/dto"
	"e-wallet/server"

	"github.com/go-playground/assert/v2"
)

func TestHandler_NoRouteHandlerFunc(t *testing.T) {
	var (
		path = "/"

		expectedRes = dto.ErrResponse{
			Message: "page not found",
			Code:    http.StatusNotFound,
		}
	)

	t.Run("should return message response", func(t *testing.T) {
		// given
		router := server.Setup(server.RouterOpt{})
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, path, nil)
		expectedResBytes, _ := json.Marshal(expectedRes)

		// when
		router.ServeHTTP(w, req)

		// then
		assert.Equal(t, string(expectedResBytes), w.Body.String())
	})
}
