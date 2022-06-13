package ginp

import (
	"github.com/spf13/viper"
	"os"
	"strings"
)

const (
	DefaultConfigDir  = "configs"
	DefaultConfigType = "yaml"
)

func NewConfig() *iConfig {
	vp := viper.New()
	vp.SetConfigType(DefaultConfigType)
	vp.AddConfigPath(DefaultConfigDir)
	vp.AddConfigPath(".")

	return &iConfig{
		vp,
	}
}

type iConfig struct {
	*viper.Viper
}

func FormatEnvKey(s string) string {
	return strings.ToUpper(strings.Replace(s, ".", "_", -1))
}

func (i *iConfig) ReadData() error {
	return i.ReadInConfig()
}

func (i *iConfig) GetWithEnv(key string) interface{} {
	v := os.Getenv(FormatEnvKey(key))
	if len(v) > 0 {
		return v
	} else {
		return i.Get(key)
	}
}
