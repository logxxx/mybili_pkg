package runutil

import (
	"pkg.logxxx.com/utils/log"
	"runtime"
)

// 安全地执行goroutine。包装了recover()捕获异常
func RunSafe(fn func()) {
	if fn == nil {
		return
	}
	go func() {
		defer func() {
			if err := recover(); err != nil {
				buf := make([]byte, 1024)
				_ = runtime.Stack(buf, false)
				log.Errorf("RunSafe panic:%v stack:%s", err, buf)
			}
		}()

		fn()
	}()
}
