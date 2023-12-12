package main

import (
	"context"

	"github.com/guatom999/sharePrice/config"
	"github.com/guatom999/sharePrice/database"
	"github.com/guatom999/sharePrice/server"
)

func main() {

	ctx := context.Background()

	cfg := config.GetConfig()

	db := database.NewPostgresDB(&cfg).GetDb()

	server.NewEchoServer(db, &cfg).Start(ctx)
}
