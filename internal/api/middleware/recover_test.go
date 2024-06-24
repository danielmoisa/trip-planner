package middleware_test

import (
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/danielmoisa/trip-planner/internal/api"
	"github.com/danielmoisa/trip-planner/internal/api/middleware"
	"github.com/danielmoisa/trip-planner/internal/test"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"

	"github.com/stretchr/testify/require"
)

func TestLogErrorFuncWithRequestInfo(t *testing.T) {
	test.WithTestServer(t, func(s *api.Server) {
		path := "/testing-e87bc94c-2d1f-4342-9ec2-f158c63ac6da"

		s.Echo.Use(echoMiddleware.RecoverWithConfig(echoMiddleware.RecoverConfig{
			LogErrorFunc: middleware.LogErrorFuncWithRequestInfo,
		}))

		s.Echo.POST(path, func(c echo.Context) error {

			var val *int
			fmt.Println(*val)

			return c.NoContent(http.StatusNoContent)
		})

		res := test.PerformRequest(t, s, "POST", path, nil, nil)
		require.Equal(t, http.StatusInternalServerError, res.Result().StatusCode)

		body, err := io.ReadAll(res.Body)
		require.NoError(t, err)

		test.Snapshoter.SaveString(t, string(body))
	})
}
