package spec

import (
	"encoding/hex"
	"errors"
	"fmt"
	"strconv"
)

type Mti struct {
	value  string
	encode string
}

func NewMti(value string, encode string) *Mti {
	mti := &Mti{value: value, encode: encode}
	return mti
}
func (m *Mti) SetMti(value string, encode string) {
	m.value = value
	m.encode = encode
}
func (m *Mti) encodeMti() ([]byte, error) {
	if m.value == "" {
		return nil, errors.New("MTI is required")
	}
	if len(m.value) != 4 {
		return nil, errors.New("MTI is invalid")
	}

	// check MTI, it must contain only digits
	if _, err := strconv.Atoi(m.value); err != nil {
		return nil, errors.New("MTI is invalid")
	}

	switch m.encode {
	case "BCD":
		return bcd([]byte(m.value)), nil
	default:
		return []byte(m.value), nil
	}
}
func (m *Mti) decodeMti(raw []byte) int {
	mtiLen := 4
	if m.encode == "BCD" {
		mtiLen = 2
	}
	if len(raw) < mtiLen {
		fmt.Errorf("Error len data")
		return 0
	}

	var mti string
	switch m.encode {
	case "ASCII":
		mti = string(raw[:mtiLen])
	case "BCD":
		mti = string(bcd2Ascii(raw[:mtiLen]))
	default:
		return 0
	}
	m.value = mti

	return mtiLen
}

func bcd2Ascii(data []byte) []byte {
	out := make([]byte, len(data)*2)
	n := hex.Encode(out, data)
	return out[:n]
}
func bcd(data []byte) []byte {
	out := make([]byte, len(data)/2+1)
	n, err := hex.Decode(out, data)
	if err != nil {
		panic(err.Error())
	}
	return out[:n]
}
