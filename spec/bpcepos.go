package spec

import (
	fieldValidator "github.com/istabraq/iso8583/FieldValidator"
)

func BPCEPOSBody() TemplateDef {

	body := TemplateDef{
		2:   AsciiVar(2, 19, fieldValidator.N()),
		3:   AsciiFixed(6, fieldValidator.N()),
		4:   AsciiNumeric(12),
		5:   AsciiFixed(12, fieldValidator.Rev87AmountValidator()),
		6:   AsciiNumeric(12),
		7:   AsciiFixed(10, fieldValidator.N()),
		10:  AsciiFixed(8, fieldValidator.N()),
		11:  AsciiFixed(6, fieldValidator.N()),
		12:  AsciiFixed(12, fieldValidator.N()),
		14:  AsciiFixed(4, fieldValidator.N()),
		15:  AsciiFixed(6, fieldValidator.N()),
		18:  AsciiFixed(4, fieldValidator.N()),
		22:  AsciiFixed(3, fieldValidator.N()),
		23:  AsciiFixed(3, fieldValidator.N()),
		24:  AsciiFixed(3, fieldValidator.N()),
		25:  AsciiFixed(2, fieldValidator.N()),
		28:  AsciiFixed(9, fieldValidator.An()),
		34:  AsciiVar(2, 19, fieldValidator.N()),
		35:  AsciiVar(2, 37, fieldValidator.None()),
		37:  AsciiFixed(12, fieldValidator.Ans()),
		38:  AsciiFixed(6, fieldValidator.An()),
		39:  AsciiFixed(2, fieldValidator.An()),
		41:  AsciiFixed(8, fieldValidator.Ans()),
		42:  AsciiFixed(15, fieldValidator.Ans()),
		43:  AsciiFixed(40, fieldValidator.None()),
		47:  AsciiVar(3, 999, fieldValidator.None()),
		48:  AsciiVar(3, 999, fieldValidator.None()),
		49:  AsciiFixed(3, fieldValidator.N()),
		50:  AsciiFixed(3, fieldValidator.N()),
		51:  AsciiFixed(3, fieldValidator.N()),
		52:  AsciiFixed(16, fieldValidator.An()),
		53:  BinaryFixed(16),
		54:  AsciiVar(3, 123, fieldValidator.Ans()),
		60:  AsciiVar(3, 999, fieldValidator.Ans()),
		62:  AsciiVar(3, 999, fieldValidator.Ans()),
		64:  BinaryFixed(16),
		70:  AsciiFixed(3, fieldValidator.N()),
		90:  AsciiFixed(20, fieldValidator.N()),
		91:  AsciiFixed(1, fieldValidator.An()),
		95:  AsciiFixed(12, fieldValidator.N()),
		102: AsciiVar(2, 99, fieldValidator.Ans()),
		103: AsciiVar(2, 99, fieldValidator.Ans()),
		128: BinaryFixed(16),
	}
	return body
}

func BPCEPOSHeader() TemplateDef {

	header := TemplateDef{
		0: AsciiFixed(4, fieldValidator.An()),
	}
	return header
}

func BPCEPOSEcho() map[string]string{
	output := make(map[string]string)
	output["H0"] = "0800"
	output["7"] = "#MMDDhhmmss#"
	output["11"] = "#STAN6"
	output["12"] = "#hhmm#"
	output["13"] = "#MMDD#"
	output["70"] = "001"
	return output
}