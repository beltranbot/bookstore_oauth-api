package accesstoken

import (
	"strings"
	"time"

	"github.com/beltranbot/bookstore_oauth-api/utils/errors"
)

const (
	expirationTime = 24
)

// AccessToken struct
type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserID      int64  `json:"user_id"`
	ClientID    int64  `json:"client_id"`
	Expires     int64  `json:"expires"`
}

// GetNewAccessToken func
func GetNewAccessToken() *AccessToken {
	return &AccessToken{
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

// IsExpired func
func (at AccessToken) IsExpired() bool {
	return time.Unix(at.Expires, 0).Before(time.Now().UTC())
}

// Validate func
func (at *AccessToken) Validate() *errors.RestErr {
	at.AccessToken = strings.TrimSpace(at.AccessToken)
	if at.AccessToken == "" {
		return errors.NewBadRequestError("invalid access token id")
	}
	if at.UserID <= 0 {
		return errors.NewBadRequestError("invalid user id")
	}
	if at.ClientID <= 0 {
		return errors.NewBadRequestError("invalid client id")
	}
	if at.Expires <= 0 {
		return errors.NewBadRequestError("invalid expiration time")
	}
	return nil
}
