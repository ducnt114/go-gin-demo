package controllers

import (
	"net/http"

	"github.com/ducnt114/go-gin-demo/dtos"
	"github.com/ducnt114/go-gin-demo/services"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AuthController struct {
	BaseController
}

func NewAuthController(sp services.ServiceProvider) *AuthController {
	c := &AuthController{}
	c.serviceProvider = sp
	return c
}

func (c *AuthController) Login(ctx *gin.Context) {
	var request *dtos.LoginRequest
	if err := ctx.Bind(&request); err != nil {
		zap.S().Error(ctx, "invalid format request", err)
		c.buildErrorResponse(ctx, http.StatusBadRequest, &dtos.LoginResponse{
			Meta: dtos.BadRequestMeta,
		})
		return
	}

	response := c.serviceProvider.GetAuthService().Login(request.Username, request.Password)

	c.buildResponse(ctx, response.Meta.Code, response)
}

func (c *AuthController) RefreshToken(ctx *gin.Context) {
	var request *dtos.RefreshTokenReq
	if err := ctx.Bind(&request); err != nil {
		zap.S().Error(ctx, "invalid format request", err)
		c.buildErrorResponse(ctx, http.StatusBadRequest, &dtos.LoginResponse{
			Meta: dtos.BadRequestMeta,
		})
		return
	}

	response := c.serviceProvider.GetAuthService().RefreshToken(request.RefreshToken)

	c.buildResponse(ctx, response.Meta.Code, response)
}
