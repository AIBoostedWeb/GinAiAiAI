package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type InitConfigError struct {
	initError error
}

func (err *InitConfigError) Error() string {
	return fmt.Sprintf("InitConfigError: %s", err.initError.Error())
}

func InitConf(pathToConf string, configName string, configType string) error {
	viper.AddConfigPath(pathToConf)
	viper.SetConfigName(configName)
	viper.SetConfigType(configType)

	var err error
	err = loadConfig()
	if err != nil {
		newErr := InitConfigError{err}
		return &newErr
	}

	err = viper.Unmarshal(&CONFIG)
	if err != nil {
		return err
	}

	return nil

}

func loadConfig() error {
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	return nil
}
