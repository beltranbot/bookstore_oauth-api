package db

import (
	"github.com/beltranbot/bookstore_oauth-api/clients/cassandra"
	"github.com/beltranbot/bookstore_oauth-api/domain/accesstoken"
	"github.com/beltranbot/bookstore_oauth-api/utils/errors"
)

// Repository interface
type Repository interface {
	GetByID(string) (*accesstoken.AccessToken, *errors.RestErr)
}

type repository struct {
}

// NewRepository func
func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetByID(id string) (*accesstoken.AccessToken, *errors.RestErr) {
	session, err := cassandra.GetSession()
	if err != nil {
		panic(err)
	}
	defer session.Close()
	return nil, errors.NewInternalServerError("database connection not implemente yet")
}
