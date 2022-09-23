package codegen

import "bitbucket.org/codegen/config"

type UpdateCmd interface {
	Update(conf *config.Config) (config.Config, error)
}
