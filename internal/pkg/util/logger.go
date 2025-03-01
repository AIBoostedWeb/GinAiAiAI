package util

import (
	"fmt"
	"go.uber.org/zap"
)

func InitLogger(stage string) (*zap.Logger, error) {
	var logger *zap.Logger
	if stage == "debug" {
		var err error
		logger, err = zap.NewDevelopment()
		return logger, err
	}
	if stage == "release" {
		var err error
		logger, err = zap.NewDevelopment()
		return logger, err
	}
	return nil, fmt.Errorf("invalid stage " + stage)
}
