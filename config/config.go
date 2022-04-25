package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

var Config Configuration

type Configuration struct {
	DbAddr     string `json:"dbAddr"`
	DbPassWord string `json:"dbPassWord"`
	DbName     string `json:"dbName"`
}

// Load reads and loads configuration to Config variable
func Load() {
	var err error

	confLen := len("config.json")
	if confLen != 0 {
		err = readConfigFromJSON("config.json")
	}

	if err != nil {
		fmt.Printf(`Configuration not found. Please specify configuration. err: %s`, err.Error())
	}
}

// nolint: gosec
func readConfigFromJSON(configFilePath string) error {
	log.Printf("Looking for JSON config file (%s)", configFilePath)

	cfgFile, err := os.Open(configFilePath)
	if err != nil {
		log.Printf("Reading configuration from JSON (%s) failed: %v\n", configFilePath, err)
		return err
	}
	defer func() {
		err := cfgFile.Close()
		if err != nil {
			log.Printf("Cannot close the configuration file [%s]: %v\n", cfgFile.Name(), err)
		}
	}()

	err = json.NewDecoder(cfgFile).Decode(&Config)
	if err != nil {
		log.Printf("Reading configuration from JSON (%s) failed: %s\n", configFilePath, err)
		return err
	}

	log.Printf("Configuration has been read from JSON (%s) successfully\n", configFilePath)
	return nil
}
