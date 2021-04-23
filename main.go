package main

import (
	"go.uber.org/zap"
)

func init() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	undo := zap.ReplaceGlobals(logger)
	defer undo()

	zap.L().Info("init")
}

func main() {
	zap.L().Info("replaced zap's global loggers")
}
