package models

type Config struct {
	App struct {
		Server struct {
			Port string `mapstructure:"port"`
			Host string `mapstructure:"host"`
		} `mapstructure:"server"`
		Redis struct {
			Password string `mapstructure:"password"`
			Host string `mapstructure:"host"`
		} `mapstructure:"redis"`
	} `mapstructure:"app"`
}
