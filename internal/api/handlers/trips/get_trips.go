package trips

import (
	"fmt"
	"net/http"

	"github.com/danielmoisa/trip-planner/internal/api"
	"github.com/danielmoisa/trip-planner/internal/data"
	"github.com/danielmoisa/trip-planner/internal/models"
	"github.com/danielmoisa/trip-planner/internal/types/trips"
	"github.com/danielmoisa/trip-planner/internal/util"
	"github.com/danielmoisa/trip-planner/internal/util/db"
	"github.com/labstack/echo/v4"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func GetTripsRoute(s *api.Server) *echo.Route {
	return s.Router.APIV1Trips.GET("", getTripsHandler(s))
}

func getTripsHandler(s *api.Server) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		params := trips.NewGetTripsRouteParams()
		log := util.LogFromContext(ctx)

		err := util.BindAndValidatePathAndQueryParams(c, &params)
		if err != nil {
			return err
		}

		queryMods := []qm.QueryMod{
			qm.Limit(int(*params.PerPage)),
			qm.Offset(int(*params.PerPage) * (int(*params.Page) - 1)),
			qm.OrderBy("id"),
		}

		if params.TripName != nil {
			val := fmt.Sprintf("%%%s%%", *params.TripName)
			queryMods = append(queryMods, qm.Expr(db.ILike(val, models.TableNames.Trips, models.TripColumns.Name)))
		}

		trips, err := models.Trips(queryMods...).All(ctx, s.DB)

		if err != nil {
			log.Debug().Err(err).Msg("Failed to load trips")
			return err
		}

		response, err := data.MarschalTrips(trips)

		if err != nil {
			return err
		}

		return util.ValidateAndReturn(c, http.StatusOK, response)

	}
}
