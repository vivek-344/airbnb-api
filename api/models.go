package api

import (
	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/vivek-344/airbnb-api/db/sqlc"
)

// RoomData holds detailed information about a room and its availability for API responses.
type RoomData struct {
	RoomID              int32                             `json:"room_id"`
	RatePerNight        []DateData                        `json:"rate_per_night"`
	MaxGuests           int32                             `json:"max_guests"`
	AvailableDates      []string                          `json:"available_dates"`
	OccupancyPercentage []db.GetAvailabilityPercentageRow `json:"occupancy_percentage"`
	AverageRate         float64                           `json:"average_rate"`
	HighestRate         int32                             `json:"highest_rate"`
	LowestRate          int32                             `json:"lowest_rate"`
	Balcony             bool                              `json:"balcony"`
	MiniFridge          bool                              `json:"fridge"`
	IndoorPool          bool                              `json:"indoor_pool"`
	GamingConsole       bool                              `json:"gaming_console"`
}

// DateData holds information about the room's availability and the nightly rate for a specific date.
type DateData struct {
	Date        pgtype.Date `json:"date"`
	IsAvailable bool        `json:"is_available"`
	NightRate   int32       `json:"night_rate"`
}
