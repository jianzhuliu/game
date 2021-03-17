package core

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"

	"gitee.com/jianzhuliu/game/iface"
)

//数据解包装包对象
type Dbpack struct{}

//构造对象
func NewDbpack() *Dbpack {
	return &Dbpack{}
}

//获取头部数据长度
func (db *Dbpack) GetHeaderLen() int {
	//数据格式， 长度uint32(4) + ID uint32(4) + 数据
	return 8
}

//打包
func (db *Dbpack) Pack(msg iface.Imsg) ([]byte, error) {
	buf := bytes.NewBuffer([]byte{})

	//写入长度
	if err := binary.Write(buf, binary.LittleEndian, msg.GetMsgLen()); err != nil {
		return nil, err
	}

	//写入ID
	if err := binary.Write(buf, binary.LittleEndian, msg.GetMsgId()); err != nil {
		return nil, err
	}

	//写入数据
	if err := binary.Write(buf, binary.LittleEndian, msg.GetData()); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil

}

//拆包
func (db *Dbpack) Unpack(r io.Reader) (iface.Imsg, error) {
	headerData := make([]byte, db.GetHeaderLen())

	//读取头部
	if _, err := io.ReadFull(r, headerData); err != nil {
		return nil, err
	}

	//解析头部长度与id
	msg, err := db.unpack(headerData)
	if err != nil {
		return nil, err
	}

	if msg.GetMsgLen() > 0 {
		//读取数据
		data := make([]byte, msg.GetMsgLen())
		if _, err := io.ReadFull(r, data); err != nil {
			return nil, err
		}

		msg.SetData(data)
	}

	return msg, nil
}

//拆包
func (db *Dbpack) unpack(data []byte) (iface.Imsg, error) {
	buf := bytes.NewReader(data)

	msg := &Message{}

	//读取长度
	if err := binary.Read(buf, binary.LittleEndian, &msg.MsgLen); err != nil {
		return nil, err
	}

	//读取id
	if err := binary.Read(buf, binary.LittleEndian, &msg.MsgId); err != nil {
		return nil, err
	}

	//包长度验证
	if ConfObj.MaxMsgLen > 0 && int(msg.GetMsgLen()) > ConfObj.MaxMsgLen {
		Logger.Printf("包长度超出上限，msgId=%d, msgLen=%d, maxMsgLen=%d", msg.GetMsgId(), msg.GetMsgLen(), ConfObj.MaxMsgLen)
		return nil, errors.New("超出包长度上限")
	}

	return msg, nil
}
