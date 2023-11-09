package main

/*
go build -ldflags "\
-X main.version=$(git describe --tag --abbrev=0) \
-X main.revision=$(git rev-list -1 HEAD) \
-X main.build=$(git describe --tags) \
" flags/ldflags/main.go
*/
var (
	version  string
	revision string
	build    string
)

func main() {
	println("version:", version)
	println("revision:", revision)
	println("build:", build)
}
