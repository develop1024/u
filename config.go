package u

import (
	"github.com/spf13/viper"
	"sync"
)

var _config *viper.Viper
var _configOnce sync.Once

// Cfg 获取config对象
func Cfg() *viper.Viper {
	_configOnce.Do(func() {
		_config = viper.New()
	})
	return _config
}
