package main

import (
	"encoding/json"
	"os"
)

type AppConfig struct {
	DB struct {
		Host string `json:"host"`
		Port string `json:"port"`
		User string `json:"user"`
		Pass string `json:"pass"`
	} `json:"db"`
	Esi struct {
		ClientID    string `json:"client_id"`
		SecretKey   string `json:"secret_key"`
		CallbackURL string `json:"callback_url"`
	} `json:"esi"`
}

// ReadConfigFromFile attempts to read the application configuration from the provided JSON file.
func ReadConfigFromFile(filename string) AppConfig {
	// Try to open the specified file.
	file, err := os.Open(filename)
	if err != nil {
		panic("Failed to open config file: " + err.Error())
	}

	// Pre-allocate a []byte buffer of size equal to that of the file.
	stat, _ := file.Stat()
	buf := make([]byte, stat.Size())

	// Read the data from the file into the buffer.
	_, err = file.Read(buf)
	if err != nil {
		panic("Failed to read from file: " + err.Error())
	}

	// Try to convert the buffer to JSON compatible with the AppConfig.
	var config AppConfig
	err = json.Unmarshal(buf, &config)
	if err != nil {
		panic("Failed to unmarshal config: " + err.Error())
	}

	return config
}
