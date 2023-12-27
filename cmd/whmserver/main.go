package main

import (
	ctx "context"
	"fmt"
	"os"

	"github.com/WormholeMapper/esiclient"
	"github.com/WormholeMapper/whmcfg"
)

const DEFAULT_CONFIG_PATH = "/whmserver.json"

func main() {
	configPath, ok := os.LookupEnv("CONFIG_PATH")
	if !ok {
		configPath = DEFAULT_CONFIG_PATH
	}
	config := whmcfg.ReadConfigFromFile(configPath)
	fmt.Println("hello from whmserver")
	fmt.Printf("read config from %v: %+v\n", configPath, config)
	client := esiclient.CreateClient(nil, "Wormhole Mapper dev / Squish Padecain")
	status, resp, err := client.ESI.StatusApi.GetStatus(ctx.Background(), nil)
	fmt.Printf("GET /status: %+v, %+v, %+v", status, resp, err)
}
