package router

import (
	"fmt"
	"time"

	"github.com/ducnt114/go-gin-demo/repositories"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/rs/zerolog/log"
)

func InitRouter() (*gin.Engine, error) {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	serviceProvider, err := createServiceProvider()
	if err != nil {
		log.Error("Error when createServiceProvider, detail: ", err)
		return nil, err
	}

	categoryController := controllers.NewCategoryController(serviceProvider)
	itemController := controllers.NewItemController(serviceProvider)
	itemControllerV2 := controllers.NewItemControllerV2(serviceProvider)
	defaultController := controllers.DefaultController{}

	v1 := r.Group("/auth/v1")
	{
		defaultHandler := v1.Group("/health")
		{
			defaultHandler.GET("/", defaultController.Get)
		}

	}

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
	serviceProvider := services.NewServiceProvider(repoProvider)

	return serviceProvider, nil
}
