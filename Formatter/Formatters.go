package formatter

type formatterFunc func() IFormatter

var (
	Binary formatterFunc
	Ascii  formatterFunc
	Bcd    formatterFunc
	Ebcdic formatterFunc
)

func init() {
	Binary = NewBinaryFormatter
	Ascii = NewAsciiFormatter
	Bcd = NewBcdFormatter
}
