package main

import (
	"fmt"
	"net"
	"runtime"
	"time"

	"gitee.com/jianzhuliu/game/core"
	"gitee.com/jianzhuliu/game/examples/services"

	"gitee.com/jianzhuliu/game/examples/pb"
	"github.com/golang/protobuf/proto"
)

func main() {
	//设置最大goroutine 调度器个数
	numGO := runtime.NumCPU() * 3 / 4
	if numGO < 1 {
		numGO = 1
	}
	runtime.GOMAXPROCS(numGO)

	conn, err := net.Dial("tcp", "127.0.0.1:8999")
	if err != nil {
		core.Logger.Fatal("fail to connect", err)
	}

	go func(conn net.Conn) {

		//读数据
		dbpack := core.NewDbpack()
		for {
			proto_data, err := dbpack.Unpack(conn)
			if err != nil {
				core.Logger.Print("msg unpack fail, err=", err)
				return
			}

			//解析数据包
			msg := &pb.Talk{}
			err = proto.Unmarshal(proto_data.GetData(), msg)
			if err != nil {
				core.Logger.Print("proto.Unmarshal err,", err)
				return
			}

			core.Logger.Print("received ======> ", msg.GetContent())
		}

	}(conn)

	dbpack := core.NewDbpack()

	for i := 0; i < 10; i++ {
		content := fmt.Sprintf("hello %d", i)
		//写数据
		//构造数据包
		proto_data := &pb.Talk{
			Content: content,
		}

		//数据打包
		data, err := proto.Marshal(proto_data)
		if err != nil {
			core.Logger.Print("proto.Marshal err,", err)
			return
		}

		msg := core.NewMessage(services.C_MSG_ID_CHAT, data)
		binaryMsg, err := dbpack.Pack(msg)
		if err != nil {
			core.Logger.Print("msg pack fail,", i, ", err=", err)
			break
		}
		conn.Write(binaryMsg)

		time.Sleep(1 * time.Second)
	}
}
