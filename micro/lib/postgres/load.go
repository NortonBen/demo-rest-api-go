package postgres

import (
	"apm/micro/load"
	"context"
	"github.com/go-pg/pg/v10"
	"github.com/micro/micro/v3/service/config"
	"github.com/micro/micro/v3/service/logger"
	"github.com/urfave/cli/v2"
)

type Postgres struct {
	db         *pg.DB
	Config     *Config
	configName string
}

func (r *Postgres) Flag() []cli.Flag {
	return []cli.Flag{}
}

func (r *Postgres) Loader() load.Loader {
	return r
}

func (p *Postgres) Start() error {

	val, err := config.Get(p.configName)
	if err != nil {
		logger.Error("Config Get", err)
		return err
	}
	err = val.Scan(&p.Config)
	if err != nil {
		logger.Error("Config Get", err)
		return err
	}

	p.db = ConnectPostgres(p.Config)

	err = p.db.Ping(context.TODO())
	if err != nil {
		return err
	}

	return nil
}

func NewPostgres(configName string) *Postgres {
	return &Postgres{
		configName: configName,
		Config:     &Config{},
	}
}

func (p *Postgres) Database() *pg.DB {
	return p.db
}

func (p Postgres) Name() string {
	return "postgres"
}

func (p *Postgres) Stop() error {
	return p.db.Close()
}
