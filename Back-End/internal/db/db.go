package db

import (
	"context"
	"log"

	"github.com/Bobby-P-dev/todo-listgo.git/internal/helpers"
	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnectToDB() *pgxpool.Pool {

	dbUrl := helpers.GetEnv("DBURL")
	log.Print("Connecting to database at URL:", dbUrl)

	conf, err := pgxpool.ParseConfig(dbUrl)
	if err != nil {
		log.Fatal("Error parsing database configuration:", err)
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), conf)
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}
	return pool
}
