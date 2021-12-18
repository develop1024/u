package u

import (
	"sync"
	"github.com/spf13/viper"
)

var config *viper.Viper
var configOnce sync.Once

// Cfg 获取config对象
func Cfg() *viper.Viper {
	configOnce.Do(func() {
		config = viper.New()
	})
	return config
}
