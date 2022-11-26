package main

import (
	"fmt"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"log"
)

const src = `package main
type IFRepository interface {
	Get() int
}

type Repository struct {}

func (r *Repository) Get() int {
	return 1
}

type Service struct {
	repo IFRepository
}

func NewService(r *Repository) *Service {
	return &Service{repo: r}
}

func main() {
	s := NewService(nil)
	print(s)
}`

func main() {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "main.go", src, 0)
	if err != nil {
		/* エラー処理 */
		log.Fatal(err)
	}
	cfg := &types.Config{Importer: importer.Default()}
	info := &types.Info{
		Types: make(map[ast.Expr]types.TypeAndValue),
	}
	pkg, err := cfg.Check("main", fset, []*ast.File{file}, info)
	if err != nil {
		/* エラー処理 */
		log.Fatal(err)
	}
	fmt.Println("package is", pkg.Path())
	ul := make(map[string]types.Type)
	el := make(map[*types.Struct]ast.Expr)
	for e, t := range info.Types {
		// Collect the underlying types.
		ul[t.Type.String()] = t.Type.Underlying()
		// fmt.Printf("type=%+v, underlying=%+v\n", t.Type.String(), t.Type.Underlying())
		// Collect structs to determine the fields of a receiver.
		if v, ok := t.Type.(*types.Struct); ok {
			el[v] = e
			if v.NumFields() == 0 {
				continue
			}
			fmt.Printf("struct=%+v, e=%+v\n", v.Field(0).Type().Underlying(), e)
		}
		fmt.Println("-----------------")
	}
}
