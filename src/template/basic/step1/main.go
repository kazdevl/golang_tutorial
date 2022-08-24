package main

import (
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"text/template"
	"time"
)

func main() {
	// 1. 入力データの用意
	type Field struct {
		Name string
		Type string
		IsPK bool
	}
	type input struct {
		Name   string
		Fields []Field
	}

	is := []input{
		{
			Name: "user",
			Fields: []Field{
				{Name: "ID", Type: reflect.TypeOf(int64(0)).String(), IsPK: true},
				{Name: "Name", Type: reflect.TypeOf("").String()},
				{Name: "LastLoginDate", Type: reflect.TypeOf(time.Time{}).String()},
			},
		},
	}
	toUpperCase := func(s string) string {
		return strings.ToUpper(s[0:1]) + s[1:]
	}

	// 2. Template構造体のインスタンスを作成する
	// Funcsで独自定義関数の追加
	// template.ParseFilesでmodel.tplを解析している(内部的にtemplateを表現するTreeという構造体を生成)
	firstFilePath := "model.tpl"
	tmpl := template.Must(template.New(firstFilePath).Funcs(template.FuncMap{
		"toUpperCase": toUpperCase,
	}).ParseFiles(firstFilePath))

	// 3. インスタンスに入力データを渡して実行結果をファイルに書き込む
	// 定義した入力データ
	for _, i := range is {
		f, err := os.Create(filepath.Join("result", i.Name+"modelrepository.go"))
		if err != nil {
			panic(err)
		}

		if err = tmpl.Execute(f, i); err != nil {
			panic(err)
		}
		f.Close()
	}
}
