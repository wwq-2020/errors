package stack

import (
	"fmt"
	"path/filepath"
	"runtime"
)

// Get 获取堆栈
func Get(depth int) string {
	pc, _, line, ok := runtime.Caller(depth + 1)
	if !ok {
		return ""
	}
	return fmt.Sprintf("%v:%v", filepath.Base(runtime.FuncForPC(pc).Name()), line)
}
