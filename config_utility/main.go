package main

import (
	"flag"
	"log"

	config_utils "hh-go-utilities/config_utility/pkg"
)

var (
	configPath = flag.String("config_path", "data/config.env", "Config file to overide defaults for config env vars.")
)

func main() {
	flag.Parse()

	config, err := config_utils.LoadTestConfig(*configPath)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Test Config defined in pkg/config_utils_test_utils.go: \n", *config)
}
