package config

import (
	"encoding/json"
	"os"
)

type Configuration struct {
	Chan   string
	Key    string
	Server string
}

var c *Configuration

func Init() *Configuration {
	if os.Getenv("ENV") == "production" {
		//TODO: build it off a slice or something
		c = &Configuration{
			Chan: os.Getenv("CHAN"),
			Key:  os.Getenv("KEY"),
		}
	} else {
		file, err := os.Open("config/config.json")
		if err != nil {
			panic(err)
		}
		defer file.Close()
		decoder := json.NewDecoder(file)
		c = new(Configuration)
		err = decoder.Decode(&c)
		if err != nil {
			panic(err)
		}
	}
	return c
}

func Get() *Configuration {
	return c
}
