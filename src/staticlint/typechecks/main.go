package main

import (
	"fmt"
	"go/types"
	"log"

	"golang.org/x/tools/go/packages"
)

func main() {
	checkWithPackages()
	// checkFromOtherPackage()
	// checkFromOneFile()
}

func checkWithPackages() {
	cfg := &packages.Config{
		Mode: packages.NeedImports | packages.NeedTypes,
		// NeedTypesInfo
	}
	pkzs, err := packages.Load(cfg, "target/target.go")
	if err != nil {
		log.Fatal(err)
	}

	for k, v := range pkzs[0].Types.Scope().Names() {
		fmt.Printf("k=%+v, v=%+v\n", k, v)
	}
	obj := pkzs[0].Types.Scope().Lookup("BlogService")
	tp := obj.Type().Underlying()

	st := tp.(*types.Struct)
	for i := 0; i < st.NumFields(); i++ {
		fmt.Println(st.Field(i).Type().Underlying())
	}
}

// func checkFromOtherPackage() {
// 	fset := token.NewFileSet()
// 	file, err := parser.ParseFile(fset, "/home/kazdevl/go/src/github.com/kazdevl/presentation_materials/20221202/target/target.go", nil, 0)
// 	// loaderConf := &loader.Config{}
// 	// file, err := loaderConf.ParseFile("target/target.go", nil)
// 	if err != nil {
// 		/* エラー処理 */
// 		log.Fatal(err)
// 	}
// 	cfg := &types.Config{
// 		Importer: importer.Default(),
// 		Error:    func(error) {},
// 	}
// 	info := &types.Info{
// 		Types: make(map[ast.Expr]types.TypeAndValue),
// 	}
// 	cfg.Check("", fset, []*ast.File{file}, info)
// 	for e, t := range info.Types {
// 		// Collect the underlying types.
// 		fmt.Printf("type=%+v, underlying=%+v\n", t.Type.String(), t.Type.Underlying())
// 		// Collect structs to determine the fields of a receiver.
// 		if v, ok := t.Type.(*types.Struct); ok {
// 			if v.NumFields() == 0 {
// 				continue
// 			}
// 			fmt.Printf("struct=%+v, e=%+v\n", v.Field(0).Type().Underlying(), e)
// 		}
// 		fmt.Println("-----------------")
// 	}
// }

// func checkFromOneFile() {
// 	fset := token.NewFileSet()
// 	file, err := parser.ParseFile(fset, "main.go", src, 0)
// 	if err != nil {
// 		/* エラー処理 */
// 		log.Fatal(err)
// 	}
// 	cfg := &types.Config{Importer: importer.Default()}
// 	info := &types.Info{
// 		Types: make(map[ast.Expr]types.TypeAndValue),
// 	}
// 	pkg, err := cfg.Check("main", fset, []*ast.File{file}, info)
// 	if err != nil {
// 		/* エラー処理 */
// 		log.Fatal(err)
// 	}
// 	fmt.Println("package is", pkg.Path())
// 	for e, t := range info.Types {
// 		// Collect the underlying types.
// 		fmt.Printf("type=%+v, underlying=%+v\n", t.Type.String(), t.Type.Underlying())
// 		// Collect structs to determine the fields of a receiver.
// 		if v, ok := t.Type.(*types.Struct); ok {
// 			if v.NumFields() == 0 {
// 				continue
// 			}
// 			fmt.Printf("struct=%+v, e=%+v\n", v.Field(0).Type().Underlying(), e)
// 		}
// 		fmt.Println("-----------------")
// 	}
// }

// const src = `package main
// type IFRepository interface {
// 	Get() int
// }

// type Repository struct {}

// func (r *Repository) Get() int {
// 	return 1
// }

// type Service struct {
// 	repo IFRepository
// }

// func NewService(r *Repository) *Service {
// 	return &Service{repo: r}
// }

// func main() {
// 	s := NewService(nil)
// 	print(s)
// }`
