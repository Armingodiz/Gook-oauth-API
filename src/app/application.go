package app

import (
	"github.com/ArminGodiz/Gook-oauth-API/src/http"
	"github.com/ArminGodiz/Gook-oauth-API/src/repository/db"
	"github.com/ArminGodiz/Gook-oauth-API/src/repository/rest"
	"github.com/ArminGodiz/Gook-oauth-API/src/services"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	atHandler := http.NewHandler(services.NewService(db.NewRepository(),rest.NewRepository()))
	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.POST("/oauth/access_token", atHandler.Create)
	router.Run(":2222")
}
