package formatter

import (
	"encoding/hex"
	"errors"
	validator "iso8583/Validator"
	"strings"
)

type BinaryFormatter struct {
}

func NewBinaryFormatter() IFormatter {
	return &BinaryFormatter{}
}

func (bf *BinaryFormatter) GetBytes(value string) ([]byte, error) {

	if !validator.IsHex(value) {
		return nil, errors.New("value is not in hexadecimal format")
	}

	return hex.DecodeString(value)
}

func (bf *BinaryFormatter) GetString(value []byte) string {
	return strings.ToUpper(hex.EncodeToString(value))
}

func (bf *BinaryFormatter) GetPackedLength(unpackedLength int) int {
	return unpackedLength / 2
}
