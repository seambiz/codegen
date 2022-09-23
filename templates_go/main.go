package main

import "bitbucket.org/codegen"

func main() {
	conf := codegen.ReadConfig("../codegen.json.gen")

	codegen.GenGo(conf, Constants{}, "test/constants.go")
	codegen.GenGo(conf, DBTypes{}, "test/dbtypes.go")
	codegen.GenGo(conf, Router{}, "test/router.go")

	codegen.Lint("test/", nil)
}
