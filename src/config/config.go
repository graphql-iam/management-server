package config

import (
	"errors"
	"gopkg.in/yaml.v3"
	"io"
	"os"
)

type Config struct {
	Port     int    `yaml:"port"`
	MongoUrl string `yaml:"mongoUrl"`
}

const ConfigPathEnvName = "MANAGEMENT_SERVER_CONFIG_PATH"

func NewConfig() Config {
	configPath := "./config.yaml"

	if os.Getenv(ConfigPathEnvName) != "" {
		configPath = os.Getenv(ConfigPathEnvName)
	}

	cfg, err := getConfig(configPath)
	if err != nil {
		panic(err)
	}

	return cfg
}

func getConfig(path string) (Config, error) {
	var res Config

	file, err := os.Open(path)
	if err != nil {
		return res, err
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return res, err
	}

	err = yaml.Unmarshal(bytes, &res)
	if err != nil {
		return res, err
	}

	err = res.validateAndFillDefaults()
	return res, err
}

func (c *Config) validateAndFillDefaults() error {
	if c.Port <= 0 {
		c.Port = 8080
	}
	if c.MongoUrl == "" {
		return errors.New("no mongoUrl provided in config")
	}
	return nil
}
