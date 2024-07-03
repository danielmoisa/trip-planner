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

func TestGetTrips(t *testing.T) {
	t.Parallel()
	test.WithTestServer(t, func(s *api.Server) {
		fixtures := test.Fixtures()
		res := test.PerformRequest(t, s, "GET", "/api/v1/trips?per_page=10&page=1", nil, test.HeadersWithAuth(t, fixtures.User1AccessToken1.Token))
		require.Equal(t, http.StatusOK, res.Result().StatusCode)

		var response types.GetTripsResponse
		test.ParseResponseAndValidate(t, res, &response)

		assert.NotEmpty(t, response)
		assert.LessOrEqual(t, len(response), 10)

		test.Snapshoter.Save(t, response)

	})
}

func TestGetTripsWithFilter(t *testing.T) {
	t.Parallel()
	test.WithTestServer(t, func(s *api.Server) {
		fixtures := test.Fixtures()
		res := test.PerformRequest(t, s, "GET", "/api/v1/trips?per_page=10&page=1&trip_name=Winter%Getaway", nil, test.HeadersWithAuth(t, fixtures.User1AccessToken1.Token))
		require.Equal(t, http.StatusOK, res.Result().StatusCode)

		var response types.GetTripsResponse
		test.ParseResponseAndValidate(t, res, &response)

		assert.NotEmpty(t, response)
		for _, trip := range response {
			assert.Contains(t, trip.Name, "Test")
		}
		test.Snapshoter.Save(t, response)
	})
}

func TestGetTripsWithInvalidParams(t *testing.T) {
	t.Parallel()
	test.WithTestServer(t, func(s *api.Server) {
		fixtures := test.Fixtures()
		res := test.PerformRequest(t, s, "GET", "/api/v1/trips?per_page=-1&page=0", nil, test.HeadersWithAuth(t, fixtures.User1AccessToken1.Token))
		require.Equal(t, http.StatusBadRequest, res.Result().StatusCode)

		var response types.PublicHTTPError
		test.ParseResponseAndValidate(t, res, &response)
		test.Snapshoter.Save(t, response)
	})
}

func TestGetTripsWhileDBInaccessible(t *testing.T) {
	t.Parallel()
	test.WithTestServer(t, func(s *api.Server) {
		fixtures := test.Fixtures()
		s.DB.Close() // forcefully close the database!
		res := test.PerformRequest(t, s, "GET", "/api/v1/trips?per_page=10&page=1", nil, test.HeadersWithAuth(t, fixtures.User1AccessToken1.Token))
		require.Equal(t, http.StatusInternalServerError, res.Result().StatusCode)

		var response types.PublicHTTPError
		test.ParseResponseAndValidate(t, res, &response)
		test.Snapshoter.Save(t, response)
	})
}
