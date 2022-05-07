package auxiliary

import (
	"github.com/laurent22/toml-go"
	"runtime"
)

/*
 Get required field from config file.
*/
func GetFromConfig(configField string) string {
	os := runtime.GOOS
	config_path := "/crypto/config/config.toml"
	var parser toml.Parser
	if os == "windows" {
		config_path = "C:\\crypto\\config\\config.toml"
	}

	config := parser.ParseFile(config_path)
	return config.GetString(configField)
}
