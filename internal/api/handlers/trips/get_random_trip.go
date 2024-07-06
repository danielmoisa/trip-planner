package trips

import (
	"net/http"

	"github.com/danielmoisa/trip-planner/internal/api"
	"github.com/danielmoisa/trip-planner/internal/data"
	"github.com/danielmoisa/trip-planner/internal/models"
	"github.com/danielmoisa/trip-planner/internal/util"
	"github.com/labstack/echo/v4"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func GetRandomBeerRoute(s *api.Server) *echo.Route {
	return s.Router.APIV1Trips.GET("/random", getRandomBeerHandler(s))
}

func getRandomBeerHandler(s *api.Server) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		log := util.LogFromContext(ctx)

		trip, err := models.Trips(
			qm.OrderBy("RANDOM()"),
			qm.Limit(1),
		).All(ctx, s.DB)

		if err != nil {
			log.Debug().Err(err).Msg("Failed to load random trip")
			return err
		}

		response, err := data.MarschalTrips(trip)

		if err != nil {
			return err
		}

		return util.ValidateAndReturn(c, http.StatusOK, response)

	}
}
