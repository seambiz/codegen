package codegen

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/seambiz/codegen/config"
	"github.com/seambiz/codegen/static"
)

func DumpTemplates(conf *config.Config) {
	uniqueFiles := make(map[string]bool)
	GenerateTemplatesConfig(conf)

	for _, schema := range conf.Schemas {
		for _, files := range schema.PreparedTemplatefiles {
			for _, filename := range files {
				uniqueFiles[filename] = true
			}
		}

		for _, table := range schema.Tables {
			for _, files := range table.PreparedTemplatefiles {
				for _, filename := range files {
					uniqueFiles[filename] = true
				}
			}
		}
	}

	fmt.Println(uniqueFiles)
	for filename := range uniqueFiles {
		contents, err := static.ReadFile(filename)
		if err != nil {
			panic(err)
		}

		os.WriteFile(filepath.Join(conf.TemplateFolder, filename), contents, 0o644)
	}
}
