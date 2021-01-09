package app

import (
	"bookstore/bookstore-oauth-api/src/domain/access_token"
	"bookstore/bookstore-oauth-api/src/http"
	"bookstore/bookstore-oauth-api/src/repository/db"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApp() {
	dbRepository := db.NewRepository()
	atService := access_token.NewService(dbRepository)
	atHandler := http.NewHandler(atService)

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetByID)

	router.Run(":8080")
}
