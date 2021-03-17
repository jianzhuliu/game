package core

import (
	"fmt"

	"gitee.com/jianzhuliu/game/iface"
)

//数据对象，数据结构采用 [长度+类型+数据] 结构
type Message struct {
	MsgLen uint32 //数据长度
	Data   []byte //数据
	MsgId  uint32 //id
}

//构造数据对象
func NewMessage(msgId uint32, data []byte) iface.Imsg {
	return &Message{
		MsgLen: uint32(len(data)),
		MsgId:  msgId,
		Data:   data,
	}
}

//获取数据id
func (m *Message) GetMsgId() uint32 {
	return m.MsgId
}

//获取数据长度
func (m *Message) GetMsgLen() uint32 {
	return m.MsgLen
}

//获取数据
func (m *Message) GetData() []byte {
	return m.Data
}

//设置数据id
func (m *Message) SetMsgId(id uint32) {
	m.MsgId = id
}

//设置数据长度
func (m *Message) SetMsgLen(msgLen uint32) {
	m.MsgLen = msgLen
}

//设置数据
func (m *Message) SetData(data []byte) {
	m.Data = data
}

//打印数据
func (m *Message) String() string {
	return fmt.Sprintf("msgId=%d, msgLen=%d, msgData=%v",
		m.GetMsgId(),
		m.GetMsgLen(),
		string(m.GetData()),
	)
}
