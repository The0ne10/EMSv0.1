package config

import (
	"errors"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
	"time"
)

type Config struct {
	Env        string `yaml:"env" env-default:"local"`
	HTTPServer `yaml:"http_server"`
	Postgres   `yaml:"postgres"`
}

type HTTPServer struct {
	Addr     string        `yaml:"addr" env-default:":8080"`
	Timeout  time.Duration `yaml:"timeout" env-default:"5s"`
	TimeIdle time.Duration `yaml:"time_idle" env-default:"30s"`
}

type Postgres struct {
	Host     string `yaml:"host" env-default:"localhost"`
	Port     string `yaml:"port" env-default:"5432"`
	User     string `yaml:"user" env-required:"true"`
	Password string `yaml:"password" env-required:"true"`
	DBName   string `yaml:"dbname" env-required:"true"`
	SSLMode  string `yaml:"ssl_mode" env-default:"false"`
}

func MustLoad() *Config {
	var cfg Config

	// ПРОВЕРКА: есть ли переменная окружения с нашим конфигом
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		panic("CONFIG_PATH environment variable not set")
	}

	// ПРОВЕРКА: существует ли файл
	if _, err := os.Stat(configPath); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			panic("file does not exist")
		}
	}

	// запись файла в cfg
	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		panic(err)
	}

	return &cfg
}
