package config

import (
	"github.com/spf13/viper"
	"log"
)

type Env struct {
	DBUsername  string `mapstructure:"DATABASE_USERNAME"`
	DBPassword  string `mapstructure:"DATABASE_PASSWORD"`
	DBHost      string `mapstructure:"DATABASE_HOST"`
	DBPort      string `mapstructure:"DATABASE_PORT"`
	DBName      string `mapstructure:"DATABASE_NAME"`
	SecretToken string `mapstructure:"SECRET_TOKEN"`
}

func NewEnv(path string) (*Env, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	// Muat konfigurasi (baik dari variabel lingkungan atau file)
	err := viper.ReadInConfig() // Pakai hanya jika konfigurasi dari file diperlukan
	if err != nil {
		log.Print(err)
		return nil, err
	}

	var env Env
	err = viper.Unmarshal(&env)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	return &env, nil
}
