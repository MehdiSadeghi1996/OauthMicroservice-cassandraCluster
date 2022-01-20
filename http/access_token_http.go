package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"oauthService/domain/access_token"
	"oauthService/utils/errors"
	"strings"
)

type AccessTokenHandler interface {
	GetById(c *gin.Context)
	Create(c *gin.Context)
}
type accessTokenHandler struct {
	service access_token.Service
}

func NewHandler(service access_token.Service) AccessTokenHandler {
	return &accessTokenHandler{
		service: service,
	}
}

func (handler *accessTokenHandler) GetById(c *gin.Context) {
	accessTokenId := strings.TrimSpace(c.Param("access_token_id"))
	accessToken, err := handler.service.GetById(accessTokenId)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, accessToken)
}
func (handler *accessTokenHandler) Create(c *gin.Context) {
	// اینجا باید کاربر لاگین کنه و بعد از احراز هویت توکن جنریت بشه
	//فعلا عملیات کراد توی کاساندرا
	var at access_token.AccessToken
	err := c.ShouldBindJSON(&at)
	if err != nil {
		resErr := errors.NewBadRequestError("invalid json Body")
		c.JSON(resErr.Status, resErr)
		return
	}
	er := handler.service.Create(at)
	if er != nil {
		c.JSON(er.Status, er)
		return
	}
	c.JSON(http.StatusCreated, at)
}
