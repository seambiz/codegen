package codegen

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/seambiz/codegen/config"
)

func ReadConfig(filename string) *config.Config {
	conf := &config.Config{}

	jsonBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(jsonBytes, &conf)
	if err != nil {
		panic(err)
	}

	return conf
}

func WriteConfig(filename string) {
	conf := config.Config{}

	jsonBytes, err := json.MarshalIndent(conf, "", "  ")
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(filename, jsonBytes, 0o644)
	if err != nil {
		panic(err)
	}
}
