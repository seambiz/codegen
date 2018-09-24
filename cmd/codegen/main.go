package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"bitbucket.org/seambiz/codegen"
	_ "github.com/go-sql-driver/mysql"
)

/*

TODO: update with bitset to allow custom updates with partial fields


*/

func main() {
	// Subcommands
	updateCommand := flag.NewFlagSet("update", flag.ExitOnError)
	genCommand := flag.NewFlagSet("gen", flag.ExitOnError)

	configFile := flag.String("config", "codegen.json", "config file (required)")
	flag.Parse()

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

	if conf.FieldsPerTable == 0 {
		conf.FieldsPerTable = 50
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
