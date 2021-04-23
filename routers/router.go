package routers

import (
	"fmt"
	"time"

	"github.com/ducnt114/go-gin-demo/conf"
	"github.com/ducnt114/go-gin-demo/controllers"
	"github.com/ducnt114/go-gin-demo/repositories"
	"github.com/ducnt114/go-gin-demo/services"
	"github.com/ducnt114/go-gin-demo/utils"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

func InitRouter() (*gin.Engine, error) {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	serviceProvider, err := createServiceProvider()
	if err != nil {
		zap.S().Error("Error when createServiceProvider, detail: ", err)
		return nil, err
	}

	authController := controllers.NewAuthController(serviceProvider)

	v1 := r.Group("/auth/v1")

	v1.POST("/login", authController.Login)

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

	jwtHelper, err := utils.NewJWTHelper(conf.EnvConfig.JWT.PublicKey, conf.EnvConfig.JWT.PrivateKey)
	if err != nil {
		return nil, err
	}

	serviceProvider := services.NewServiceProvider(repoProvider, jwtHelper)

	return serviceProvider, nil
}
