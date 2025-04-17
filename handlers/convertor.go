package handlers

import (
	"PopServe/config"
	"PopServe/utils"
	"fmt"
	"os"
	"strings"
)

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

	app := GenerateStructsFromMap(newConfig.App)
	user := GenerateStructsFromMap(newConfig.User)
	
	for _, model := range newConfig.Models {
		newModel := GenerateStructsFromMap(model)
		os.WriteFile("./models/models.go", []byte(newModel), 0644)

	}
	os.WriteFile("./models/models.go", []byte(app), 0644)
	os.WriteFile("./models/models.go", []byte(user), 0644)


	return newConfig, nil
}

func GenerateStructsFromMap(input map[string]interface{}) string {
	var result strings.Builder

	for key, v := range input {
		fieldsMap, ok := v.(map[string]interface{})
		if !ok {
			continue
		}

		result.WriteString(fmt.Sprintf("type %s struct {\n", utils.Capitalize(key)))
		for fieldName, fieldValue := range fieldsMap {
			goType := utils.InferGoType(fieldValue)
			result.WriteString(fmt.Sprintf("    %s %s `json:\"%s\"`\n", utils.Capitalize(fieldName), goType, fieldName))
		}
		result.WriteString("}\n\n")
	}

	return result.String()
}
