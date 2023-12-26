package main

import (
	"fmt"
	"os"

	"github.com/WormholeMapper/esiclient"
)

const DEFAULT_CONFIG_PATH = "/whmserver.json"

func main() {
	configPath, ok := os.LookupEnv("CONFIG_PATH")
	if !ok {
		configPath = DEFAULT_CONFIG_PATH
	}
	config := ReadConfigFromFile(configPath)
	fmt.Println("hello from whmserver")
	fmt.Printf("read config from %v: %+v\n", configPath, config)
	esiclient.CreateClient(nil, "Wormhole Mapper dev / Squish Padecain")
}
