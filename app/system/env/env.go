package env

import (
	"sync"

	"github.com/jinzhu/configor"
)

type Config struct {
	SAMPLE_VARIABLE string `required:"true" env:"_"`
}

var (
	configOnce sync.Once
	configInst *Config
)

func GetConfig() *Config {
	configOnce.Do(func() {
		if err := configor.Load(&configInst); err != nil {
			panic(err)
		}
	})
	return configInst
}
