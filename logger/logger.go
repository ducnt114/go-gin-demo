package logger

import "go.uber.org/zap"

func InitLogger() {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	logger.WithOptions(zap.AddCallerSkip(1))
	zap.ReplaceGlobals(logger)
}
