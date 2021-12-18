package u

import (
	"github.com/robfig/cron/v3"
	"sync"
)

var _cron *cron.Cron
var _cronOnce sync.Once

// Cron 获取cron对象
func Cron() *cron.Cron {
	_cronOnce.Do(func() {
		_cron = cron.New()
	})
	return _cron
}
