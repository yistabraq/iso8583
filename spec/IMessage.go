package spec

type IMessage interface {
	ToMsg() []byte
	Unpack([]byte, int, int) (int, error)
}
