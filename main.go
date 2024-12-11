package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/vivek-344/airbnb-api/api"
	"github.com/vivek-344/airbnb-api/util"
)

func main() {
	// Load the application configuration from the environment.
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	// Establish a connection pool to the PostgreSQL database.
	conn, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to the database", err)
	}

	// Initialize the database store with the connection pool.
	store := api.NewStore(conn)

	// Create a new API server with the initialized store.
	server := api.NewServer(*store)

	// Uncomment the following line to populate the `room` table with test data.
	// util.FeedRoomData(store.Queries)

	// Populate the `room_availability` table with test data.
	util.FeedAvailabilityData(store.Queries)

	// Start the server and listen on the configured address.
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}
