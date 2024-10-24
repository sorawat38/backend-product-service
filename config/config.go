package config

import "github.com/spf13/viper"

type Config struct {
	App      App
	Log      Log
	Database Database
}

type App struct {
	Name string
	Port string
}

type Log struct {
	Env string
}

type Database struct {
	Username string
	Password string
	Host     string
	Port     int
	DBName   string
}

func InitConfig() (Config, error) {

	viper.SetConfigName("config")  // name of config file (without extension)
	viper.SetConfigType("yaml")    // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("config/") // optionally look for config in the working directory
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return Config{}, err
		}
	}

	var c Config
	err := viper.Unmarshal(&c)
	if err != nil {
		return Config{}, err
	}

	return c, nil
}
