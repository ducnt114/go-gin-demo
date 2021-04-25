package routers

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/ducnt114/go-gin-demo/conf"
	"github.com/ducnt114/go-gin-demo/controllers"
	"github.com/ducnt114/go-gin-demo/dtos"
	"github.com/ducnt114/go-gin-demo/repositories"
	"github.com/ducnt114/go-gin-demo/services"
	"github.com/ducnt114/go-gin-demo/utils"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

	"go.uber.org/zap"
)

const (
	TokenTypeJWTAuthen = "Bearer"
)

var jwtHelper utils.JWTHelper

func InitRouter() (*gin.Engine, error) {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	serviceProvider, err := createServiceProvider()
	if err != nil {
		zap.S().Error("Error when createServiceProvider, detail: ", err)
		return nil, err
	}

	authController := controllers.NewAuthController(serviceProvider)
	userController := controllers.NewUserController(serviceProvider)

	v1 := r.Group("/auth/v1")

	v1.POST("/login", authController.Login)
	v1.POST("/refresh-token", authController.RefreshToken)
	v1.Use(middlewareJWTAuthen()).GET("/info", userController.GetInfo)

	return r, nil
}

func createServiceProvider() (services.ServiceProvider, error) {
	dborm, err := gorm.Open("mysql", fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True",
		conf.EnvConfig.MySQL.User, conf.EnvConfig.MySQL.Password,
		conf.EnvConfig.MySQL.Host, conf.EnvConfig.MySQL.Port, conf.EnvConfig.MySQL.DB))
	if err != nil {
		return nil, err
	}

	dborm.DB().SetConnMaxLifetime(5 * time.Minute)

	if conf.EnvConfig.Environment == conf.EnvironmentLocal {
		dborm.LogMode(true)
	}

	repoProvider, err := repositories.NewRepositoryProvider(dborm)
	if err != nil {
		return nil, err
	}

	jHelper, err := utils.NewJWTHelper(conf.EnvConfig.JWT.PublicKey, conf.EnvConfig.JWT.PrivateKey)
	if err != nil {
		return nil, err
	}
	jwtHelper = jHelper

	serviceProvider := services.NewServiceProvider(repoProvider, jHelper)

	return serviceProvider, nil
}

func middlewareJWTAuthen() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		tokens := strings.Split(token, " ")
		if len(tokens) != 2 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, &dtos.BaseResponse{
				Meta: &dtos.Meta{
					Code:    http.StatusUnauthorized,
					Message: "Unauthorized"}})
			return
		}

		switch tokens[0] {
		case TokenTypeJWTAuthen:
			var userClaim dtos.AuthClaims
			err := jwtHelper.ParseClaims(tokens[1], &userClaim)
			if err != nil {
				zap.S().Error("Error when parse jwt token, detail: ", err)
				c.AbortWithStatusJSON(http.StatusUnauthorized, &dtos.BaseResponse{
					Meta: &dtos.Meta{
						Code:    http.StatusUnauthorized,
						Message: "Unauthorized"}})
				return
			}
			c.Set("user_id", userClaim.UserID)
		default:
			c.AbortWithStatusJSON(http.StatusUnauthorized, &dtos.BaseResponse{
				Meta: &dtos.Meta{
					Code:    http.StatusUnauthorized,
					Message: "Unauthorized"}})
			return
		}
		c.Next()
	}
}
