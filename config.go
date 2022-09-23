package codegen

import (
	"encoding/json"
	"io/ioutil"

	"bitbucket.org/codegen/config"
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
