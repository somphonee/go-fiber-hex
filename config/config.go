package config


import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	AppName     string `mapstructure:"APP_NAME"`
	AppEnv      string `mapstructure:"APP_ENV"`
	AppPort     string `mapstructure:"APP_PORT"`
	DBHost      string `mapstructure:"DB_HOST"`
	DBPort      string `mapstructure:"DB_PORT"`
	DBUser      string `mapstructure:"DB_USER"`
	DBPassword  string `mapstructure:"DB_PASSWORD"`
	DBName      string `mapstructure:"DB_NAME"`
}

func LoadConfig() *Config {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("Warning: %v", err)
	}

	config := &Config{
		AppName:     viper.GetString("APP_NAME"),
		AppEnv:      viper.GetString("APP_ENV"),
		AppPort:     viper.GetString("APP_PORT"),
		DBHost:      viper.GetString("DB_HOST"),
		DBPort:      viper.GetString("DB_PORT"),
		DBUser:      viper.GetString("DB_USER"),
		DBPassword:  viper.GetString("DB_PASSWORD"),
		DBName:      viper.GetString("DB_NAME"),
	}

	// ตั้งค่า default values
	if config.AppPort == "" {
		config.AppPort = "8080"
	}
	if config.DBHost == "" {
		config.DBHost = "localhost"
	}
	if config.DBPort == "" {
		config.DBPort = "5434"
	}
	if config.DBUser == "" {
		config.DBUser = "postgres"
	}
	if config.DBPassword == "" {
		config.DBPassword = "postgres"
	}
	if config.DBName == "" {
		config.DBName = "postgres"
	}

	return config
}
