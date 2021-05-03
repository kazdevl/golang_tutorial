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
	makeValuesForInsert := func(bs []bookForInsert) string {
		values := make([]string, len(bs))
		for index, b := range bs {
			values[index] = fmt.Sprintf("('%s',%d)", b.name, b.value) //脆弱性が残っている実装になっている...https://www.ipa.go.jp/files/000017320.pdf
		}
		return strings.Join(values, ",")
	}
	query := fmt.Sprintf("insert into book(name,value) values %s", makeValuesForInsert(bs))
	if _, err := handler.DB.Exec(query); err != nil {
		return err
	}
	return nil
	// makePlaceholder := func(length int) string {
	// 	values := make([]string, length)
	// 	for index := 0; index < length; index++ {
	// 		values[index] = "(?, ?)"
	// 	}
	// 	return strings.Join(values, ",")
	// }
	// query := fmt.Sprintf("INSERT INTO book(name, value) values %s", makePlaceholder(len(bs)))

	// converterForExec := func(bs []bookForInsert) []interface{} {
	// 	bind := make([]interface{}, len(bs))
	// 	for index, b := range bs {
	// 		bind[index] = b
	// 	}
	// 	return bind
	// }
	// if _, err := handler.DB.Exec(query, converterForExec(bs)); err != nil {//ここでエラーが出る
	// 	return err
	// }
	// return nil
}

func (handler *SqlHandler) BulkGet(ids []int) ([]bookModel, error) {
	makePlaceholder := func(ids []int) string {
		values := make([]string, len(ids))
		for index := 0; index < len(ids); index++ {
			values[index] = fmt.Sprintf("%d", ids[index])
		}
		return strings.Join(values, ",")
	}
	query := fmt.Sprintf("SELECT * FROM book WHERE id in (%s)", makePlaceholder(ids))
	rows, err := handler.DB.Query(query)
	if err != nil {
		return []bookModel{}, err
	}

	bms := make([]bookModel, len(ids))
	for index := 0; rows.Next(); index++ {
		if err := rows.Scan(&bms[index].id, &bms[index].name, &bms[index].value); err != nil {
			return []bookModel{}, err
		}
	}
	return bms, nil
}

func BulkInsertWithConcurrency() {
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

			loopNum := getNumOfExecBulkInsert(contentsNumPerThread)
			contentsNumPerLoop := contentsNumPerThread / loopNum
			for i := 0; i < loopNum; i++ {
				m.Lock()
				if err := handler.BulkInsert(books[index : index+contentsNumPerLoop]); err != nil {
					log.Fatal(err)
				}
				index += contentsNumPerLoop
				m.Unlock()
			}
		}()
	}
	wg.Wait()
	if index <= len(books)-1 {
		if err := handler.BulkInsert(books[index:]); err != nil {
			log.Fatal(err)
		}
	}

	data := make([]int, 50)
	for i := 0; i < len(data); i++ {
		data[i] = i
	}
	bms, err := handler.BulkGet(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v", bms)
}
