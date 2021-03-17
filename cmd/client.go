package main

import (
	"net"
	"runtime"
	"strconv"
	"time"

	"gitee.com/jianzhuliu/game/core"
)

//模拟客户端

func main() {
	//不全部占用cpu个数
	n := runtime.NumCPU() * 3 / 4
	if n < 1 {
		n = 1
	}

	runtime.GOMAXPROCS(n)

	conn, err := net.Dial("tcp", "127.0.0.1:8999")
	if err != nil {
		core.Logger.Fatal("fail to connect", err)
	}

	go func(conn net.Conn) {

		//读数据
		dbpack := core.NewDbpack()
		for {
			msg, err := dbpack.Unpack(conn)
			if err != nil {
				core.Logger.Print("msg unpack fail, err=", err)
				return
			}
			core.Logger.Print("received ======> ", msg)
		}

	}(conn)

	for i := 0; i < 10; i++ {
		//写数据
		dbpack := core.NewDbpack()
		msg := core.NewMessage(uint32(i%2+1), []byte("Hello "+strconv.Itoa(i)))
		binaryMsg, err := dbpack.Pack(msg)
		if err != nil {
			core.Logger.Print("msg pack fail,", i, ", err=", err)
			break
		}
		conn.Write(binaryMsg)

		time.Sleep(1 * time.Second)
	}
}
