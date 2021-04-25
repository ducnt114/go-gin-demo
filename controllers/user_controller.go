package controllers

import (
	"net/http"

	"github.com/ducnt114/go-gin-demo/dtos"
	"github.com/ducnt114/go-gin-demo/services"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserController struct {
	BaseController
}

func NewUserController(sp services.ServiceProvider) *UserController {
	c := &UserController{}
	c.serviceProvider = sp
	return c
}

func (c *UserController) GetInfo(ctx *gin.Context) {
	userID, exist := ctx.Get("user_id")
	if !exist {
		zap.S().Error("error when get user_id from token")
		c.buildErrorResponse(ctx, http.StatusBadRequest, &dtos.GetUserInfoResponse{
			Meta: dtos.BadRequestMeta,
		})
		return
	}
	res := c.serviceProvider.GetUserService().GetUserInfo(userID.(uint))
	c.buildResponse(ctx, res.Meta.Code, res)
}
