package app

import (
	"github.com/gin-gonic/gin"
	"oauthService/domain/access_token"
	"oauthService/http"
	"oauthService/repository/db"
)

var (
	router = gin.Default()
)

func StartApplication() {

	dbRepository := db.NewRepository()
	atService := access_token.NewService(dbRepository)
	atHandler := http.NewHandler(atService)

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.Run(":8080")

}
