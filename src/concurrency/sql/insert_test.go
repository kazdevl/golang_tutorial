package concurrency_test

import (
	"database/sql"
	"fmt"
	"os"
	"runtime"
	"testing"

	concurrency "github.com/kazdevl/golang_tutorial/concurrency/sql"

	"github.com/joho/godotenv"
)

func BenchmarkInsert(b *testing.B) {
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
	handler := concurrency.NewSqlHandler(db)
	// dbに保存したいデータ
	books := make([]concurrency.BookModel, 500)
	for index := 0; index < len(books); index++ {
		books[index] = concurrency.BookModel{Name: "test", Value: index + 100}
	}

	insertForBnechMark := func(bms []concurrency.BookModel) {
		for _, b := range bms {
			if err := handler.Insert(b.Name, b.Value); err != nil {
				fmt.Println(err)
			}
		}
	}

	bulkInsertForBnechMark := func(bms []concurrency.BookModel) {
		// bulkinsertWithConcurrencyと同時に挿入しているレコードの数を合わせている
		threadNum := runtime.NumCPU()
		contentsNumPerThread := len(bms) / threadNum
		getNumOfGoroutineOfBulkInsert := func(contentsNum int) int {
			if contentsNum > 200 {
				return 4
			}
			return 2
		}

		numOfGoroutine := getNumOfGoroutineOfBulkInsert(contentsNumPerThread)
		contentsNumOfInsert := contentsNumPerThread / numOfGoroutine
		for i := 0; i < len(bms)/contentsNumOfInsert; i++ {
			if err := handler.BulkInsert(bms[i*contentsNumOfInsert : (i+1)*contentsNumOfInsert]); err != nil {
				fmt.Println(err)
			}
		}
	}

	tests := []struct {
		name     string
		function func(bms []concurrency.BookModel)
	}{
		{
			name:     "Insert",
			function: insertForBnechMark,
		},
		{
			name:     "BulkInsert",
			function: bulkInsertForBnechMark,
		},
		{
			name:     "BulkInsertWithConcurrency",
			function: handler.BulkInsertWithConcurrency,
		},
	}
	for _, test := range tests {
		b.Run(test.name, func(b *testing.B) {
			b.ResetTimer()
			test.function(books)
		})
	}

}
