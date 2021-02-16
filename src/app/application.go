package app

import (
	"github.com/ArminGodiz/Gook-oauth-API/src/domain/access_token"
	"github.com/ArminGodiz/Gook-oauth-API/src/http"
	"github.com/ArminGodiz/Gook-oauth-API/src/repository/db"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	atHandler := http.NewHandler(access_token.NewService(db.NewRepository()))
	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.Run(":8080")
}
