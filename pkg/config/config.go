package config

import (
	"log"
	"os"

	"sync"

	"gopkg.in/yaml.v3"
)

type App struct {
	Port string `yaml:"port"`
	Env  string `yaml:"env"`
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

var once sync.Once

func Get() (cf *config) {
	var instance *config = nil
	once.Do(func() {
		cPath := os.Getenv("CONFIG_PATH")
		cf.loadConfig(cPath)
		instance = new(config)
	})
	return instance
}

func (cg *config) loadConfig(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("ReadConfigFile: ", err)
	}

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&cg)
	if err != nil {
		log.Fatal("DecodeConfigFile: ", err)
	}
}
