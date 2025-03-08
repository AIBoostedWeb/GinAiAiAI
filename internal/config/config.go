package config

type ServerConfig struct {
	Port string   `mapstructure:"port"`
	Host string   `mapstructure:"host"`
	AI   AIConfig `mapstructure:"ai"`
}

type DatabaseConfig struct {
	Username        string `mapstructure:"username"`
	Password        string `mapstructure:"password"`
	Host            string `mapstructure:"host"`
	Port            string `mapstructure:"port"`
	Name            string `mapstructure:"name"`
	Charset         string `mapstructure:"charset"`
	ParseTime       bool   `mapstructure:"parse_time"`
	MaxOpenConns    int    `mapstructure:"max_open_conns"`
	MaxIdleConns    int    `mapstructure:"max_idle_conns"`
	ConnMaxLifetime string `mapstructure:"conn_max_lifetime"`
}

type GlobalConfig struct {
	RunMode string `mapstructure:"run_mode"`
}

type ServiceConfig struct {
	JWTSecret string `mapstructure:"jwt_secret"`
}
type AIConfig struct {
	AuthKey string `mapstructure:"auth_key"`
	BaseURL string `mapstructure:"base_url"`
	Model   string `mapstructure:"model"`
}
type Config struct {
	Global   GlobalConfig   `mapstructure:"global"`
	Server   ServerConfig   `mapstructure:"server"`
	Database DatabaseConfig `mapstructure:"database"`
	Service  ServiceConfig  `mapstructure:"service"`
}
