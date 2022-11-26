package selectquery_test

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	selectquery "github.com/kazdevl/golang_tutorial/sql"

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
	data1, err := so.SelectGenderWithOverAvgIncome(selectquery.IsFemail)
	t.Log("\n\n**data1**\n")
	for index, value := range data1 {
		t.Logf("index(%d): %+v\n", index, value)
	}
	if err != nil {
		t.Error(err)
	}

	data2, err := so.SelectMaxIncomeWithUnderAvgAgeAndGenderGroup()
	t.Log("\n\n**data2**\n")
	for index, value := range data2 {
		t.Logf("index(%d): %+v\n", index, value)
	}
	if err != nil {
		t.Error(err)
	}

	data3, err := so.SelectDepartmentInfo()
	t.Log("\n\n**data3**\n")
	for index, value := range data3 {
		t.Logf("index(%d): %+v\n", index, value)
	}
	if err != nil {
		t.Error(err)
	}

	data4, err := so.SelectEmployeeInfo()
	t.Log("\n\n**data4**\n")
	for index, value := range data4 {
		t.Logf("index(%d): %+v\n", index, value)
	}
	if err != nil {
		t.Error(err)
	}
}
