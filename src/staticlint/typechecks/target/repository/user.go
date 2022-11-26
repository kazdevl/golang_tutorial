package repository

//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAGE

import (
	"github.com/jmoiron/sqlx"
	"github.com/kazdevl/presentation_materials/20221202/target/model"
)

type IFUserModelRepository interface {
	Get(uk model.UserPK) (*model.User, error)
	FindByName(name string) (model.Users, error)
}

type UserModelRepository struct {
	client *sqlx.DB
}

func (r *UserModelRepository) Get(uk model.UserPK) (*model.User, error) {
	model := new(model.User)
	if err := r.client.Select(&model, `SELECT * FROM user WHERE
    ID=?
    `,
		uk.ID,
	); err != nil {
		return nil, err
	}
	return model, nil
}

func (r *UserModelRepository) FindByName(name string) (model.Users, error) {
	var models model.Users
	if err := r.client.Select(&models, "SELECT * FROM user WHERE name=?", name); err != nil {
		return nil, err
	}
	return models, nil
}
