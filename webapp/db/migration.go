package db

import (
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"streamer/webapp/db/models"
)

func MigrateModels(db *pg.DB) error {
	dbModels := []interface{}{
		(*models.User)(nil),
	}

	for _, model := range dbModels {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			IfNotExists: true,
			Temp:        false,
		})

		if err != nil {
			return err
		}
	}
	return nil
}
