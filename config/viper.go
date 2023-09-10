package config

import (
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
)

var (
	_, b, _, _ = runtime.Caller(0)
	RootPath   = filepath.Join(filepath.Dir(b), "../")
)

type Instance struct {
	Package *viper.Viper
}

func loadEnv(v *viper.Viper) {
	envFilePath := filepath.Join(RootPath, ".env")
	//check .env is exist
	_, err := os.Stat(envFilePath)
	if !os.IsNotExist(err) {
		v.SetConfigFile(envFilePath)
		err = v.ReadInConfig() // Find and read the config file
		if err != nil {        // Handle errors reading the config file
			log.Fatalf("Fatal error config:%v", err)
		}
	}
	v.AllowEmptyEnv(true)
	v.AutomaticEnv()
}

func loadYaml(v *viper.Viper) {
	v.SetConfigName("development")
	v.SetConfigType("yaml")
	v.AddConfigPath(RootPath)
	err := v.ReadInConfig() // Find and read the config file
	if err != nil {         // Handle errors reading the config file
		log.Fatalf("Fatal error config:%v", err)
	}
}

func NewViper(t string) *viper.Viper {
	v := viper.New()
	switch t {
	case "env":
		loadEnv(v)
	case "yaml":
		loadYaml(v)
	}

	return v
}

func (i *Instance) GetString(name string) string {
	return i.Package.GetString(name)
}

func (i *Instance) GetInt(name string) int {
	return i.Package.GetInt(name)
}

func (i *Instance) GetBool(name string) bool {
	return i.Package.GetBool(name)
}

func (i *Instance) GetStringSlice(name string) []string {
	return i.Package.GetStringSlice(name)
}

func (i *Instance) Set(key string, value interface{}) {
	i.Package.Set(key, value)
}
