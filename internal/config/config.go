package config

import (
	"github.com/spf13/viper"
)

const (
	ConfigName string = "app"
	ConfigType string = "env"
)

type Config struct {
	DSN  string `mapstructure:"DSN"`
	Host string `mapstructure:"Host"`
	Port string `mapstructure:"Port"`
}

// Reading config file for setting application
func LoadConfig(path string) (Config, error) {

	conf := Config{}

	viper.AddConfigPath(path)
	viper.SetConfigName(ConfigName)
	viper.SetConfigType(ConfigType)
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return conf, err
	}

	err = viper.Unmarshal(&conf)
	if err != nil {
		return conf, err
	}

	return conf, nil
}
