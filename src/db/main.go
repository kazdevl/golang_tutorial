package main

import (
	"context"
	"database/sql"
	"fmt"

	"app/db/models"

	_ "github.com/lib/pq"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func main() {
	db, err := sql.Open("postgres", "dbname=postgres host=localhost user=postgres password=pass sslmode=disable")
	dieIf(err)

	err = db.Ping()
	dieIf(err)
	fmt.Println("connected")

	a := &models.Author{ID: "0003", Name: "sample3", Age: null.Int{Int: 25}}
	err = a.Insert(context.Background(), db, boil.Infer())
	dieIf(err)

	got, err := models.Authors().One(context.Background(), db)
	dieIf(err)
	fmt.Println("got author:", got.ID)

	found, err := models.FindAuthor(context.Background(), db, a.ID)
	dieIf(err)
	fmt.Println("found author:", found.ID, ",", found.Name)
}

func dieIf(err error) {
	if err != nil {
		panic(err)
	}
}
