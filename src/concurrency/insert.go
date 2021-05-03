package concurrency

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
	"sync"

	_ "github.com/go-sql-driver/mysql"

	"github.com/joho/godotenv"
)

type bookModel struct {
	id    int
	name  string
	value int
}

type bookForInsert struct {
	name  string
	value int
}

type SqlHandler struct {
	DB *sql.DB
}

func NewSqlHandler(db *sql.DB) *SqlHandler {
	return &SqlHandler{
		DB: db,
	}
}

func (handler *SqlHandler) BulkInsert(bs []bookForInsert) error {
	// makeValuesForInsert := func(bs []bookForInsert) string {
	// 	values := make([]string, len(bs))
	// 	for index, b := range bs {
	// 		values[index] = fmt.Sprintf("('%s', %d)", b.name, b.value) //脆弱性が残っている実装になっている...https://www.ipa.go.jp/files/000017320.pdf
	// 	}
	// 	return strings.Join(values, ",")
	// }
	// query := "INSERT INTO book(name, value) values ?"
	// if _, err := handler.DB.Exec(query, makeValuesForInsert(bs)); err != nil {
	// 	return err
	// }
	// return nil

	makePlaceholder := func(length int) string {
		values := make([]string, length)
		for index := 0; index < length; index++ {
			values[index] = "(?, ?)"
		}
		return strings.Join(values, ",")
	}
	query := fmt.Sprintf("INSERT INTO book(name, value) values %s", makePlaceholder(len(bs)))

	converterForExec := func(bs []bookForInsert) []interface{} {
		bind := make([]interface{}, len(bs))
		for index, b := range bs {
			bind[index] = b
		}
		return bind
	}
	if _, err := handler.DB.Exec(query, converterForExec(bs)); err != nil {
		return err
	}
	return nil

}

func (handler *SqlHandler) Get(ids []int) ([]bookModel, error) {
	makePlaceholder := func(length int) string {
		values := make([]string, length)
		for index := 0; index < length; index++ {
			values[index] = "?"
		}
		return strings.Join(values, ",")
	}
	query := fmt.Sprintf("SELECT * FROM book WHERE id in (%s)", makePlaceholder(len(ids)))
	converterForExec := func(is []int) []interface{} {
		bind := make([]interface{}, len(is))
		for index, i := range is {
			bind[index] = i
		}
		return bind
	}
	rows, err := handler.DB.Query(query, converterForExec(ids))
	if err != nil {
		return []bookModel{}, err
	}

	bm := make([]bookModel, len(ids))
	for index := 0; rows.Next(); index++ {
		if err := rows.Scan(&bm[index]); err != nil {
			return []bookModel{}, err
		}
	}
	return bm, nil
}

func BulkInsertWithConcurrency() {
	// dbとの接続
	if err := godotenv.Load("./../.env"); err != nil {
		fmt.Printf("%v\n", err)
	}
	fmt.Printf("DB_USER: %s\n", os.Getenv("DSN"))
	db, err := sql.Open("mysql", os.Getenv("DSN"))
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	// handlerの初期化
	handler := NewSqlHandler(db)

	// dbに保存したいデータ
	books := make([]bookForInsert, 3000)
	for index := 0; index < len(books); index++ {
		books[index] = bookForInsert{name: "test", value: index + 100}
	}
	// goroutineの数を算出
	threadNum := runtime.NumCPU()
	// 1gorouitneごとに処理するコンテンツの数
	contentsNumPerThread := len(books) / threadNum

	// 排他処理を制御する変数
	var m sync.Mutex
	// goroutineの終了を待つ変数
	var wg sync.WaitGroup
	// bulkInsertの位置を調整する変数
	var index int = 0
	for i := 0; i < threadNum; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// コンテンツを処理するのに実行するbulkinsertの回数
			getNumOfExecBulkInsert := func(contentsNum int) int {
				if contentsNum > 200 {
					return 4
				}
				return 2
			}

			for loopNum := 0; loopNum < getNumOfExecBulkInsert(contentsNumPerThread); loopNum++ {
				m.Lock()
				if err := handler.BulkInsert(books[index : index+contentsNumPerThread]); err != nil {
					log.Fatal(err)
				}
				index += contentsNumPerThread
				m.Unlock()
			}
		}()
	}
	wg.Wait()
	if err := handler.BulkInsert(books[index:]); err != nil {
		log.Fatal(err)
	}

}
