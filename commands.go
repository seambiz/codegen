package codegen

import "github.com/seambiz/codegen/config"

type UpdateCmd interface {
	Update(conf *config.Config) (config.Config, error)
}
