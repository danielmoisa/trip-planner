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

func TestGetTrip(t *testing.T) {
	t.Parallel()

	test.WithTestServer(t, func(s *api.Server) {

		fixtures := test.Fixtures()

		res := test.PerformRequest(t, s, "GET", "/api/v1/trips/98ad176b-af90-44b7-b991-d9ebfc5dd911", nil, test.HeadersWithAuth(t, fixtures.User1AccessToken1.Token))
		require.Equal(t, http.StatusOK, res.Result().StatusCode)

		var response types.GetTripsResponse
		test.ParseResponseAndValidate(t, res, &response)

		assert.Len(t, response, 1)

		assert.Equal(t, "98ad176b-af90-44b7-b991-d9ebfc5dd911", (string(*response[0].ID)))

		test.Snapshoter.Save(t, response)
	})
}

func TestGetNonFoundTrip(t *testing.T) {

	t.Parallel()

	test.WithTestServer(t, func(s *api.Server) {
		fixtures := test.Fixtures()

		res := test.PerformRequest(t, s, "GET", "/api/v1/trips/c81618fb-3a83-40ab-935a-e3f639949e00", nil, test.HeadersWithAuth(t, fixtures.User1AccessToken1.Token))
		require.Equal(t, http.StatusNotFound, res.Result().StatusCode)

		var response types.PublicHTTPError
		test.ParseResponseAndValidate(t, res, &response)

		test.Snapshoter.Save(t, response)
	})

}

// Only here to showcase that an inaccessable database will result in a 500
// We won't actually hit the get_trips controller, as the authorization middleware will kick in before (try debugging)
func TestGetTripsWhileDBInaccessable(t *testing.T) {
	t.Parallel()

	test.WithTestServer(t, func(s *api.Server) {
		fixtures := test.Fixtures()

		s.DB.Close() // forcefully close the database!

		res := test.PerformRequest(t, s, "GET", "/api/v1/trips/98ad176b-af90-44b7-b991-d9ebfc5dd911", nil, test.HeadersWithAuth(t, fixtures.User1AccessToken1.Token))
		require.Equal(t, http.StatusInternalServerError, res.Result().StatusCode)

		var response types.PublicHTTPError
		test.ParseResponseAndValidate(t, res, &response)

		test.Snapshoter.Save(t, response)
	})
}
