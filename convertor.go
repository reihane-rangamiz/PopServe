package main

import "PopServe/config"

type Config struct {
	User   map[string]interface{}   `koanf:"user"`
	Models []map[string]interface{} `koanf:"models"`
	App    map[string]interface{}   `koanf:"app"`
	// DataBase map[string]interface{}   `yml:"database"`

}

func AppConfig() (*Config, error) {
	var newConfig = &Config{}
	newKoanf := config.LoadConfig()
	if err := newKoanf.Unmarshal("config", newConfig); err != nil {
		return nil, err
	}

	return newConfig, nil
}
