package config

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
)

// Config -
var (
	Config config
	Flags  flags
)

type config struct {
	JWTSecret string   `json:"JWTSecret"`
	Address   string   `json:"Address"`
	Database  database `json:"Database"`
}

type database struct {
	Address  string `json:"Address"`
	User     string `json:"User"`
	Password string `json:"Password"`
	Name     string `json:"Name"`
}

type flags struct {
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
