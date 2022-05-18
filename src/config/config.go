package config

import (
	"io/ioutil"
	"log"

	"github.com/json-iterator/go"
)

type Mongodb struct {
	DBname string ` json:"name" `
	Host   string ` json:"host" `
	Port   string ` json:"port" `
	User   string ` json:"user" `
	Pwd    string ` json:"pwd"  `
}
type Config struct {
	Mongo    Mongodb `json:"mongo"`
	JWTToken string  `json:"jwtToken"`
	SavePath string  `json:"save path"`
}

var (
	GlobalConfig *Config
)

func init() {
	configFile := "config.json"
	data, err := ioutil.ReadFile(configFile)

	if err != nil {
		log.Println("Read config error!")
		log.Panic(err)
		return
	}

	config := &Config{}

	err = jsoniter.Unmarshal(data, config)

	if err != nil {
		log.Println("Unmarshal config error!")
		log.Panic(err)
		return
	}

	GlobalConfig = config
	log.Println("Config " + configFile + " loaded.")

}
