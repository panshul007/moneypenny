package main

import (
	"github.com/panshul007/moneypenny/cmd"
	"go.uber.org/zap"
)

func main() {
	//loggerConfig := zap.NewDevelopmentConfig()
	//loggerConfig.OutputPaths = []string{"stdout"}
	logger, _ := zap.NewDevelopment()
	zap.ReplaceGlobals(logger)
	defer logger.Sync()
	logger.Info("Welcome to MoneyPenny!")
	cmd.Execute()
}
