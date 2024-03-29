package modelrepository

//go:generate mockgen -destinition=mock_$GOFILE -package=$GOPACKAGE

import (
    "time"
)

type {{toUpperCase .Name}} struct {
    {{- range $key, $element := .Fields}}
    {{$element.Name}} {{$element.Type}}
    {{- end}}
}

type {{toUpperCase .Name}}PK struct {
    {{- range $index, $element := .Fields}}
    {{- if $element.IsPK}}
    {{$element.Name}} {{$element.Type}}
    {{- end}}
    {{- end}}
}