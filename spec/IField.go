package spec

type IField interface {
	FieldNumber() int
	PackedLength() int
	Value() string
	SetValue(string)
	ToMsg() []byte
	Unpack([]byte, int) (int, error)
	ToString(string) string
	SubFieldValue(int) string
	SetSubFieldValue(int, string)
	ToJSON() string
}
