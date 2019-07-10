package dbservice

import (
	// "fmt"
	"github.com/go-pg/pg"
)

var DB *pg.DB

type dbLogger struct{}

func (d dbLogger) BeforeQuery(q *pg.QueryEvent) {
}

func (d dbLogger) AfterQuery(q *pg.QueryEvent) {
	// fmt.Println(q.FormattedQuery())
}

func Init() {
	DB = pg.Connect(&pg.Options{
		User:     "postgres",
		Password: "postgres",
		Database: "db",
	})
	DB.AddQueryHook(dbLogger{})
}
