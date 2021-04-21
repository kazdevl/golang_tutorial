# golang_tutorial
golangを完全に理解するためのリポジトリ

## 参考URL
- [チュートリアル](https://go-tour-jp.appspot.com/basics/11)
- [言語仕様書](https://golang.org/ref/spec)
- [公式ドキュメント](https://golang.org/doc/)
- [標準パッケージ](https://golang.org/pkg/)
- [golangパッケージ](https://pkg.go.dev/)
- [goのロードマップ](https://github.com/Alikhll/golang-developer-roadmap/blob/master/i18n/ja-JP/ReadMe-ja-JP.md)
## 確認しておきたい内容
- Variables・Interface・Type・Strcutの基礎知識
    - sturctのポインタ型の、フィールドへのアクセスに関して、struct.Fieldという形式でアクセス可能なのは、Goコンパイラが、利便性のために暗黙的変換を行っているためである。実際には、(*struct).Fieladと解釈される。
    - 要確認: array・structは値型。Slice・Mapは参照型であることに注意
- Defer
- Pointer
- Goroutines
    - mutex
- Channels
- 標準パッケージ
    - net
        - http
        - mail
    - os
    - reflect...ORMapperとの関連性がありそう
    - io
    - log
    - database/sql
    - errors
    - time