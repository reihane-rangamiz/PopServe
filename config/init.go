package config

import (
	"log"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
)

func LoadConfig() *koanf.Koanf {
	config := "./config/config.yml"

	var knf = koanf.New(".")
	if err := knf.Load(file.Provider(config), yaml.Parser()); err != nil {
		log.Fatalf("error loading app config: %v", err)
	}

	return knf
}
