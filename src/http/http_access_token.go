package http

import (
	"net/http"

	atDomain "github.com/beltranbot/bookstore_oauth-api/domain/accesstoken"
	services "github.com/beltranbot/bookstore_oauth-api/services/accesstoken"
	"github.com/beltranbot/bookstore_oauth-api/utils/errors"
	"github.com/gin-gonic/gin"
)

// AccessTokenHandler interface
type AccessTokenHandler interface {
	GetByID(*gin.Context)
	Create(*gin.Context)
	// UpdateExpirationTime(*gin.Context)
}

type accessTokenHandler struct {
	service services.Service
}

// NewHandler func
func NewHandler(service services.Service) AccessTokenHandler {
	return &accessTokenHandler{
		service: service,
	}
}

func (handler *accessTokenHandler) GetByID(c *gin.Context) {
	accessToken, err := handler.service.GetByID(c.Param("access_token_id"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, accessToken)
}

func (handler *accessTokenHandler) Create(c *gin.Context) {
	var request atDomain.Request
	if bindingErr := c.ShouldBindJSON(&request); bindingErr != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	accessToken, err := handler.service.Create(request)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusCreated, accessToken)
}
