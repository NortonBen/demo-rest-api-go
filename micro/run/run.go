package run

import (
	"apm/bootstrap"
	"apm/micro/lib/postgres"
	"apm/micro/load"
	"apm/micro/web"
	"apm/pkg"
	server2 "apm/pkg/server"
	"context"
	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/config"
	"github.com/micro/micro/v3/service/logger"
	"time"
)

func Handler(postgresLoad *postgres.Postgres, server *server2.Server, ctx context.Context) error {
	var jwtToken string

	val, err := config.Get("apm-jwt-token")
	if err != nil {
		logger.Error(err)
	}
	err = val.Scan(&jwtToken)
	if err != nil {
		logger.Error(err)
	}

	return bootstrap.Load(postgresLoad.Database(), server, jwtToken, ctx)
}

func Run() {
	ctx, cancel := context.WithCancel(context.Background())
	postgresLoad := postgres.NewPostgres("apm-pg")

	server := server2.NewServer(context.TODO())

	srv := service.New(
		service.Version(pkg.VERSION),
	)

	srv.Init(
		service.BeforeStart(func() error {
			logger.Infof("Start Service on: %s", time.Now().Format(time.RFC3339))
			return nil
		}),
		service.AfterStart(func() error {
			return Handler(postgresLoad, server, ctx)
		}),
		service.BeforeStop(func() error {
			cancel()
			return nil
		}),
		load.Load(
			postgresLoad,
		),
		web.Handler(server),
	)

	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
