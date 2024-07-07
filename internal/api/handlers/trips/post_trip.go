package trips

import (
	"net/http"
	"time"

	"github.com/danielmoisa/trip-planner/internal/api"
	"github.com/danielmoisa/trip-planner/internal/api/auth"
	"github.com/danielmoisa/trip-planner/internal/models"
	"github.com/danielmoisa/trip-planner/internal/types"
	"github.com/danielmoisa/trip-planner/internal/util"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/strfmt/conv"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func PostTripRoute(s *api.Server) *echo.Route {
	return s.Router.APIV1Trips.POST("", postTripHandler(s))
}

func postTripHandler(s *api.Server) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		user := auth.UserFromContext(ctx)
		log := util.LogFromContext(ctx)

		var body types.PostTripPayload
		if err := util.BindAndValidateBody(c, &body); err != nil {
			return err
		}

		// Parse string dates
		startDate, err := time.Parse("2006-01-02", *body.StartDate)
		if err != nil {
			log.Debug().Err(err).Msg("Failed to parse start date")
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid start date format")
		}

		endDate, err := time.Parse("2006-01-02", *body.EndDate)
		if err != nil {
			log.Debug().Err(err).Msg("Failed to parse end date")
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid end date format")
		}

		// Insert new Trip
		newTrip := models.Trip{
			ID:        uuid.New().String(),
			Name:      null.StringFrom(*body.Name),
			StartDate: startDate,
			EndDate:   endDate,
			UserID:    user.ID,
		}

		if err := newTrip.Insert(ctx, s.DB, boil.Infer()); err != nil {
			log.Debug().Msg("Failed to insert trip.")
			return err
		}

		log.Debug().Msg("Successfully inserted a new trip.")

		// Prepare response
		response := &types.Trip{
			ID:        conv.UUID4(strfmt.UUID4(newTrip.ID)),
			Name:      newTrip.Name.Ptr(),
			StartDate: body.StartDate,
			EndDate:   body.EndDate,
		}

		return util.ValidateAndReturn(c, http.StatusOK, response)
	}
}
