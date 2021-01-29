package configs

import (
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	Env string
	Auth AuthConfig
	DB DBConfig
	Server ServerConfig
	Job JobConfig
}

func GetConfig(path string) (Config, error) {
	var cfg Config
	f, err := os.Open(path)
	if err != nil {
		return cfg, err
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		return cfg, err
	}
	return cfg, nil
}