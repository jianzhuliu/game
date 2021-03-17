package core

import (
	"encoding/json"
	"fmt"
	"os"
)

type Conf struct {
	Network        string //服务器协议类型，tcp tcp4 tcp6
	MaxConn        int    //最大连接数
	MaxMsgLen      int    //最大包长度
	WorkerPoolSize int    //工作池招聘工人数
	WorkerTaskSize int    //每个工人最大任务数
}

var ConfObj *Conf

func initConf() {
	ConfObj = &Conf{
		Network:        "tcp4",
		MaxConn:        1000000,
		MaxMsgLen:      4096,
		WorkerPoolSize: 10,
		WorkerTaskSize: 1024,
	}

	ConfObj.reload()
}

//基本配置信息打印
func (c *Conf) String() string {
	return fmt.Sprintf("Network=%s, MaxConn=%d, MaxMsgLen=%d, WorkerPoolSize=%d, WorkerTaskSize=%d",
		c.Network,
		c.MaxConn,
		c.MaxMsgLen,
		c.WorkerPoolSize,
		c.WorkerTaskSize,
	)
}

//加载配置文件
func (c *Conf) reload() {
	//fs, err := os.OpenFile(C_CONF_PATH, )
	data, err := os.ReadFile(C_CONF_PATH)
	if err != nil {
		if os.IsNotExist(err) {
			return
		}
	}

	err = json.Unmarshal(data, c)
	if err != nil {
		panic(fmt.Sprintf("conf file(%s) is not right, err=%v", C_CONF_PATH, err))
	}
}
