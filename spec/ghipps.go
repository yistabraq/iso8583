package spec

import (
	formatter "github.com/istabraq/iso8583/Formatter"

	fieldValidator "github.com/istabraq/iso8583/FieldValidator"
)

func GhippsBody() TemplateDef {

	body := TemplateDef{
		2:   AsciiVar(2, 19, fieldValidator.N()),
		3:   AsciiFixed(6, fieldValidator.N()),
		4:   AsciiNumeric(12),
		5:   AsciiNumeric(12),
		6:   AsciiNumeric(12),
		7:   AsciiFixed(10, fieldValidator.N()),
		9:   AsciiFixed(8, fieldValidator.N()),
		10:  AsciiFixed(8, fieldValidator.N()),
		11:  AsciiFixed(6, fieldValidator.N()),
		12:  AsciiFixed(12, fieldValidator.N()),
		14:  AsciiFixed(4, fieldValidator.N()),
		15:  AsciiFixed(6, fieldValidator.N()),
		16:  AsciiFixed(4, fieldValidator.N()),
		18:  AsciiFixed(4, fieldValidator.N()),
		19:  AsciiFixed(3, fieldValidator.N()),
		21:  AsciiFixed(3, fieldValidator.N()),
		22:  AsciiFixed(12, fieldValidator.An()),
		23:  AsciiFixed(3, fieldValidator.N()),
		24:  AsciiFixed(3, fieldValidator.N()),
		25:  AsciiFixed(4, fieldValidator.N()),
		27:  AsciiFixed(1, fieldValidator.N()),
		30:  AsciiFixed(24, fieldValidator.N()),
		32:  AsciiVar(2, 11, fieldValidator.N()),
		33:  AsciiVar(2, 11, fieldValidator.N()),
		35:  AsciiVar(2, 37, fieldValidator.Ans()),
		37:  AsciiFixed(12, fieldValidator.None()),
		38:  AsciiFixed(6, fieldValidator.An()),
		39:  AsciiFixed(3, fieldValidator.N()),
		41:  AsciiFixed(8, fieldValidator.Ans()),
		42:  AsciiFixed(23, fieldValidator.Ans()),
		43:  AsciiVar(2, 99, fieldValidator.None()),
		45:  AsciiVar(2, 76, fieldValidator.Ans()),
		46:  AsciiVar(3, 204, fieldValidator.Ans()),
		48:  AsciiVar(3, 999, fieldValidator.None()),
		49:  AsciiFixed(3, fieldValidator.N()),
		50:  AsciiFixed(3, fieldValidator.N()),
		51:  AsciiFixed(3, fieldValidator.N()),
		52:  AsciiFixed(16, fieldValidator.An()),
		53:  AsciiVar(2, 99, fieldValidator.N()),
		54:  AsciiVar(3, 120, fieldValidator.Ans()),
		55:  BinaryVar(3, 255, formatter.Binary()),
		56:  AsciiVar(2, 35, fieldValidator.N()),
		60:  AsciiVar(3, 999, fieldValidator.Ans()),
		61:  AsciiVar(3, 999, fieldValidator.Ans()),
		62:  AsciiVar(3, 999, fieldValidator.Ans()),
		73:  AsciiFixed(6, fieldValidator.An()),
		74:  AsciiFixed(10, fieldValidator.N()),
		75:  AsciiFixed(10, fieldValidator.N()),
		76:  AsciiFixed(10, fieldValidator.N()),
		77:  AsciiFixed(10, fieldValidator.N()),
		83:  AsciiFixed(10, fieldValidator.N()),
		84:  AsciiFixed(10, fieldValidator.N()),
		86:  AsciiFixed(16, fieldValidator.N()),
		87:  AsciiFixed(16, fieldValidator.N()),
		88:  AsciiFixed(16, fieldValidator.N()),
		89:  AsciiFixed(16, fieldValidator.N()),
		93:  AsciiVar(2, 11, fieldValidator.N()),
		94:  AsciiVar(2, 11, fieldValidator.N()),
		97:  AsciiFixed(17, fieldValidator.Rev87AmountValidator()),
		99:  AsciiVar(2, 11, fieldValidator.N()),
		101: AsciiVar(2, 17, fieldValidator.Ans()),
		102: AsciiVar(2, 28, fieldValidator.Ans()),
		103: AsciiVar(2, 28, fieldValidator.Ans()),
		105: AsciiFixed(16, fieldValidator.N()),
		106: AsciiFixed(16, fieldValidator.N()),
		107: AsciiFixed(16, fieldValidator.N()),
		108: AsciiFixed(16, fieldValidator.N()),
		128: BinaryFixed(16),
	}
	return body
}

func GhippsHeader() TemplateDef {

	header := TemplateDef{
		-2: AsciiFixed(3, fieldValidator.A()),
		-1: AsciiFixed(8, fieldValidator.N()),
		0:  AsciiFixed(4, fieldValidator.An()),
	}
	return header
}
func GhippsEcho() map[string]string{
	output := make(map[string]string)
	output["H0"] = "0800"
	output["7"] = "#MMDDhhmmss#"
	output["11"] = "#STAN6"
	output["12"] = "#hhmm#"
	output["13"] = "#MMDD#"
	output["70"] = "001"
	return output
}