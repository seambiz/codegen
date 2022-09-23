package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/seambiz/codegen"
	"github.com/seambiz/codegen/db"
	"github.com/seambiz/codegen/updater"
	"github.com/seambiz/seambiz/sdb"
)

var (
	version      string
	gitCommit    string
	gitTreeState string
	gitSha       string
	gitTag       string
	buildTime    string
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	// Subcommands
	updateCommand := flag.NewFlagSet("update", flag.ExitOnError)
	genCommand := flag.NewFlagSet("gen", flag.ExitOnError)
	versionFlag := flag.Bool("v", false, "Print the current version and exit")

	configFile := flag.String("config", "codegen.json", "config file (required)")
	flag.Parse()

	if *versionFlag {
		fmt.Printf("version:      %s\n", version)
		fmt.Printf("build time:   %s\n", buildTime)
		fmt.Printf("git commit:   %s\n", gitCommit)
		fmt.Printf("git sha:      %s\n", gitSha)
		fmt.Printf("git treestate:%s\n", gitTreeState)
		fmt.Printf("git tag:      %s\n", gitTag)
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

	switch flag.Args()[0] {
	case "update":
		err := updateCommand.Parse(os.Args[2:])
		if err != nil {
			panic(err)
		}

	case "gen":
		err := genCommand.Parse(os.Args[2:])
		if err != nil {
			panic(err)
		}

	default:
		flag.PrintDefaults()
		os.Exit(1)
	}

	if genCommand.Parsed() {
		conf := codegen.ReadConfig(*configFile + ".gen")
		codegen.Generate(conf)
	}

	if updateCommand.Parsed() {
		var up codegen.UpdateCmd
		var err error

		conf := codegen.ReadConfig(*configFile)

		ctx := &codegen.Context{
			Log: &log.Logger,
		}

		conn := sdb.OpenDatabaseDSN(conf.Database.DSN)
		repoTable := db.NewTablesRepo(conn)
		repoStats := db.NewStatisticsRepo(conn)
		repoKeyCol := db.NewKeyColumnUsageRepo(conn)
		repoCols := db.NewColumnsRepo(conn)

		up = updater.NewMysqlUpdate(ctx, repoTable, repoCols, repoKeyCol, repoStats)
		conf2, err := up.Update(conf)
		if err != nil {
			panic(err)
		}

		jsonBytes, err := json.MarshalIndent(conf2, "", "\t")
		if err != nil {
			panic(err)
		}

		err = os.WriteFile(*configFile+".gen", jsonBytes, 0o644)
		if err != nil {
			panic(err)
		}
	}
}
