package main

import (
	"log"
	"os"
	"text/template"

	"github.com/jmoiron/sqlx"
)

type Sample struct {
	A string
	B int
}

func main() {
	sample := Sample{"Hello", 2}
	tmpl, err := template.New("sample").Parse("Aの背丈は{{.A}}. B is {{.B}}")
	if err != nil {
		log.Fatal(err)
	}
	if err := tmpl.Execute(os.Stdout, sample); err != nil {
		log.Fatal(err)
	}
}

// template内容
// 全PJで共通しているものにする
// modelとそのメソッド
//
// go generateでmockgenのやつも

type Human struct {
	ID     int64  `custom:"PK"`
	Name   string `cutom:"SK"`
	Height int
}

type Accessory struct {
	Name    string `custom:"PK"`
	HumanID int64  `custom:"PK"`
}

// 以下は自動生成して欲しいやつ
// 使うやつ, text・spaces・if・range・functions(toLowerCamelCase・tagの取得)・associated template

// go:generate mockgen -sorce=humanmodelrepository.go -destinition
func (h *Human) toHumanPK() HumanPK {
	return HumanPK{ID: h.ID}
}

type HumanPK struct {
	ID int64
}

type Humans []*Human

func (hs *Humans) ToMap() map[HumanPK]*Human {
	m := make(map[HumanPK]*Human, len(*hs))
	for _, h := range *hs {
		m[h.toHumanPK()] = h
	}
	return m
}

type HumanModelRepository struct {
	client *sqlx.DB
}

func (r *HumanModelRepository) Get(hk HumanPK) (*Human, error) {
	h := new(Human)
	if err := r.client.Select(&h, "SELECT * FROM human WHERE id=?", hk.ID); err != nil {
		return nil, err
	}
	return h, nil
}

// SKありの場合
func (r *HumanModelRepository) FindByName(name string) (Humans, error) {
	var hs Humans
	if err := r.client.Select(&hs, "SELECT * FROM human WHERE name=?", name); err != nil {
		return nil, err
	}
	return hs, nil
}

// Name・toLowerCase・ShortName・isPK・Fields・subOne
