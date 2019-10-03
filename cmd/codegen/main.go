package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"bitbucket.org/codegen"
	_ "github.com/go-sql-driver/mysql"
)

var buildTime string
var gitRevision string
var gitBranch string

func main() {
	// Subcommands
	updateCommand := flag.NewFlagSet("update", flag.ExitOnError)
	genCommand := flag.NewFlagSet("gen", flag.ExitOnError)
	versionFlag := flag.Bool("v", false, "Print the current version and exit")

	configFile := flag.String("config", "codegen.json", "config file (required)")
	flag.Parse()

	if *versionFlag {
		fmt.Printf("build time:   %s\n", buildTime)
		fmt.Printf("git revision: %s\n", gitRevision)
		fmt.Printf("git branch:   %s\n", gitBranch)
		return
	}

	if *configFile == "" {
		fmt.Println("config file is required")
		flag.PrintDefaults()
		os.Exit(1)
	}

	if len(flag.Args()) == 0 {
		fmt.Println("subcommand < update | gen > is required")
		flag.PrintDefaults()
		os.Exit(1)
	}

	var conf codegen.Config
	jsonBytes, err := ioutil.ReadFile(*configFile)
	if err != nil {
		fmt.Println("invalid file")
		fmt.Println(err)
		os.Exit(1)
	}
	err = json.Unmarshal(jsonBytes, &conf)
	if err != nil {
		fmt.Println("config not valid json")
		fmt.Println(err)
		os.Exit(1)
	}

	switch flag.Args()[0] {
	case "update":
		err = updateCommand.Parse(os.Args[2:])
		if err != nil {
			panic(err)
		}

	case "gen":
		err = genCommand.Parse(os.Args[2:])
		if err != nil {
			panic(err)
		}

	default:
		flag.PrintDefaults()
		os.Exit(1)
	}

	if genCommand.Parsed() {
		codegen.Generate(&conf)
	}

	if updateCommand.Parsed() {
		contents, err := codegen.Update(&conf)
		if err != nil {
			panic(err)
		}
		err = ioutil.WriteFile(*configFile, contents, 0x644)
		if err != nil {
			panic(err)
		}
	}
}
