package ginp

import (
	"fmt"
	"github.com/spf13/viper"
)

const (
	DefaultConfigDir  = "configs"
	DefaultConfigType = "yaml"
)

func NewConfig() *inConfig {

	v := viper.New()
	v.SetConfigType(DefaultConfigType)
	v.AddConfigPath(DefaultConfigDir)
	v.AddConfigPath(".")

	if err := v.ReadInConfig(); err != nil {
		fmt.Println(err.Error())
		panic(err.Error())
	}
	return &inConfig{
		v,
	}
}

type inConfig struct {
	*viper.Viper
}
