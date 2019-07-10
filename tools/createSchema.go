package main

import (
	"log"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	bfdb "github.com/reanox/BitfinexLend/db"
	"github.com/reanox/BitfinexLend/service/dbservice"
)

func main() {
	dbservice.Init()
	createTable(dbservice.DB)
}

func createTable(db *pg.DB) error {
	for _, model := range []interface{}{
		(*bfdb.User)(nil),
	} {
		err := db.DropTable(model, nil)
		if err != nil {
			log.Println("createSchema error :", err)
		}
		err = db.CreateTable(model, &orm.CreateTableOptions{
			Temp: false,
		})
		if err != nil {
			log.Println("createSchema error :", err)
			continue
		}
	}
	return nil
}
