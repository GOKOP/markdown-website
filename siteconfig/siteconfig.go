package siteconfig

import (
	"io/ioutil"
	"log"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Port string
	Files []string
}

func Read(file string) Config {

	var config Config
	source, err := ioutil.ReadFile(file)

	if err != nil {
		log.Fatal("Reading config: "+err.Error())
	}

	err = yaml.Unmarshal(source, &config)

	if err != nil {
		log.Fatal("Decoding config: "+err.Error())
	}

	return config
}
