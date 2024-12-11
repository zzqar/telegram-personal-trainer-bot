package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
)

type Config struct {
	Env     string `env:"ENV" env-default:"prod"`
	Token   string `env:"TELEGRAM_TOKEN" env-required:"true"`
	AdminId int64  `env:"ADMIN_ID" env-required:"true"`
	Redis   Redis  `env-prefix:"REDIS_" env-required:"true"`
}

type Redis struct {
	Addr     string `env:"ADDR" env-default:"localhost:6379"`
	Password string `env:"PASSWORD" env-default:""`
	DB       int    `env:"DB" env-default:"0"`
}

func MustConfig() *Config {
	var cfg Config
	if err := cleanenv.ReadConfig(".env", &cfg); err != nil {
		log.Fatalf("Error reading config: %v", err)
	}
	return &cfg
}
