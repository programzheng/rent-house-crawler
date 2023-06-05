package config

var Cfg = New()

type Config interface {
	GetString(name string) string
	GetInt(name string) int
	GetBool(name string) bool
	GetStringSlice(name string) []string
	Set(key string, value interface{})
}

func New() Config {
	config := &Instance{
		Package: NewViper("yaml"),
	}

	return config
}

func GetProductionStatus() bool {
	return Cfg.GetString("env") == "production"
}
