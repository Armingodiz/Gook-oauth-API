package http

import (
	"github.com/ArminGodiz/Gook-oauth-API/src/domain/access_token"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AccessTokenHandler interface {
	GetById(ctx *gin.Context)
}
type accessTokenHandler struct {
	service access_token.Service
}

func NewHandler(service access_token.Service) AccessTokenHandler {
	return &accessTokenHandler{
		service: service,
	}
}

func (ath *accessTokenHandler) GetById(c *gin.Context) {
	accessTokenId := c.Param("access_token_id")
	accessToken, err := ath.service.GetById(accessTokenId)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}
	c.JSON(http.StatusOK, accessToken)
}

func (ath *accessTokenHandler) Create(c *gin.Context) {
	var at access_token.AccessToken
	err := c.ShouldBindJSON(&at)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	err2 := ath.service.Create(at)
	if err2 != nil {
		c.JSON(err2.Code, err2)
		return
	}
	c.JSON(http.StatusCreated, at)
}
