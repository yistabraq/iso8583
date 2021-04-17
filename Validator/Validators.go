package validator

import (
	fieldValidator "github.com/istabraq/iso8583/FieldValidator"
)


func IsHex(value string) bool {
	v := &fieldValidator.HexFieldValidator{}
	return v.IsValid(value)
}
