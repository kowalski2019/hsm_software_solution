package auxiliary

import (
	"github.com/laurent22/toml-go"
)

/*
 Get required field from config file.
*/
func GetFromConfig(configField string) string {
	var parser toml.Parser
	config := parser.ParseFile("/crypto/config/config.toml")
	return config.GetString(configField)
}
