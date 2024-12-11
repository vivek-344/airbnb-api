package api

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	db "github.com/vivek-344/airbnb-api/db/sqlc"
)

// Store provides all functions to execute db queries.
type Store struct {
	*db.Queries
	conn *pgxpool.Pool
}

// NewStore creates a new Store that wraps the database connection pool and query methods.
func NewStore(conn *pgxpool.Pool) *Store {
	return &Store{
		conn:    conn,
		Queries: db.New(conn),
	}
}

// getRoomRequest defines the expected URI parameters for fetching room data.
type getRoomRequest struct {
	RoomID int32 `binding:"required,min=1" uri:"room_id"`
}

// getRoomData is a handler method that delegates the request to the store.
func (server *Server) getRoomData(ctx *gin.Context) {
	server.store.getRoomData(ctx)
}

// getRoomData fetches detailed information about a room, including its availability, rates, and stats.
func (store *Store) getRoomData(ctx *gin.Context) {
	var req getRoomRequest

	// Bind URI parameters to the request struct and validate them.
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()}) // Respond with 400 Bad Request on validation failure.
		return
	}

	// Fetch room details based on RoomID.
	room, err := store.GetRoom(ctx, req.RoomID)
	if err != nil {
		ctx.JSON(404, gin.H{"error": "Room not found"}) // Respond with 404 if the room doesn't exist.
		return
	}

	// Fetch room availability and nightly rates for the next 30 days.
	ratePerNight, err := store.ListRoomAvailability(ctx, req.RoomID)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to fetch room availability"}) // Respond with 500 if fetching fails.
		return
	}

	// Convert the availability data into a format suitable for the API response.
	var dateData []DateData
	for _, availability := range ratePerNight {
		dateData = append(dateData, DateData{
			Date:        availability.Date,
			IsAvailable: availability.IsAvailable,
			NightRate:   availability.NightRate,
		})
	}

	// Fetch a list of all available dates for the room.
	availableDates, err := store.ListAvailableDates(ctx, req.RoomID)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to fetch available dates"}) // Respond with 500 on failure.
		return
	}

	// Convert the available dates from `pgtype.Date` to string for API response.
	var availableDateStrings []string
	for _, date := range availableDates {
		availableDateStrings = append(availableDateStrings, date.Time.Format("2006-01-02"))
	}

	// Fetch the occupancy percentage for the room and handle errors gracefully.
	var occupancyPercentage []db.GetAvailabilityPercentageRow
	occupancyPercentage, err = store.GetAvailabilityPercentage(ctx, req.RoomID)
	if err != nil {
		log.Printf("Error fetching occupancy percentage for room %d: %v", req.RoomID, err)
		occupancyPercentage = []db.GetAvailabilityPercentageRow{} // Use an empty slice instead of nil.
	}

	// Fetch the average nightly rate for the room.
	averageRate, err := store.GetAverageRate(ctx, req.RoomID)
	if err != nil {
		log.Printf("Error fetching average rate for room %d: %v", req.RoomID, err)
		averageRate = 0 // Default to 0 if the average rate can't be calculated.
	}

	// Fetch the highest nightly rate for the room.
	highestRate, err := store.GetMaximumRate(ctx, req.RoomID)
	if err != nil {
		log.Printf("Error fetching highest rate for room %d: %v", req.RoomID, err)
		highestRate = 0 // Default to 0 if no rate is found.
	}

	// Fetch the lowest nightly rate for the room.
	lowestRate, err := store.GetMinimumRate(ctx, req.RoomID)
	if err != nil {
		log.Printf("Error fetching lowest rate for room %d: %v", req.RoomID, err)
		lowestRate = 0 // Default to 0 if no rate is found.
	}

	// Prepare the response object with all room details and statistics.
	roomData := RoomData{
		RoomID:              room.RoomID,
		RatePerNight:        dateData,
		MaxGuests:           room.MaxGuests,
		AvailableDates:      availableDateStrings,
		OccupancyPercentage: occupancyPercentage,
		AverageRate:         averageRate,
		HighestRate:         highestRate,
		LowestRate:          lowestRate,
		Balcony:             room.Balcony,
		MiniFridge:          room.Fridge,
		IndoorPool:          room.IndoorPool,
		GamingConsole:       room.GamingConsole,
	}

	// Send the room data as a JSON response with HTTP status 200.
	ctx.JSON(200, roomData)
}
