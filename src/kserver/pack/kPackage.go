package pack

import (
	"bytes"
	"encoding/binary"
	"errors"
)

const (
	code           = 12321
	PACKAGEHEADLEN = 6
)

type Head struct {
	msgID     uint16
	bodyLen   uint16
	fixedCode uint16
}
type Body struct {
	jsonKey byte
	jsonLen uint16
	jsonStr string
}
type PackageData struct {
	head *Head
	body *Body
}

func (head *Head) encodeHead() ([]byte, error) {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, head.msgID)
	if err != nil {
		return []byte(""), errors.New("write msgid error")
	}
	err = binary.Write(buf, binary.LittleEndian, head.bodyLen)
	if err != nil {
		return []byte(""), errors.New("write bodyLen error")
	}
	err = binary.Write(buf, binary.LittleEndian, head.fixedCode)
	if err != nil {
		return []byte(""), errors.New("write fixedCode error")
	}
	return buf.Bytes(), nil
}
func DecodeHead(by []byte) (head *Head) {
	head = new(Head)
	head.msgID = binary.LittleEndian.Uint16(by[0:2])
	head.bodyLen = binary.LittleEndian.Uint16(by[2:4])
	head.fixedCode = binary.LittleEndian.Uint16(by[4:6])
	return
}
func (body *Body) encodeBody() ([]byte, error) {
	buf := new(bytes.Buffer)
	//err := binary.Write(buf, binary.LittleEndian, )
	err := buf.WriteByte(body.jsonKey)

	if err != nil {
		return []byte(""), errors.New("write jsonKey error")
	}
	err = binary.Write(buf, binary.LittleEndian, body.jsonLen)
	if err != nil {
		return []byte(""), errors.New("write jsonLen error")
	}
	_, serr := buf.WriteString(body.jsonStr)

	if serr != nil {

		return []byte(""), errors.New("write jsonStr error")
	}
	return buf.Bytes(), nil
}
func DecodeBody(by []byte) (body *Body) {
	body = new(Body)
	body.jsonKey = by[0]

	body.jsonLen = binary.LittleEndian.Uint16(by[1:3])
	body.jsonStr = string(by[3:])
	return
}
func (data *PackageData) encodePackageData() ([]byte, error) {
	buf := new(bytes.Buffer)
	headby, err := data.head.encodeHead()
	if err != nil {
		return []byte(""), errors.New("encodeHead  error")
	}
	err = binary.Write(buf, binary.LittleEndian, headby)
	if err != nil {
		return []byte(""), errors.New("write head error")
	}
	bodyby, erro := data.body.encodeBody()
	if erro != nil {
		return []byte(""), errors.New("encodeBody  error")
	}
	err = binary.Write(buf, binary.LittleEndian, bodyby)
	if err != nil {
		return []byte(""), errors.New("write body error")
	}
	return buf.Bytes(), nil
}

func NewHead(msgID uint16, bodyLen uint16) *Head {
	return &Head{msgID, bodyLen, code}
}
func NewBody(jsonKey byte, jsonStr string) *Body {

	return &Body{jsonKey, uint16(len(jsonStr)), jsonStr}
}
func NewPackage(head *Head, body *Body) *PackageData {
	return &PackageData{head, body}
}
