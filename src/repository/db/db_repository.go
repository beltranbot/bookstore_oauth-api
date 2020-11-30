package db

import (
	"github.com/beltranbot/bookstore_oauth-api/clients/cassandra"
	"github.com/beltranbot/bookstore_oauth-api/domain/accesstoken"
	"github.com/beltranbot/bookstore_oauth-api/utils/errors"
	"github.com/gocql/gocql"
)

const (
	queryGetAccessToken    = "SELECT access_token, user_id, client_id, expires FROM access_tokens WHERE access_token = ?;"
	queryCreateAccessToken = "INSERT INTO access_tokens(access_token, user_id, client_id, expires) VALUES (?, ?, ?, ?);"
	queryUpdateExpires     = "UPDATE access_tokens set expires = ? where access_token = ?;"
)

// Repository interface
type Repository interface {
	GetByID(string) (*accesstoken.AccessToken, *errors.RestErr)
	Create(accesstoken.AccessToken) *errors.RestErr
	UpdateExpirationTime(accesstoken.AccessToken) *errors.RestErr
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
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer session.Close()
	var result accesstoken.AccessToken
	if err := session.Query(queryGetAccessToken, id).Scan(
		&result.AccessToken,
		&result.UserID,
		&result.ClientID,
		&result.Expires,
	); err != nil {
		if err == gocql.ErrNotFound {
			return nil, errors.NewNotFoundError("no access token found with given id")
		}
		return nil, errors.NewInternalServerError(err.Error())
	}
	return &result, nil
}

func (r *repository) Create(at accesstoken.AccessToken) *errors.RestErr {
	session, err := cassandra.GetSession()
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer session.Close()

	if err := session.Query(queryCreateAccessToken,
		at.AccessToken,
		at.UserID,
		at.ClientID,
		at.Expires,
	).Exec(); err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	return nil
}

func (r *repository) UpdateExpirationTime(at accesstoken.AccessToken) *errors.RestErr {
	session, err := cassandra.GetSession()
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer session.Close()

	if err := session.Query(queryUpdateExpires,
		at.Expires,
		at.AccessToken,
	).Exec(); err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	return nil
}
