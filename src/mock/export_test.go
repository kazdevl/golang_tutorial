package mock

var ExportFn func() string = wantExport

// 型をexportするには、Go1.9に変更する必要がある。
// それで、型エイリアスの機能を使うと、型に別名をつけられる
// https://golang.org/doc/go1.9#language
// type ExportUser = user
