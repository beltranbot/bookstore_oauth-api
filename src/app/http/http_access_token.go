package http

import (
	"net/http"

	"github.com/beltranbot/bookstore_oauth-api/domain/accesstoken"
	"github.com/gin-gonic/gin"
)

// AccessTokenHandler interface
type AccessTokenHandler interface {
	GetByID(*gin.Context)
}

type accessTokenHandler struct {
	service accesstoken.Service
}

// NewHandler func
func NewHandler(service accesstoken.Service) AccessTokenHandler {
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
