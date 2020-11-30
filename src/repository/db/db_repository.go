package db

import (
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
	return nil, errors.NewInternalServerError("database connection not implemented yet")
}
