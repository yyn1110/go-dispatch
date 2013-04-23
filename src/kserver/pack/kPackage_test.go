package pack

import (
	"fmt"
	"testing"
)

func checkLog(arg string) {
	fmt.Println("--------" + arg + "-------")

}
func TestHead(t *testing.T) {
	fmt.Println("test TestHead")
	msgID := 1
	bodyLen := 10

	head := NewHead(uint16(msgID), uint16(bodyLen))

	//head := &Head{[2]byte(msgID), [2]byte(bodyLen), [2]byte(fixcode)}
	by, err := head.encodeHead()
	if err != nil {
		t.Log(by)
		fmt.Println(by)

	}

}
func TestDecodeHead(t *testing.T) {
	checkLog("test TestDecodeHead")
	msgID := 1
	bodyLen := 10

	head := NewHead(uint16(msgID), uint16(bodyLen))

	//head := &Head{[2]byte(msgID), [2]byte(bodyLen), [2]byte(fixcode)}
	by, _ := head.encodeHead()
	he := DecodeHead(by)
	fmt.Println(he.msgID, he.bodyLen, he.fixedCode)
	t.Log(he.msgID, he.bodyLen, he.fixedCode)
}
func TestDecodeBody(t *testing.T) {
	checkLog("test TestDecodeBody")
	body := NewBody('A', "hello")
	by, err := body.encodeBody()
	if err != nil {
		t.Log(by)
		fmt.Println(by)
		fmt.Println(err)
	}

	b := DecodeBody(by)
	fmt.Println(byte(b.jsonKey), b.jsonLen, b.jsonStr)
	t.Log(byte(b.jsonKey), b.jsonLen, b.jsonStr)
}
func TestBody(t *testing.T) {
	checkLog("test TestBody")
	body := NewBody('A', "hello")
	by, err := body.encodeBody()
	if err != nil {
		t.Log(by)
	}
}

func TestPackageData(t *testing.T) {
	head := NewHead(uint16(1), uint16(20))
	body := NewBody('A', "hello")
	data := NewPackage(head, body)
	by, err := data.encodePackageData()
	if err != nil {
		t.Log(by)
	}
}
