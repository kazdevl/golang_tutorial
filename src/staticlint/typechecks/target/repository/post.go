package repository

//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAGE

import (
	"github.com/jmoiron/sqlx"
	"github.com/kazdevl/presentation_materials/20221202/target/model"
)

type IFPostModelRepository interface {
	Get(pk model.PostPK) (*model.Post, error)
	FindByUserID(userid string) (model.Posts, error)
}

type PostModelRepository struct {
	client *sqlx.DB
}

func (r *PostModelRepository) Get(pk model.PostPK) (*model.Post, error) {
	model := new(model.Post)
	if err := r.client.Select(&model, `SELECT * FROM post WHERE
    ID=?
    `,
		pk.ID,
	); err != nil {
		return nil, err
	}
	return model, nil
}

func (r *PostModelRepository) FindByUserID(userid string) (model.Posts, error) {
	var models model.Posts
	if err := r.client.Select(&models, "SELECT * FROM post WHERE userid=?", userid); err != nil {
		return nil, err
	}
	return models, nil
}
