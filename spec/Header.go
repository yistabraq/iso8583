package spec

import (
	"bytes"
	"fmt"
	"sort"
	"strconv"
)

//type CreateFieldFunc func(int) IField
type Header struct {
	MsgTemplate         *Template
	Fields              map[int]IField
	json                map[string]string
	CreateFieldCallback CreateFieldFunc
}

func NewHeader(tmpl *Template) *Header {
	header := &Header{MsgTemplate: tmpl, Fields: make(map[int]IField), json: make(map[string]string)}
	header.CreateFieldCallback = header.CreateField

	return header
}
func (header *Header) CreateField(field int) IField {

	if _, ok := header.MsgTemplate.templateDefinition[field]; ok {
		return NewField(field, header.MsgTemplate.templateDefinition[field])
	}

	return nil
}
func (header *Header) PackedLength() int {

	length := 0
	for _, i := range SortKeys(header.MsgTemplate.templateDefinition) {
		if header.Fields[i] != nil {
			length += header.Fields[i].PackedLength()
		}
	}
	return length
}
func (header *Header) Unpack(data []byte, startingOffset int) (int, error) {
	offset := 0
	for _, i := range SortKeys(header.MsgTemplate.templateDefinition) {
		field := header.CreateFieldCallback(i)
		off, err := field.Unpack(data, offset)
		if err != nil {
			return field.FieldNumber(),err
		}
		header.Fields[i] = field
		header.json[fmt.Sprintf("%d", i)] = header.Fields[i].Value()
		offset = off
	}
	return offset, nil
}

func (header *Header) GetField(field int) (IField, error) {

	if header.Fields[field] == nil {
		header.Fields[field] = header.CreateFieldCallback(field)
		return header.Fields[field], nil
	}
	return header.Fields[field], nil
}

func (header *Header) ToString(prefix string) string {
	var buffer bytes.Buffer
	for _, i := range SortKeys(header.MsgTemplate.templateDefinition) {
		_, err := header.GetField(i)
		if err != nil {
			return ""
		}
		buffer.WriteString(header.FieldsToString(i, prefix) + "\n")
	}
	return buffer.String()
}

func (header *Header) FieldsToString(field int, prefix string) string {
	return header.Fields[field].ToString(prefix)
}
func (header *Header) String() string {
	return header.ToString("   ")
}

func (header *Header) ToMsg() []byte {

	packedLength := header.PackedLength()
	data := make([]byte, packedLength)
	offset := 0
	for _, i := range SortKeys(header.MsgTemplate.templateDefinition) {
		field, err := header.GetField(i)
		if err != nil {
			return nil
		}
		copy(data[offset:], field.ToMsg())
		offset += field.PackedLength()
	}

	return data
}

func SortKeys(m TemplateDef) []int {
	var keys []int
	for k, _ := range m {
		keys = append(keys, k)
	}
	//sort.Sort(sort.Reverse(sort.IntSlice(keys)))
	sort.Ints(keys)
	return keys
}

func (header *Header) GetFieldValue(field int) string {

	return header.Fields[field].Value()
}
func (header *Header) SetFieldValue(field int, value string) error {

	header.Fields[field] = header.CreateFieldCallback(field)
	header.Fields[field].SetValue(value)
	header.json[fmt.Sprintf("%d", field)] = header.Fields[field].Value()
	return nil
}
func (header *Header) ClearField(field int) {
	delete(header.Fields, field)
}

func (header *Header) ToFields() map[string]string {

	return header.json
}

func (header *Header) GetResponse() string {
	mti := header.GetFieldValue(0)
	r := mti[2:3]
	i, _ := strconv.Atoi(r)
	i++
	return mti[:2] + (strconv.Itoa(i)) + mti[3:]
}
