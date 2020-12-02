package rest

import (
	"encoding/json"
	"time"

	"github.com/beltranbot/bookstore_oauth-api/domain/users"
	"github.com/beltranbot/bookstore_oauth-api/utils/errors"
	"github.com/federicoleon/golang-restclient/rest"
)

var (
	restClient = rest.RequestBuilder{
		BaseURL: "http://localhost:8080", // http is important here for testing purposes
		Timeout: 100 * time.Millisecond,
	}
)

// UserRepository type interface
type UserRepository interface {
	LoginUser(string, string) (*users.User, *errors.RestErr)
}

type usersRepository struct{}

// NewRepository func
func NewRepository() UserRepository {
	return &usersRepository{}
}

func (ur *usersRepository) LoginUser(email string, password string) (*users.User, *errors.RestErr) {
	request := users.UserLoginRequest{
		Email:    email,
		Password: password,
	}
	response := restClient.Post("/users/login", request)

	if response == nil || response.Response == nil {
		return nil, errors.NewInternalServerError("invalid restclient response when trying to login user")
	}

	if response.StatusCode > 299 {
		apiErr, err := errors.NewRestErrorFromBytes(response.Bytes())
		if err != nil {
			return nil, errors.NewInternalServerError("invalid error interface when trying to login user")
		}
		return nil, apiErr
	}

	var user users.User
	if err := json.Unmarshal(response.Bytes(), &user); err != nil {
		return nil, errors.NewInternalServerError("error when trying to unmarshal users login response")
	}
	return &user, nil
}
