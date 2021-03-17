package core

import (
	"bytes"
	"reflect"
	"testing"
)

func TestDbpack(t *testing.T) {
	data := []byte("Hello World!!!")
	msg := NewMessage(1, data)

	//打包
	dbpack := NewDbpack()
	binaryData, err := dbpack.Pack(msg)
	if err != nil {
		t.Fatal(err)
	}

	//t.Log(binaryData)
	//解包
	msgUnpack, err := dbpack.Unpack(bytes.NewReader(binaryData))
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("msgLen=%d, msgId=%d, msgData=%v", msgUnpack.GetMsgLen(), msgUnpack.GetMsgId(), string(msgUnpack.GetData()))
	//深度比较
	if !reflect.DeepEqual(data, msgUnpack.GetData()) {
		t.Fatalf("expect %v, but got %v", string(data), string(msgUnpack.GetData()))
	}

}
