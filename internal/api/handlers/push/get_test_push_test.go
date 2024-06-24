package push_test

import (
	"net/http"
	"testing"

	"github.com/danielmoisa/trip-planner/internal/api"
	"github.com/danielmoisa/trip-planner/internal/test"
	"github.com/stretchr/testify/assert"
)

func TestGetTestPush(t *testing.T) {
	test.WithTestServer(t, func(s *api.Server) {
		fixtures := test.Fixtures()

		res := test.PerformRequest(t, s, "GET", "/api/v1/push/test", nil, test.HeadersWithAuth(t, fixtures.User1AccessToken1.Token))

		assert.Equal(t, http.StatusOK, res.Result().StatusCode)
	})
}

func TestGetTestPushUnauthorized(t *testing.T) {
	test.WithTestServer(t, func(s *api.Server) {
		res := test.PerformRequest(t, s, "GET", "/api/v1/push/test", nil, nil)

		assert.Equal(t, http.StatusUnauthorized, res.Result().StatusCode)
	})
}
