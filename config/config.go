package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// Config -
var Config configStruct

type configStruct struct {
	JWTSecret string `json:"JWTSecret"`
	Address   string `json:"Address"`
}

// Init - read config file
func Init() {

	log.Println("Reading config file...")

	file, err := ioutil.ReadFile("./config.json")

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%s\n", string(file))

	err = json.Unmarshal(file, &Config)

	if err != nil {
		log.Println(err)
	}

}
