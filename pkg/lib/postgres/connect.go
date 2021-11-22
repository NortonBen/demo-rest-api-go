package postgres

import (
	"context"
	"github.com/go-pg/pg/v10"
	"log"
)

func ConnectPostgres(postgres *Config) *pg.DB {
	schemaName := "public"
	if postgres.Schema != "" {
		schemaName = postgres.Schema
	}

	db := pg.Connect(&pg.Options{
		Addr:            postgres.Addr,
		User:            postgres.User,
		Password:        postgres.Password,
		Database:        postgres.Database,
		ApplicationName: "APM",
		OnConnect: func(c context.Context, conn *pg.Conn) error {
			_, err := conn.ExecContext(c, "set search_path=?", schemaName)
			if err != nil {
				log.Fatalln(err)
			}
			return nil
		},
	}).WithParam("search_path", schemaName)

	return db
}
