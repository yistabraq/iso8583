package spec

import (
	fieldValidator "github.com/istabraq/iso8583/FieldValidator"
	formatter "github.com/istabraq/iso8583/Formatter"
	lengthFormatters "github.com/istabraq/iso8583/LengthFormatters"
)

type IFieldDescriptor interface {
	Adjuster() Adjuster
	Formatter() formatter.IFormatter
	LengthFormatter() lengthFormatters.ILengthFormatter
	Validator() fieldValidator.IFieldValidator
	IsComposite() bool
	CompositeTemplate() *Template
	Display(string, string, string) string
	GetPackedLength(string) int
	Pack(int, string) ([]byte, error)
	Unpack(int, []byte, int) (string, int, error)
	GetLengthOfBitmap() int
}
