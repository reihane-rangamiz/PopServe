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

	app := GenerateStructsFromMap(newConfig.App, "app")
	user := GenerateStructsFromMap(newConfig.User, "user")

	for _, model := range newConfig.Models {
		for modelName, modelfields := range model {
			newModel := GenerateStructsFromMap(map[string]interface{}{modelName: modelfields}, modelName)
			os.WriteFile("./models/models.go", []byte(newModel), os.ModeAppend|0644)
		}

		// newModel := GenerateStructsFromMap(model, "model")
		// os.WriteFile("./models/models.go", []byte(newModel), 0644)

	}
	os.WriteFile("./models/models.go", []byte(app), 0644)
	os.WriteFile("./models/models.go", []byte(user), 0644)

	return newConfig, nil
}

func GenerateStructsFromMap(input map[string]interface{}, structName string) string {
	var result strings.Builder

	result.WriteString(fmt.Sprintf("type %s struct {\n", utils.Capitalize(structName)))
	for fieldName, value := range input {

		var gormTag string
		fieldType, skip, tag := utils.DetectType(fieldName, value)
		if skip {
			continue
		}
		gormTag = tag
		jsonTag := fmt.Sprintf(`json:"%s"`, fieldName)

		tags := fmt.Sprintf("`%s", jsonTag)
		if gormTag != "" {
			tags += fmt.Sprintf(` gorm:"%s"`, gormTag)
		}
		tags += "`"

		result.WriteString(fmt.Sprintf("    %s %s %s\n", utils.Capitalize(fieldName), fieldType, tags))

		result.WriteString("}\n\n")
	}

	return result.String()
}
