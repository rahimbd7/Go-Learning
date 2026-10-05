package main

import (
	"go_tickets/internal/config"
	"go_tickets/internal/server"
)

func main() {
	cfg := config.LoadEnv()
	db := config.ConnectDatabase(cfg)
	server.Start(db, cfg)
}
