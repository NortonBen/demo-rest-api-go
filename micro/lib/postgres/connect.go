package postgres

import (
	"apm/pkg"
	"context"
	"github.com/go-pg/pg/v10"
	"github.com/micro/micro/v3/service/logger"
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
		ApplicationName: pkg.NAME,
		OnConnect: func(c context.Context, conn *pg.Conn) error {
			_, err := conn.ExecContext(c, "set search_path=?", schemaName)
			if err != nil {
				logger.Fatal(err)
			}
			return nil
		},
	}).WithParam("search_path", schemaName)

	return db
}
