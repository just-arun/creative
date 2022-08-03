package models

type Config struct {
	App struct {
		Server struct {
			Port string `mapstructure:"port"`
		} `mapstructure:"server"`
	} `mapstructure:"app"`
}
