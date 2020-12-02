package accesstoken

import (
	"fmt"
	"strings"
	"time"

	"github.com/beltranbot/bookstore_oauth-api/utils/cryptoutils"
	"github.com/beltranbot/bookstore_oauth-api/utils/errors"
)

const (
	expirationTime             = 24
	grantTypePassword          = "password"
	grandTypeClientCredentials = "client_credentials"
)

// AccessToken struct
type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserID      int64  `json:"user_id"`
	ClientID    int64  `json:"client_id"`
	Expires     int64  `json:"expires"`
}

// Request type
type Request struct {
	GrantType string `json:"grant_type"`
	Scope     string `json:"scope"`

	// Used for password grant type
	Username string `json:"username"`
	Password string `json:"password"`

	// Used for client_credentials grant type
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

// Validate func, validates Request
func (request *Request) Validate() *errors.RestErr {
	switch request.GrantType {
	case grantTypePassword:
		break

	case grandTypeClientCredentials:
		break

	default:
		return errors.NewBadRequestError("invalid grant_type parameter")
	}

	//TODO: Validate parameters for each grant_type
	return nil
}

// GetNewAccessToken func
func GetNewAccessToken(userID int64) *AccessToken {
	return &AccessToken{
		UserID:  userID,
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

// Generate func
func (at *AccessToken) Generate() {
	at.AccessToken = cryptoutils.GetMD5(fmt.Sprintf("at-%d-%d-ran", at.UserID, at.Expires))
}
