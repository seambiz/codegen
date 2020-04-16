package codegen

type UpdateCmd interface {
	Update(conf *Config) (Config, error)
}
