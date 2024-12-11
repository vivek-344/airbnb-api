package util

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/vivek-344/airbnb-api/db/sqlc"
)

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
			log.Fatal("error: ", err)
		}
	}
}

func FeedAvailabilityData(queries *db.Queries) {
	err := queries.DeleteOldRoomAvailabilityData(context.Background())
	if err != nil {
		log.Fatalf("Failed to delete old availability data: %v", err)
	}

	startDate := time.Now()

	roomIDs, err := queries.ListAllRoomIDs(context.Background())
	if err != nil {
		log.Fatalf("Failed to list all room IDs: %v", err)
	}

	maxDate, err := queries.GetMaxDate(context.Background())
	if err != nil {
		log.Printf("Failed to get max date: %v", err)
	}

	if maxDate.Time.After(startDate) {
		startDate = maxDate.Time
	}

	dateCount, err := queries.GetDateCount(context.Background())
	if err != nil {
		log.Fatalf("Failed to get date count: %v", err)
	}

	for dateCount < 150 {
		for _, roomID := range roomIDs {
			availabilityDate := startDate

			params := db.CreateRoomAvailabilityParams{
				RoomID:      roomID,
				Date:        pgtype.Date{Time: availabilityDate, Valid: true},
				IsAvailable: RandomBool(),
				NightRate:   RandomPrice(),
			}

			_, err := queries.CreateRoomAvailability(context.Background(), params)
			if err != nil {
				log.Printf("Failed to create room availability for room %d: %v", roomID, err)
				continue
			}
		}
		startDate = startDate.AddDate(0, 0, 1)
		dateCount++
	}
}
