package util

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/vivek-344/airbnb-api/db/sqlc"
)

// FeedRoomData populates the `room` table with random test data for 25 rooms.
func FeedRoomData(queries *db.Queries) {
	for i := 100; i < 125; i++ {
		params := db.CreateRoomParams{
			RoomID:        int32(i),
			MaxGuests:     RandomGuests(),
			Balcony:       RandomBool(),
			Fridge:        RandomBool(),
			IndoorPool:    RandomBool(),
			GamingConsole: RandomBool(),
		}

		_, err := queries.CreateRoom(context.Background(), params)
		if err != nil {
			log.Fatal("error: ", err) // Logs an error and exits if room creation fails.
		}
	}
}

// FeedAvailabilityData populates the `room_availability` table with random test data
// for all rooms, ensuring data is available for the next 150 days.
func FeedAvailabilityData(queries *db.Queries) {
	// Clean up old room availability data before adding new data.
	err := queries.DeleteOldRoomAvailabilityData(context.Background())
	if err != nil {
		log.Fatalf("Failed to delete old availability data: %v", err)
	}

	// Start date for availability data generation.
	startDate := time.Now()

	// Retrieve all room IDs from the database.
	roomIDs, err := queries.ListAllRoomIDs(context.Background())
	if err != nil {
		log.Fatalf("Failed to list all room IDs: %v", err)
	}

	// Fetch the maximum existing date in the `room_availability` table.
	maxDate, err := queries.GetMaxDate(context.Background())
	if err != nil {
		log.Printf("Failed to get max date, hence using current date")
	}

	// If data exists in the future, adjust the start date accordingly.
	if maxDate.Time.After(startDate) {
		startDate = maxDate.Time
	}

	// Retrieve the current count of availability dates.
	dateCount, err := queries.GetDateCount(context.Background())
	if err != nil {
		log.Fatalf("Failed to get date count: %v", err)
	}

	// Continue adding availability data until 150 days are covered.
	for dateCount < 150 {
		for _, roomID := range roomIDs {
			availabilityDate := startDate

			// Parameters for creating a room's availability entry.
			params := db.CreateRoomAvailabilityParams{
				RoomID:      roomID,
				Date:        pgtype.Date{Time: availabilityDate, Valid: true},
				IsAvailable: RandomBool(),
				NightRate:   RandomPrice(),
			}

			// Attempt to create a room availability entry, logging and skipping on error.
			_, err := queries.CreateRoomAvailability(context.Background(), params)
			if err != nil {
				log.Printf("Failed to create room availability for room %d: %v", roomID, err)
				continue
			}
		}
		// Increment the start date and the date count.
		startDate = startDate.AddDate(0, 0, 1)
		dateCount++
	}
}
