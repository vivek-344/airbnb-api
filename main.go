package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	db "github.com/vivek-344/airbnb-api/db/sqlc"
	"github.com/vivek-344/airbnb-api/util"
)

var queries *db.Queries
var DB *pgxpool.Pool

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	DB, err = pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to the database", err)
	}

	queries = db.New(DB)

	// util.FeedRoomData(queries)

	util.FeedAvailabilityData(queries)
}
