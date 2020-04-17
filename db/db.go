package db

import (
	// "fmt"
	"github.com/go-pg/pg"
)

type dbLogger struct{}

func (d dbLogger) BeforeQuery(q *pg.QueryEvent) {
}

func (d dbLogger) AfterQuery(q *pg.QueryEvent) {
	// fmt.Println(q.FormattedQuery())
}

func NewDB(user string, password string, dbname string) *pg.DB {
	db := pg.Connect(&pg.Options{
		User:     user,
		Password: password,
		Database: dbname,
	})

	db.AddQueryHook(dbLogger{})
	err := CreateTable(db)
	if err != nil {
		panic(err)
	}

	return db
}
