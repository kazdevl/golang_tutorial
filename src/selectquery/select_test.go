package selectquery_test

import (
	"app/selectquery"
	"database/sql"
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestSelect(t *testing.T) {
	// dbとの接続
	if err := godotenv.Load(); err != nil {
		fmt.Printf("%v\n", err)
	}
	db, err := sql.Open("mysql", os.Getenv("DSN"))
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	// handlerの初期化
	so := selectquery.NewSelectOperator(db)
	data, err := so.SelectGenderWithOverAvgIncome(selectquery.IsFemail)
	t.Log(data)
	if err != nil {
		t.Error(err)
	}
}
