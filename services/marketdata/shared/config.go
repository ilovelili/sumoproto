package marketdata

import (
	"encoding/json"
	"os"
	"path"
)

// GetConfig get config defined in config.json
func GetConfig() (config *Config, err error) {
	path := path.Join(Getwd(), "config", "config.json")
	configFile, err := os.Open(path)
	defer configFile.Close()

	if err != nil {
		return
	}

	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)
	return
}

// Account account entry
type Account struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

// Config config entry
type Config struct {
	Account `json:"account"`
}
