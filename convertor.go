package main

import (
	"log"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
)

type Config struct {
	User   map[string]interface{}   `koanf:"user"`
	Models []map[string]interface{} `koanf:"models"`
	App    map[string]interface{}   `koanf:"app"`
	// DataBase map[string]interface{}   `yml:"database"`

}

func AppConfig() (*Config, error) {
	var newConfig = &Config{}
	newKoanf := LoadConfig()
	if err := newKoanf.Unmarshal("config", newConfig); err != nil {
		return nil, err
	}

	return newConfig, nil
}

func LoadConfig() *koanf.Koanf {
	config := "config.yml"

	var knf = koanf.New(".")
	if err := knf.Load(file.Provider(config), yaml.Parser()); err != nil {
		log.Fatalf("error loading app config: %v", err)
	}

	return knf
}
