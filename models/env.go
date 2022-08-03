package models

type Config struct {
	App struct {
		Server struct {
			Port string `mapstructure:"port"`
			Host string `mapstructure:"host"`
		} `mapstructure:"server"`
	} `mapstructure:"app"`
}
