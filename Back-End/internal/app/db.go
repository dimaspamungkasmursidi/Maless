package api

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func ConnectToDB() *pgx.Conn {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	dburl := os.Getenv("DBURL")
	if dburl == "" {
		log.Fatal("DBURL not found in .env")
	}

	con, err := pgx.Connect(context.Background(), dburl)
	if err != nil {
		log.Fatal("Error parsing DB URL:", err)
	}

	fmt.Println("Connected to DB")
	return con
}
