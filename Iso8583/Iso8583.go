package Iso8583

import (
	spec "iso8583/spec"
)

type (
	Iso8583 struct {
		AMessage spec.AMessage
		AHeader  spec.Header
		EchoData map[string]string
	}
)

func NewIso8583(name string) *Iso8583 {
	body, header, echo := spec.NewSpec(name)
	iso8583 := &Iso8583{AMessage: body, AHeader: header, EchoData: echo}
	return iso8583
}
func (msg *Iso8583) GetEchoData() map[string]string {
	return msg.EchoData
}
func (msg *Iso8583) GetSubFieldValue(field, subField int) string {

	if msg.AMessage.Bitmap.IsFieldSet(field) {
		return msg.AMessage.Fields[field].SubFieldValue(subField)
	}

	return ""
}

func (msg *Iso8583) SetSubFieldValue(field, subField int, value string) error {

	fld, err := msg.AMessage.GetField(field)
	if err != nil {
		return err
	}
	fld.SetSubFieldValue(subField, value)
	return nil
}
func (msg *Iso8583) UnpackFull(data []byte, start int) (int, error) {
	header, _ := msg.AHeader.Unpack(data, 0)
	return msg.AMessage.Unpack(data, header, 16)
}
func (msg *Iso8583) ToStringFull(prefix string) string {
	header := "\t*************************** Header Fields ****************************\n"
	fields := "\t*************************** Data Fields ****************************\n"
	output := header + msg.AHeader.String() + fields + msg.AMessage.String()
	return output
}

func (msg *Iso8583) ToMsgFull() []byte {
	headerLength := msg.AHeader.PackedLength()
	bodyLength := msg.AMessage.PackedLength()
	data := make([]byte, headerLength+bodyLength)
	copy(data, msg.AHeader.ToMsg())
	copy(data[headerLength:], msg.AMessage.ToMsg())
	return data

}
func (msg *Iso8583) SetFieldValue(field int, value string) error {
	return msg.AMessage.SetFieldValue(field, value)
}
func (msg *Iso8583) GetFieldValue(field int) string {

	return msg.AMessage.GetFieldValue(field)
}
func (msg *Iso8583) SetHeaderValue(field int, value string) error {
	return msg.AHeader.SetFieldValue(field, value)
}
func (msg *Iso8583) GetHeaderValue(field int) string {

	return msg.AHeader.GetFieldValue(field)
}
func (msg *Iso8583) SetMti(mti string) {
	msg.AHeader.SetFieldValue(0, mti)
}

func (msg *Iso8583) ToFields() (h, b map[string]string) {

	return msg.AHeader.ToFields(), msg.AMessage.ToFields()
}
func (msg *Iso8583) ToCompare(input map[int]string) bool {

	return msg.AMessage.ToCompare(input)
}

func (msg *Iso8583) UnpackH(data []byte, start int) (int, error) {
	return msg.AHeader.Unpack(data, 0)
}
func (msg *Iso8583) ToFieldsH() (h map[string]string) {

	return msg.AHeader.ToFields()
}
