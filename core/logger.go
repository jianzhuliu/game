package core

import (
	"log"
	"os"

	"gitee.com/jianzhuliu/game/iface"
)

//日志处理
var Logger iface.Ilogger

func initLogger() {
	Logger = log.New(os.Stderr, "", log.LstdFlags|log.Lshortfile)
}
