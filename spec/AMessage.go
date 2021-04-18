package spec

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

type CreateFieldFunc func(int) IField

type AMessage struct {
	Bitmap              *Bitmap
	MsgTemplate         *Template
	Fields              map[int]IField
	json                map[string]string
	CreateFieldCallback CreateFieldFunc
}

func NewAMessage(tmpl *Template) *AMessage {
	msg := &AMessage{MsgTemplate: tmpl, Fields: make(map[int]IField), Bitmap: NewBitmap(tmpl.BitmapFormatter), json: make(map[string]string)}
	msg.CreateFieldCallback = msg.CreateField

	return msg
}

func (msg *AMessage) PackedLength() int {

	length := msg.Bitmap.PackedLength()
	for i := 2; i < 128; i++ {
		if msg.Bitmap.IsFieldSet(i) {
			length += msg.Fields[i].PackedLength()
		}
	}
	return length
}

func (msg *AMessage) ClearField(field int) {
	msg.Bitmap.SetField(field, false)
	delete(msg.Fields, field)
}

func (msg *AMessage) IsFieldSet(field int) bool {
	return msg.Bitmap.IsFieldSet(field)
}

func (msg *AMessage) ToMsg() []byte {

	packedLength := msg.PackedLength()
	data := make([]byte, packedLength)
	bmap := msg.Bitmap.ToMsg()
	copy(data, bmap)
	offset := msg.Bitmap.PackedLength()
	for i := 2; i < 128; i++ {
		if msg.Bitmap.IsFieldSet(i) {
			field := msg.Fields[i]
			copy(data[offset:], field.ToMsg())
			offset += field.PackedLength()
		}
	}

	return data
}

func (msg *AMessage) String() string {
	return msg.ToString("   ")
}

func (msg *AMessage) ToString(prefix string) string {
	var buffer bytes.Buffer
	bitmap := "[ "
	for i := 2; i < 128; i++ {
		if msg.Bitmap.IsFieldSet(i) {
			bitmap += strconv.Itoa(i)
			buffer.WriteString(msg.FieldsToString(i, prefix) + "\n")
			bitmap += ", "
		}
	}

	b := fmt.Sprintf("%s[Fixed    b        16 0016] 001\t%s]\n", prefix, bitmap)
	return b + buffer.String()
}

func (msg *AMessage) FieldsToString(field int, prefix string) string {

	return msg.Fields[field].ToString(prefix)
}

func (msg *AMessage) CreateField(field int) IField {

	if _, ok := msg.MsgTemplate.templateDefinition[field]; ok {
		return NewField(field, msg.MsgTemplate.templateDefinition[field])
	}

	return nil
}

func (msg *AMessage) GetField(field int) (IField, error) {

	_, ok := msg.Fields[field]
	if (!msg.Bitmap.IsFieldSet(field)) || (!ok) {
		if msg.Fields[field] = msg.CreateFieldCallback(field); msg.Fields[field] != nil {
			msg.Bitmap.SetField(field, true)
		} else {
			return nil, errors.New(fmt.Sprintf("Unable to create field number %d. Possibly because template does not have a defination for the field",
				field))
		}
	}

	return msg.Fields[field], nil
}

func (msg *AMessage) Unpack(data []byte, startingOffset int, lengthOfBitmap int) (int, error) {
	offset := msg.Bitmap.Unpack(data, startingOffset, lengthOfBitmap)
	for i := 2; i < 128; i++ {
		if msg.Bitmap.IsFieldSet(i) {
			field, err := msg.GetField(i)
			if err != nil {
				return i, err
			}
			offset, err = field.Unpack(data, offset)
			if err != nil {
				fmt.Println(err)
				return field.FieldNumber(), err
			}
			msg.json[fmt.Sprintf("%03d", i)] = msg.Fields[i].Value()
		}
	}
	return offset, nil
}

func (msg *AMessage) GetFieldValue(field int) string {

	if msg.Bitmap.IsFieldSet(field) {
		return msg.Fields[field].Value()
	}
	return ""
}

func (msg *AMessage) SetFieldValue(field int, value string) error {

	if value == "" {
		msg.ClearField(field)
		return nil
	}

	fld, err := msg.GetField(field)
	if err != nil {
		return err
	}

	ok,out := fld.SetValue(value)
	if !ok {
		return fmt.Errorf(out)
	}
	msg.json[fmt.Sprintf("%03d", field)] = msg.Fields[field].Value()
	return nil
}

func (msg *AMessage) ToFields() map[string]string {
	return msg.json
}
func (msg *AMessage) ToCompare(input map[int]string) bool {

	return reflect.DeepEqual(input, msg.json)
}
func (msg *AMessage) ToJSON() string {
	var buffer bytes.Buffer
	for i := 1; i < 128; i++ {
		if msg.Bitmap.IsFieldSet(i) {
			buffer.WriteString(msg.FieldsToJson(i))

		}
	}

	return buffer.String()
}
func (msg *AMessage) FieldsToJson(field int) string {
	defer func() {
		if err := recover(); err != nil {
			if field != 1 {
				fmt.Println("panic occurred this field ", field, " : ", err)
			}

		}
	}()
	return msg.Fields[field].ToJSON()
}
