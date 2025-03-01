package config

type ServerConfig struct {
	Port string `mapstructure:"port"`
}

type DatabaseConfig struct {
	DSN string `mapstructure:"dsn"`
}

type GlobalConfig struct {
	RunMode string `mapstructure:"run_mode"`
}

type Config struct {
	Global   GlobalConfig   `mapstructure:"global"`
	Server   ServerConfig   `mapstructure:"server"`
	Database DatabaseConfig `mapstructure:"database"`
}

var CONFIG Config
