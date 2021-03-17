package iface

import "io"

//日志接口
type Ilogger interface {
	//设置输出对象
	SetOutput(io.Writer)

	//基本打印方法
	Printf(format string, v ...interface{})
	Print(v ...interface{})

	//发生错误，需要退出
	Fatal(v ...interface{})
	Fatalf(format string, v ...interface{})

	//panic
	Panic(v ...interface{})
	Panicf(format string, v ...interface{})
}
