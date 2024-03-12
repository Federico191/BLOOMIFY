package config

import (
	"github.com/spf13/viper"
	"log"
)

type Env struct {
	APort          string `mapstructure:"APP_PORT"`
	AHost          string `mapstructure:"APP_HOST"`
	DBUsername     string `mapstructure:"DATABASE_USERNAME"`
	DBPassword     string `mapstructure:"DATABASE_PASSWORD"`
	DBHost         string `mapstructure:"DATABASE_HOST"`
	DBPort         string `mapstructure:"DATABASE_PORT"`
	DBName         string `mapstructure:"DATABASE_NAME"`
	SecretToken    string `mapstructure:"SECRET_TOKEN"`
	ExpiredToken   string `mapstructure:"EXPIRED_TOKEN"`
	EmailFrom      string `mapstructure:"EMAIL_FROM"`
	SMTPHost       string `mapstructure:"SMTP_HOST"`
	SMTPPort       string `mapstructure:"SMTP_PORT"`
	SMTPUser       string `mapstructure:"SMTP_USER"`
	SMTPPassword   string `mapstructure:"SMTP_PASSWORD"`
	ServerKey      string `mapstructure:"SERVER_KEY"`
	SupabaseUrl    string `mapstructure:"SUPABASE_URL"`
	SupabaseKey    string `mapstructure:"SUPABASE_KEY"`
	SupabaseBucket string `mapstructure:"SUPABASE_BUCKET"`
}

func NewEnv(path string) (*Env, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
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
