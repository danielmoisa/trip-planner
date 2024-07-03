package data

import (
	"github.com/danielmoisa/trip-planner/internal/models"
	"github.com/danielmoisa/trip-planner/internal/types"
)

func MarschalTrips(trips models.TripSlice) (types.GetTripsResponse, error) {
	response := make(types.GetTripsResponse, len(trips))

	for i, trip := range trips {

		marschalledTrip, err := MarschalTrip(*trip)

		if err != nil {
			return nil, err
		}

		response[i] = marschalledTrip
	}

	return response, nil
}

func MarschalTrip(trip models.Trip) (*types.Trip, error) {

	startDateStr := trip.StartDate.Format("2006-01-02")
	endDateStr := trip.EndDate.Format("2006-01-02")

	return &types.Trip{
		ID:        &trip.ID,
		Name:      &trip.Name.String,
		StartDate: &startDateStr,
		EndDate:   &endDateStr,
	}, nil
}
