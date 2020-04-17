package db

import (
	"log"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"github.com/reanox/BitfinexLend/db/models"
)

func CreateTable(db *pg.DB) error {
	for _, model := range []interface{}{
		(*models.User)(nil),
	} {
		err := db.CreateTable(model, &orm.CreateTableOptions{
			Temp:        false,
			IfNotExists: true,
		})
		if err != nil {
			log.Println("createSchema error :", err)
			continue
		}
	}
	return nil
}
