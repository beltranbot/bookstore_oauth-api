package services

import (
	"strings"

	"github.com/beltranbot/bookstore_oauth-api/domain/accesstoken"
	"github.com/beltranbot/bookstore_oauth-api/repository/db"
	"github.com/beltranbot/bookstore_oauth-api/repository/rest"

	"github.com/beltranbot/bookstore_oauth-api/utils/errors"
)

// Service interface
type Service interface {
	GetByID(string) (*accesstoken.AccessToken, *errors.RestErr)
	Create(accesstoken.Request) (*accesstoken.AccessToken, *errors.RestErr)
	UpdateExpirationTime(accesstoken.AccessToken) *errors.RestErr
}

type service struct {
	restUserRepo rest.UserRepository
	DBRepo       db.Repository
}

// NewService func
func NewService(usersRepo rest.UserRepository, DBRepo db.Repository) Service {
	return &service{
		restUserRepo: usersRepo,
		DBRepo:       DBRepo,
	}
}

func (s *service) GetByID(accessTokenID string) (*accesstoken.AccessToken, *errors.RestErr) {
	accessTokenID = strings.TrimSpace(accessTokenID)
	if len(accessTokenID) == 0 {
		return nil, errors.NewBadRequestError("invalid access token id")
	}

	accessToken, err := s.DBRepo.GetByID(accessTokenID)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}

func (s *service) Create(request accesstoken.Request) (*accesstoken.AccessToken, *errors.RestErr) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	// TODO: support both grant types: client_credentials and password

	// Authenticate the user against the Users API
	user, err := s.restUserRepo.LoginUser(request.Username, request.Password)
	if err != nil {
		return nil, err
	}

	// Generate a new access token:
	at := accesstoken.GetNewAccessToken(user.ID)
	at.Generate()

	// Save the new access token in Cassandra:
	if err := s.DBRepo.Create(*at); err != nil {
		return nil, err
	}

	return at, nil
}

func (s *service) UpdateExpirationTime(at accesstoken.AccessToken) *errors.RestErr {
	if err := at.Validate(); err != nil {
		return err
	}
	return s.DBRepo.UpdateExpirationTime(at)
}
