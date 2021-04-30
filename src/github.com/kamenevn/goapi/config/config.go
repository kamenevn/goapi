package config

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	Database struct {
		Host     string `json:"host" required:"true" default:"localhost"`
		Name     string `json:"name" required:"true"`
		User     string `json:"user" required:"true"`
		Password string `json:"password" required:"true"`
	} `json:"database"`

	Log []struct {
		File  string `json:"file" required:"true"`
	} `json:"log"`
}

var AppConfig *Config

func Get() *Config {
	var config *Config

	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	// TODO: как сюда конфиг то прописать?(
	var filePath = dir + "/src/github.com/kamenevn/goapi/config/config.json"

	configFile, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Error config loading: ", err.Error())
	}

	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)

	return config
}