package main

import (
	"encoding/hex"
	"fmt"

	"github.com/istabraq/iso8583/Iso8583"
)

func main() {
	data := "49534F373031303030303031383134823000008A00000000000000000000013139303231333232353239313731373931393032313332323532353030363330303332303930343432323931373137393830303030303030303030"
	//resp := "16010200A85332250000000110004D94101201ABF5000110722222810ED0820610476134000000004300000000000000320010091247222683361010020400110006404122F9F2F8F2F1F2F2F6F8F3F3F6F0F2F2F7F6F1F0F0F7F0F0F0F0F9F4F0F0F0F0F0F0F0F0F1F0F0F0F0F7F4F0094040404040404040F20952140100119F36020005910A98D04EC5E7B3BCFF30301040000000000000000309282460430470058000000002"
	src := []byte(data)
	//srcp := []byte(resp)
	dst := make([]byte, hex.DecodedLen(len(src)))
	//dstp := make([]byte, hex.DecodedLen(len(srcp)))
	_, err := hex.Decode(dst, src)
	//_, err = hex.Decode(dstp, srcp)
	if err != nil {
		panic(err)
	}
	msg := Iso8583.NewIso8583("gim")
	msgs := Iso8583.NewIso8583("gim")

	_,err = msgs.UnpackFull(dst, 0)
	if err != nil {
		panic(err)
	}
	fmt.Println(dst)
	fmt.Println(msgs.ToStringFull(""))
	
	err = msg.SetHeaderValue(-2,"ISO")
	if err != nil {
		panic(err)
	}
	err = msg.SetHeaderValue(-1,"70100000")
	if err != nil {
		panic(err)
	}
	err = msg.SetMti("1200")
	if err != nil {
		panic(err)
	}
	err = msg.SetFieldValue(2,"4555555555555555")
	if err != nil {
		panic(err)
	}
	err = msg.SetFieldValue(3,"000000")
	if err != nil {
		panic(err)
	}
	err = msg.SetFieldValue(4,"500")
	if err != nil {
		panic(err)
	}
	fmt.Println(msg.ToStringFull(""))
	fmt.Println(string(msg.ToMsgFull()))
	fmt.Println(msg.ToFields())
}
