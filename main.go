package main

import (
	"fmt"

	"github.com/ducnt114/go-gin-demo/conf"
	"github.com/ducnt114/go-gin-demo/logger"
	"github.com/ducnt114/go-gin-demo/routers"

	"go.uber.org/zap"
)

func init() {
	logger.InitLogger()
	conf.InitConfig()
}

func main() {
	bindingAddr := fmt.Sprintf(":%v", conf.EnvConfig.Server.BindingPort)
	zap.S().Info("Starting api at: ", bindingAddr)
	router, err := routers.InitRouter()
	if err != nil {
		zap.S().Error("Error when init router, detail: ", err)
		panic(err)
	}
	_ = router.Run(bindingAddr)
}
