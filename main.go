package main

import (
	"apm/bootstrap"
	"apm/pkg/lib/postgres"
	server2 "apm/pkg/server"
	"context"
	"log"
)

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	db := postgres.ConnectPostgres(&postgres.Config{
		Addr:     "10.1.1.43:5432",
		User:     "postgres",
		Password: "postgres",
		Database: "postgres",
		Schema:   "apm",
	})

	err := db.Ping(context.TODO())
	if err != nil {
		log.Fatalln(err)
	}

	server := server2.NewServer(context.TODO())

	accessSecret := "asd35oihnempfq4jiwtf9043rmpgkl(*&*#$%#s"

	err = bootstrap.Load(db, server, accessSecret, context.TODO())
	if err != nil {
		log.Fatalln(err)
		return
	}

	server.Run(":5010")
}
