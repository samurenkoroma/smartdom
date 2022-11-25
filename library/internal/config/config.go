package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"sync"
)

type Config struct {
	MongoDB struct {
		Host     string `env:"HOST"  env-default:"localhost"`
		Port     string `env:"PORT"  env-default:"27017"`
		Database string `env:"DATABASE"  env-default:"library"`
		Username string `env:"USERDB"   env-default:"library"`
		Password string `yaml:"PASSDB"  env-default:"library"`
	}
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		instance = &Config{}
		if err := cleanenv.ReadEnv(instance); err != nil {
			log.Fatal(err)
		}
	})
	return instance
}
