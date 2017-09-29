package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"bitbucket.com/codegen"
	"github.com/BurntSushi/toml"
	_ "github.com/go-sql-driver/mysql"
)

/*

TODO: update with bitset to allow custom updates with partial fields


*/

func main() {
	// Subcommands
	updateCommand := flag.NewFlagSet("update", flag.ExitOnError)
	genCommand := flag.NewFlagSet("gen", flag.ExitOnError)

	configFile := flag.String("config", "", "config file (required)")
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
	_, err := toml.DecodeFile(*configFile, &conf)
	if err != nil {
		fmt.Println("config not valid toml")
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
