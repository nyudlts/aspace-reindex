package main

import (
	"flag"
	"fmt"
	"github.com/nyudlts/go-aspace"
	"io/ioutil"
	"os"
)

var (
	config      string
	environment string
)

func init() {
	flag.StringVar(&config, "config", "", "location of go-aspace config file")
	flag.StringVar(&environment, "environment", "", "aspace environment")
}

func main() {
	flag.Parse()
	client, err := aspace.NewClient(config, environment, 20)
	if err != nil {
		panic(err)
	}

	response, err := client.JsonRequest("/plugins/reindex", "POST", "")
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
	os.Exit(0)
}
