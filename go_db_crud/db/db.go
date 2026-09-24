package db

import (
	"context"
	"fmt"
	"os"
	"github.com/jackc/pgx/v5"
)

var Db *pgx.Conn

func ConnectDB() {
	var err error
	conString := os.Getenv("DB_STRING")
	Db, err = pgx.Connect(context.Background(), conString)
	if err != nil {
		panic(err)
	}
	fmt.Println("Database connected successfully")
}
