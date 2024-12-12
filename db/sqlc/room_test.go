package db_test

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/require"
	db "github.com/vivek-344/airbnb-api/db/sqlc"
	"github.com/vivek-344/airbnb-api/util"
)

func createRandomRoom(room_id int32, t *testing.T) db.Room {
	arg := db.CreateRoomParams{
		RoomID:        room_id,
		MaxGuests:     util.RandomGuests(),
		Balcony:       util.RandomBool(),
		Fridge:        util.RandomBool(),
		IndoorPool:    util.RandomBool(),
		GamingConsole: util.RandomBool(),
	}

	room, err := testQueries.CreateRoom(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, room)

	require.Equal(t, arg.RoomID, room.RoomID)
	require.Equal(t, arg.MaxGuests, room.MaxGuests)
	require.Equal(t, arg.Balcony, room.Balcony)
	require.Equal(t, arg.Fridge, room.Fridge)
	require.Equal(t, arg.IndoorPool, room.IndoorPool)
	require.Equal(t, arg.GamingConsole, room.GamingConsole)

	return room
}

func deleteRoom(room db.Room, t *testing.T) {
	err := testQueries.DeleteRoom(context.Background(), room.RoomID)
	require.NoError(t, err)

	room, err = testQueries.GetRoom(context.Background(), room.RoomID)
	require.Error(t, err)
	require.EqualError(t, err, pgx.ErrNoRows.Error())
	require.Empty(t, room)
}

func TestCreateAccount(t *testing.T) {
	room := createRandomRoom(1, t)
	deleteRoom(room, t)
}

func TestDeleteRoom(t *testing.T) {
	room := createRandomRoom(1, t)
	deleteRoom(room, t)
}

func TestListAllRoomIDs(t *testing.T) {
	accounts, err := testQueries.ListAllRoomIDs(context.Background())

	require.NoError(t, err)

	count, _ := testQueries.GetRoomCount(context.Background())

	require.Len(t, accounts, int(count))
}

func TestListRooms(t *testing.T) {
	var i int32
	var all_rooms []db.Room
	for i = 1; i <= 10; i++ {
		room := createRandomRoom(i, t)
		all_rooms = append(all_rooms, room)
	}

	arg := db.ListRoomsParams{
		Limit:  5,
		Offset: 5,
	}

	rooms, err := testQueries.ListRooms(context.Background(), arg)

	require.NoError(t, err)

	require.Len(t, rooms, 5)

	for _, room := range all_rooms {
		require.NotEmpty(t, room)
		deleteRoom(room, t)
	}
}

func TestUpdateMaxGuests(t *testing.T) {
	room := createRandomRoom(1, t)

	args := db.UpdateMaxGuestsParams{
		RoomID:    room.RoomID,
		MaxGuests: util.RandomGuests(),
	}

	updatedRoom, err := testQueries.UpdateMaxGuests(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, updatedRoom)

	require.Equal(t, room.RoomID, updatedRoom.RoomID)
	require.Equal(t, args.MaxGuests, updatedRoom.MaxGuests)
	require.Equal(t, room.Balcony, updatedRoom.Balcony)
	require.Equal(t, room.Fridge, updatedRoom.Fridge)
	require.Equal(t, room.IndoorPool, updatedRoom.IndoorPool)
	require.Equal(t, room.GamingConsole, updatedRoom.GamingConsole)

	deleteRoom(room, t)
}

func TestUpdateRoomConsole(t *testing.T) {
	room := createRandomRoom(1, t)

	args := db.UpdateRoomConsoleParams{
		RoomID:        room.RoomID,
		GamingConsole: !room.GamingConsole,
	}

	updatedRoom, err := testQueries.UpdateRoomConsole(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, updatedRoom)

	require.Equal(t, room.RoomID, updatedRoom.RoomID)
	require.Equal(t, room.MaxGuests, updatedRoom.MaxGuests)
	require.Equal(t, room.Balcony, updatedRoom.Balcony)
	require.Equal(t, room.Fridge, updatedRoom.Fridge)
	require.Equal(t, room.IndoorPool, updatedRoom.IndoorPool)
	require.Equal(t, args.GamingConsole, updatedRoom.GamingConsole)
	require.NotEqual(t, room.GamingConsole, updatedRoom.GamingConsole)

	deleteRoom(room, t)
}

func TestUpdateRoomFridge(t *testing.T) {
	room := createRandomRoom(1, t)

	args := db.UpdateRoomFridgeParams{
		RoomID: room.RoomID,
		Fridge: !room.Fridge,
	}

	updatedRoom, err := testQueries.UpdateRoomFridge(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, updatedRoom)

	require.Equal(t, room.RoomID, updatedRoom.RoomID)
	require.Equal(t, room.MaxGuests, updatedRoom.MaxGuests)
	require.Equal(t, room.Balcony, updatedRoom.Balcony)
	require.Equal(t, args.Fridge, updatedRoom.Fridge)
	require.Equal(t, room.IndoorPool, updatedRoom.IndoorPool)
	require.Equal(t, room.GamingConsole, updatedRoom.GamingConsole)
	require.NotEqual(t, room.Fridge, updatedRoom.Fridge)

	deleteRoom(room, t)
}
