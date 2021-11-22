package migrations

import (
	"apm/pkg/entities"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

func CreateSchema(db *pg.DB) error {
	models := []interface{}{
		(*entities.User)(nil),
		(*entities.File)(nil),
	}

	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			IfNotExists: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
