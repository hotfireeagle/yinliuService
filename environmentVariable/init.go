package environmentVariable

import "flag"

// 数据库dsn定义
var DSN *string

// 初始化系统的环境变量
func init() {
	DSN = flag.String("DSN", "user:password@/dbname", "database dsn")
	flag.Parse()
}