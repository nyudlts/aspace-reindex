package main

import (
	"flag"
	"fmt"
	"github.com/nyudlts/go-aspace"
	"io"
	"os"
)

var (
	config      string
	environment string
	help        bool
	version     bool
)

const vers = "1.0.0"

func init() {
	flag.StringVar(&config, "config", "", "location of go-aspace config file")
	flag.StringVar(&environment, "environment", "", "aspace environment")
	flag.BoolVar(&help, "help", false, "print the help message")
	flag.BoolVar(&version, "version", false, "print the version info")
}

func printHelp() {
	fmt.Printf("Usage: %s [OPTIONS]\n", os.Args[0])
	fmt.Println("Options")
	fmt.Println(" --config\t/path/to/go-aspace.yml\tmandatory")
	fmt.Println(" --environment\tenvironment to reindex\tmandatory")
	fmt.Println(" --help\tprint the help message")
	fmt.Println(" --version\tprint the version info\n")
}

func printVersion() {
	fmt.Printf("%s %s\n", os.Args[0], vers)
}

func main() {
	flag.Parse()
	if help {
		printHelp()
		os.Exit(0)
	}

	if version {
		printVersion()
		os.Exit(0)
	}

	printVersion()
	if config == "" || environment == "" {
		fmt.Println("ERROR: --config and --environment flags must both be set\n")
		printHelp()
		os.Exit(1)
	}

	client, err := aspace.NewClient(config, environment, 20)
	if err != nil {
		panic(err)
	}

	response, err := client.JsonRequest("/plugins/reindex", "POST", "")
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
	os.Exit(0)
}
