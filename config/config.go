package config

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	MinimumLogLevel int

	ListenAddress string
}

var Current *Config

func Load() {
	file, err := os.Open("config.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	config := make(map[string]interface{})
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatal(err)
	}

	config = config[config["conf_type"].(string)].(map[string]interface{})

	Current = &Config{
		MinimumLogLevel: int(config["min_log_level"].(float64)),
		ListenAddress:   config["listen_address"].(string),
	}
}
