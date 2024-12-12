package db_test

import (
	"context"
	"database/sql/driver"
	"math"
	"strconv"
	"testing"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
	db "github.com/vivek-344/airbnb-api/db/sqlc"
	"github.com/vivek-344/airbnb-api/util"
)

func createRandomRoomAvailability(entryDate pgtype.Date, room db.Room, t *testing.T) db.RoomAvailability {
	arg := db.CreateRoomAvailabilityParams{
		RoomID:      room.RoomID,
		Date:        entryDate,
		IsAvailable: util.RandomBool(),
		NightRate:   util.RandomPrice(),
	}

	room_availability, err := testQueries.CreateRoomAvailability(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, room_availability)

	require.Equal(t, arg.RoomID, room_availability.RoomID)
	require.Equal(t, arg.Date, room_availability.Date)
	require.Equal(t, arg.IsAvailable, room_availability.IsAvailable)
	require.Equal(t, arg.NightRate, room_availability.NightRate)

	return room_availability
}

func TestCreateRoomAvailability(t *testing.T) {
	room := createRandomRoom(1, t)
	entryTime := time.Now().UTC()
	entryDate := pgtype.Date{Valid: true, Time: time.Date(entryTime.Year(), entryTime.Month(), entryTime.Day(), 0, 0, 0, 0, time.UTC)}
	createRandomRoomAvailability(entryDate, room, t)
	testQueries.DeleteAllAvailabilityForRoom(context.Background(), room.RoomID)
	deleteRoom(room, t)
}

func TestDeleteOldRoomAvailabilityData(t *testing.T) {
	room := createRandomRoom(1, t)
	time := time.Date(2009, 11, 1, 0, 0, 0, 0, time.UTC)
	entryDate := pgtype.Date{Valid: true, Time: time}
	createRandomRoomAvailability(entryDate, room, t)

	err := testQueries.DeleteOldRoomAvailabilityData(context.Background())
	require.NoError(t, err)

	arg := db.GetRoomAvailabilityByDateParams{
		RoomID: room.RoomID,
		Date:   pgtype.Date{Valid: true, Time: time},
	}

	room_availability, err := testQueries.GetRoomAvailabilityByDate(context.Background(), arg)
	require.Error(t, err)
	require.EqualError(t, err, pgx.ErrNoRows.Error())
	require.Empty(t, room_availability)

	deleteRoom(room, t)
}

func TestGetAvailabilityPercentage(t *testing.T) {
	room := createRandomRoom(1, t)
	entryTime := time.Now().UTC()
	startDate := time.Date(entryTime.Year(), entryTime.Month(), entryTime.Day(), 0, 0, 0, 0, time.UTC)
	count := 0
	var i int
	var all_rooms_availabilities []db.RoomAvailability
	for i = 1; i <= 10; i++ {
		entryDate := pgtype.Date{Valid: true, Time: startDate}
		room_availability := createRandomRoomAvailability(entryDate, room, t)
		if room_availability.IsAvailable {
			count++
		}
		all_rooms_availabilities = append(all_rooms_availabilities, room_availability)
		startDate = startDate.AddDate(0, 0, 1)
		if startDate.Month() != entryTime.Month() {
			break
		}
	}

	expected_availability_percentage := float64(count) * 10
	expected_availability_percentage_float := strconv.FormatFloat(expected_availability_percentage, 'f', 2, 64)
	var calculated_availability_percentage driver.Value

	availability_percentages, err := testQueries.GetAvailabilityPercentage(context.Background(), room.RoomID)

	for _, availability_percentage := range availability_percentages {
		year, _ := availability_percentage.Year.Value()
		month, _ := availability_percentage.Month.Value()

		if year == strconv.Itoa(entryTime.Year()) && month == strconv.Itoa(int(entryTime.Month())) {
			percentage, _ := availability_percentage.AvailabilityPercentage.Value()
			calculated_availability_percentage = percentage.(string)
		}
	}

	if expected_availability_percentage_float == "0.00" {
		expected_availability_percentage_float = "0"
	}

	require.NoError(t, err)
	require.Equal(t, expected_availability_percentage_float, calculated_availability_percentage)

	testQueries.DeleteAllAvailabilityForRoom(context.Background(), room.RoomID)
	deleteRoom(room, t)
}

func TestGetAverageRate(t *testing.T) {
	room := createRandomRoom(1, t)
	entryTime := time.Now().UTC()
	startDate := time.Date(entryTime.Year(), entryTime.Month(), entryTime.Day(), 0, 0, 0, 0, time.UTC)
	var total_price int32
	var i int
	var all_rooms_availabilities []db.RoomAvailability
	for i = 1; i <= 10; i++ {
		entryDate := pgtype.Date{Valid: true, Time: startDate}
		room_availability := createRandomRoomAvailability(entryDate, room, t)
		total_price += room_availability.NightRate
		all_rooms_availabilities = append(all_rooms_availabilities, room_availability)
		startDate = startDate.AddDate(0, 0, 1)
	}

	expected_average := float64(total_price) / 10
	calculated_average, err := testQueries.GetAverageRate(context.Background(), room.RoomID)
	require.NoError(t, err)
	require.Equal(t, expected_average, calculated_average)

	testQueries.DeleteAllAvailabilityForRoom(context.Background(), room.RoomID)
	deleteRoom(room, t)
}

func TestGetDateCount(t *testing.T) {
	room := createRandomRoom(1, t)
	entryTime := time.Now().UTC()
	entryDate := pgtype.Date{Valid: true, Time: time.Date(entryTime.Year(), entryTime.Month(), entryTime.Day(), 0, 0, 0, 0, time.UTC)}
	createRandomRoomAvailability(entryDate, room, t)

	count, err := testQueries.GetDateCount(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, count)

	testQueries.DeleteAllAvailabilityForRoom(context.Background(), room.RoomID)
	deleteRoom(room, t)
}

func TestGetMaxDate(t *testing.T) {
	room := createRandomRoom(1, t)
	entryTime := time.Now().UTC()
	entryDate := pgtype.Date{Valid: true, Time: time.Date(entryTime.Year(), entryTime.Month(), entryTime.Day(), 0, 0, 0, 0, time.UTC)}
	createRandomRoomAvailability(entryDate, room, t)

	max_date, err := testQueries.GetMaxDate(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, max_date)

	testQueries.DeleteAllAvailabilityForRoom(context.Background(), room.RoomID)
	deleteRoom(room, t)
}

func TestGetMaximumRate(t *testing.T) {
	room := createRandomRoom(1, t)
	entryTime := time.Now().UTC()
	startDate := time.Date(entryTime.Year(), entryTime.Month(), entryTime.Day(), 0, 0, 0, 0, time.UTC)
	var max_price float64
	var i int
	var all_rooms_availabilities []db.RoomAvailability
	for i = 1; i <= 10; i++ {
		entryDate := pgtype.Date{Valid: true, Time: startDate}
		room_availability := createRandomRoomAvailability(entryDate, room, t)
		max_price = math.Max(float64(max_price), float64(room_availability.NightRate))
		all_rooms_availabilities = append(all_rooms_availabilities, room_availability)
		startDate = startDate.AddDate(0, 0, 1)
	}

	calculated_max_price, err := testQueries.GetMaximumRate(context.Background(), room.RoomID)
	require.NoError(t, err)
	require.Equal(t, max_price, float64(calculated_max_price))

	testQueries.DeleteAllAvailabilityForRoom(context.Background(), room.RoomID)
	deleteRoom(room, t)
}

func TestGetMinimumRate(t *testing.T) {
	room := createRandomRoom(1, t)
	entryTime := time.Now().UTC()
	startDate := time.Date(entryTime.Year(), entryTime.Month(), entryTime.Day(), 0, 0, 0, 0, time.UTC)
	min_price := float64(100001)
	var i int
	var all_rooms_availabilities []db.RoomAvailability
	for i = 1; i <= 10; i++ {
		entryDate := pgtype.Date{Valid: true, Time: startDate}
		room_availability := createRandomRoomAvailability(entryDate, room, t)
		min_price = math.Min(float64(min_price), float64(room_availability.NightRate))
		all_rooms_availabilities = append(all_rooms_availabilities, room_availability)
		startDate = startDate.AddDate(0, 0, 1)
	}

	calculated_min_price, err := testQueries.GetMinimumRate(context.Background(), room.RoomID)
	require.NoError(t, err)
	require.Equal(t, min_price, float64(calculated_min_price))

	testQueries.DeleteAllAvailabilityForRoom(context.Background(), room.RoomID)
	deleteRoom(room, t)
}

func TestGetRoomAvailabilityByDate(t *testing.T) {
	room := createRandomRoom(1, t)
	time := time.Date(2009, 11, 1, 0, 0, 0, 0, time.UTC)
	entryDate := pgtype.Date{Valid: true, Time: time}
	room_availability := createRandomRoomAvailability(entryDate, room, t)

	arg := db.GetRoomAvailabilityByDateParams{
		RoomID: room.RoomID,
		Date:   entryDate,
	}
	availability_data, err := testQueries.GetRoomAvailabilityByDate(context.Background(), arg)
	require.NoError(t, err)

	require.Equal(t, room_availability.RoomID, availability_data.RoomID)
	require.Equal(t, room_availability.Date, availability_data.Date)
	require.Equal(t, room_availability.IsAvailable, availability_data.IsAvailable)
	require.Equal(t, room_availability.NightRate, availability_data.NightRate)

	testQueries.DeleteAllAvailabilityForRoom(context.Background(), room.RoomID)
	deleteRoom(room, t)
}

func TestListAvailableDates(t *testing.T) {
	room := createRandomRoom(1, t)
	var date_list []pgtype.Date
	entryTime := time.Now().UTC()
	startDate := time.Date(entryTime.Year(), entryTime.Month(), entryTime.Day(), 0, 0, 0, 0, time.UTC)
	var all_rooms_availabilities []db.RoomAvailability
	for i := 1; i <= 10; i++ {
		entryDate := pgtype.Date{Valid: true, Time: startDate}
		room_availability := createRandomRoomAvailability(entryDate, room, t)
		if room_availability.IsAvailable {
			date_list = append(date_list, room_availability.Date)
		}
		all_rooms_availabilities = append(all_rooms_availabilities, room_availability)
		startDate = startDate.AddDate(0, 0, 1)
	}

	available_dates, err := testQueries.ListAvailableDates(context.Background(), room.RoomID)
	require.NoError(t, err)
	require.Equal(t, date_list, available_dates)

	testQueries.DeleteAllAvailabilityForRoom(context.Background(), room.RoomID)
	deleteRoom(room, t)
}

func TestListRoomAvailability(t *testing.T) {
	room := createRandomRoom(1, t)
	var all_room_availability_rows []db.ListRoomAvailabilityRow
	entryTime := time.Now().UTC()
	startDate := time.Date(entryTime.Year(), entryTime.Month(), entryTime.Day(), 0, 0, 0, 0, time.UTC)
	var all_rooms_availabilities []db.RoomAvailability
	for i := 1; i <= 10; i++ {
		entryDate := pgtype.Date{Valid: true, Time: startDate}
		room_availability := createRandomRoomAvailability(entryDate, room, t)
		room_availability_data := db.ListRoomAvailabilityRow{
			Date:        room_availability.Date,
			IsAvailable: room_availability.IsAvailable,
			NightRate:   room_availability.NightRate,
		}
		all_room_availability_rows = append(all_room_availability_rows, room_availability_data)
		all_rooms_availabilities = append(all_rooms_availabilities, room_availability)
		startDate = startDate.AddDate(0, 0, 1)
	}

	all_room_availability_data, err := testQueries.ListRoomAvailability(context.Background(), room.RoomID)
	require.NoError(t, err)
	require.Equal(t, all_room_availability_rows, all_room_availability_data)

	testQueries.DeleteAllAvailabilityForRoom(context.Background(), room.RoomID)
	deleteRoom(room, t)
}

func TestUpdateRoomAvailability(t *testing.T) {
	room := createRandomRoom(1, t)
	entryTime := time.Now().UTC()
	entryDate := pgtype.Date{Valid: true, Time: time.Date(entryTime.Year(), entryTime.Month(), entryTime.Day(), 0, 0, 0, 0, time.UTC)}
	availability_data := createRandomRoomAvailability(entryDate, room, t)

	newEntryTime := entryTime.AddDate(0, 0, 5)
	newEntryDate := pgtype.Date{Valid: true, Time: time.Date(newEntryTime.Year(), entryTime.Month(), entryTime.Day(), 0, 0, 0, 0, time.UTC)}
	arg := db.UpdateRoomAvailabilityParams{
		RoomID:      room.RoomID,
		Date:        newEntryDate,
		IsAvailable: !availability_data.IsAvailable,
		NightRate:   util.RandomPrice(),
	}
	updated_availability_data, err := testQueries.UpdateRoomAvailability(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, room.RoomID, updated_availability_data.RoomID)
	require.Equal(t, arg.Date, updated_availability_data.Date)
	require.Equal(t, arg.IsAvailable, updated_availability_data.IsAvailable)
	require.Equal(t, arg.NightRate, updated_availability_data.NightRate)

	testQueries.DeleteAllAvailabilityForRoom(context.Background(), room.RoomID)
	deleteRoom(room, t)
}
