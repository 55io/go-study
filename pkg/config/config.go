package config

import (
	"log"
	"os"

	"sync"

	"gopkg.in/yaml.v3"
)

type App struct {
	Port string `yaml:"port"`
}

type Redis struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	DB       int    `yaml:"db"`
}

type Db struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Dbname   string `yaml:"dbname"`
	Schema   string `yaml:"schema"`
	SllMode  string `yaml:"ssl_mode"`
}

type config struct {
	App   `yaml:"app"`
	Db    `yaml:"db"`
	Redis `yaml:"redis"`
}

func getInstance() (cf config) {
	var instance *config = nil
	sync.Once.Do(func() {
		cPath := os.Args[1]
		cf.LoadConfig(cPath)
		instance = new(Config)
	})
	return instance
}

func (cg *config) LoadConfig(patch string) {
	file, err := os.Open(patch)
	if err != nil {
		log.Fatal("ReadConfigFile: ", err)
	}

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&cg)
	if err != nil {
		log.Fatal("DecodeConfigFile: ", err)
	}
}
