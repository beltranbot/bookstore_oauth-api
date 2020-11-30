package app

import (
	"github.com/beltranbot/bookstore_oauth-api/app/http"
	"github.com/beltranbot/bookstore_oauth-api/clients/cassandra"
	"github.com/beltranbot/bookstore_oauth-api/domain/accesstoken"
	"github.com/beltranbot/bookstore_oauth-api/repository/db"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

// StartApplication func
func StartApplication() {
	session := cassandra.GetSession()
	defer session.Close()

	accessTokenHandler := http.NewHandler(accesstoken.NewService(db.NewRepository()))

	router.GET("/oauth/token/:access_token_id", accessTokenHandler.GetByID)
	router.POST("/oauth/token", accessTokenHandler.Create)

	router.Run(":8080")
}
