// 日志封装
// 目前只是简单的打印，后续可以接入其它操作，比如报警等
package tool

import (
	"log"
	"os"
)

// 标准日志
var InfoLog = log.New(os.Stdout, "【Info】", log.LstdFlags).Println

// 错误日志
var ErrLog = log.New(os.Stderr, "【Error】", log.LstdFlags|log.Llongfile).Println