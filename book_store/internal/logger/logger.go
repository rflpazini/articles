package logger

import (
	"book-store/pkg/utils"
	"go.uber.org/zap"
)

func NewZapLogger() (*zap.Logger, error) {
	env := utils.GetEnv()
	var lg *zap.Logger
	if env == utils.DEV {
		lg, _ = zap.NewDevelopment()
	} else {
		lg, _ = zap.NewProduction()
	}

	return lg, nil
}
