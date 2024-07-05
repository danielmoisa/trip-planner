package trips

import (
	"fmt"
	"net/http"

	"github.com/danielmoisa/trip-planner/internal/api"
	"github.com/danielmoisa/trip-planner/internal/api/httperrors"
	"github.com/danielmoisa/trip-planner/internal/data"
	"github.com/danielmoisa/trip-planner/internal/models"
	"github.com/danielmoisa/trip-planner/internal/types/trips"
	"github.com/danielmoisa/trip-planner/internal/util"
	"github.com/labstack/echo/v4"
)

func GetTripRoute(s *api.Server) *echo.Route {
	return s.Router.APIV1Trips.GET("/:id", getTripHandler(s))
}

func getTripHandler(s *api.Server) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		params := trips.NewGetTripRouteParams()
		log := util.LogFromContext(ctx)

		err := util.BindAndValidatePathAndQueryParams(c, &params)
		if err != nil {
			return err
		}

		trips, err := models.Trips(
			models.TripWhere.ID.EQ(params.ID.String()),
		).All(ctx, s.DB)

		if err != nil {
			log.Debug().Err(err).Msg("Failed to load trip")
			return err
		}

		if len(trips) == 0 {
			return httperrors.NewHTTPError(http.StatusNotFound, "TRIP_NOT_FOUND", fmt.Sprintf("No trip found that matches the ID %v", params.ID))
		}

		response, err := data.MarschalTrips(trips)

		if err != nil {
			return err
		}

		return util.ValidateAndReturn(c, http.StatusOK, response)

	}
}
