package config

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
)

// Config -
var (
	Config configFile
	Flags  configFlags
)

type configFile struct {
	JWTSecret string `json:"JWTSecret"`
	Address   string `json:"Address"`
}

type configFlags struct {
	Prod bool
	TLS  bool
}

// Init - read config file
func Init() {

	parseConfigFile()
	parseFlags()
}

func parseConfigFile() {

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

func parseFlags() {
	Flags.Prod = false
	Flags.TLS = false

	prod := flag.Bool("p", false, "Run in production")
	tls := flag.Bool("tls", false, "Use TLS")

	flag.Parse()

	Flags.Prod = *prod
	Flags.TLS = *tls

	if *prod {
		log.Println("Running in production mode")
	}

}
