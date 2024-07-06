package trips_test

import (
	"net/http"
	"testing"

	"github.com/danielmoisa/trip-planner/internal/api"
	"github.com/danielmoisa/trip-planner/internal/test"
	"github.com/danielmoisa/trip-planner/internal/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetRandomTrip(t *testing.T) {
	t.Parallel()

	test.WithTestServer(t, func(s *api.Server) {

		fixtures := test.Fixtures()

		res := test.PerformRequest(t, s, "GET", "/api/v1/trips/random", nil, test.HeadersWithAuth(t, fixtures.User1AccessToken1.Token))
		require.Equal(t, http.StatusOK, res.Result().StatusCode)

		var response types.GetTripsResponse
		test.ParseResponseAndValidate(t, res, &response)

		assert.Len(t, response, 1)

	})
}
