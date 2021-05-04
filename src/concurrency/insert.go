package concurrency

import (
	"database/sql"
	"fmt"
	"log"
	"runtime"
	"strings"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

type bookModel struct {
	id    int
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

func (handler *SqlHandler) Insert(name string, value int) error {
	query := "INSERT INTO book(name, value) values (?,?)"
	if _, err := handler.DB.Exec(query, name, value); err != nil {
		return err
	}
	return nil
}

func (handler *SqlHandler) BulkInsert(bms []bookModel) error {
	makePlaceholder := func(length int) string {
		values := make([]string, length)
		for index := 0; index < length; index++ {
			values[index] = "(?, ?)"
		}
		return strings.Join(values, ",")
	}
	query := fmt.Sprintf("INSERT INTO book(name, value) values %s", makePlaceholder(len(bms)))

	converterForExec := func(bms []bookModel) []interface{} {
		bind := make([]interface{}, len(bms)*2)
		var index int = 0
		for _, b := range bms {
			bind[index] = b.name
			bind[index+1] = b.value
			index += 2
		}
		return bind
	}
	if _, err := handler.DB.Exec(query, converterForExec(bms)...); err != nil {
		return err
	}
	return nil
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

func (handler *SqlHandler) BulkInsertWithConcurrency(bms []bookModel) {
	// goroutineの数を算出
	threadNum := runtime.NumCPU()
	// 1gorouitneごとに処理するコンテンツの数
	contentsNumPerThread := len(bms) / threadNum

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
				if err := handler.BulkInsert(bms[index : index+contentsNumPerLoop]); err != nil {
					log.Fatal(err)
				}
				index += contentsNumPerLoop
				m.Unlock()
			}
		}()
	}
	wg.Wait()
	if index <= len(bms)-1 {
		if err := handler.BulkInsert(bms[index:]); err != nil {
			log.Fatal(err)
		}
	}
}
