package config

type ServerConfig struct {
	Port string `mapstructure:"port"`
	Host string `mapstructure:"host"`
}

type DatabaseConfig struct {
	Username        string `mapstructure:"username"`
	Password        string `mapstructure:"password"`
	Host            string `mapstructure:"host"`
	Port            string `mapstructure:"port"`
	Name            string `mapstructure:"name"`
	Charset         string `mapstructure:"charset"`
	ParseTime       bool   `mapstructure:"parseTime"`
	MaxOpenConns    int    `mapstructure:"maxOpenConns"`
	MaxIdleConns    int    `mapstructure:"maxIdleConns"`
	ConnMaxLifetime string `mapstructure:"connMaxLifetime"`
}

type GlobalConfig struct {
	RunMode string `mapstructure:"run_mode"`
}

type ServiceConfig struct {
	JWTSecret string `mapstructure:"jwt_secret"`
}
type Config struct {
	Global   GlobalConfig   `mapstructure:"global"`
	Server   ServerConfig   `mapstructure:"server"`
	Database DatabaseConfig `mapstructure:"database"`
	Service  ServiceConfig  `mapstructure:"service"`
}
