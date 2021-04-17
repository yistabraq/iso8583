package spec

import (
	fieldValidator "github.com/istabraq/iso8583/FieldValidator"
	formatter "github.com/istabraq/iso8583/Formatter"
)

func KenSwitchBody() TemplateDef {
	field127Template := TemplateDef{
		2:  AsciiVar(2, 32, fieldValidator.Ans()),
		3:  AsciiFixed(48, fieldValidator.None()),
		4:  AsciiFixed(22, fieldValidator.Ans()),
		5:  AsciiFixed(73, fieldValidator.Ans()),
		6:  AsciiFixed(2, fieldValidator.None()),
		7:  AsciiVar(2, 70, fieldValidator.Ans()),
		8:  AsciiVar(3, 999, fieldValidator.Ans()),
		9:  AsciiVar(3, 255, fieldValidator.Ans()),
		10: AsciiFixed(3, fieldValidator.N()),
		11: AsciiVar(2, 32, fieldValidator.Ans()),
		12: AsciiVar(2, 25, fieldValidator.Ans()),
		13: AsciiFixed(17, fieldValidator.Ans()),
		14: AsciiFixed(8, fieldValidator.Ans()),
		15: AsciiVar(2, 29, fieldValidator.Ans()),
		16: AsciiFixed(1, fieldValidator.Ans()),
		17: AsciiVar(2, 50, fieldValidator.Ans()),
		18: AsciiVar(2, 50, fieldValidator.Ans()),
		19: AsciiFixed(31, fieldValidator.Ans()),
		20: AsciiFixed(8, fieldValidator.N()),
		21: AsciiVar(2, 12, fieldValidator.Ans()),
		22: AsciiVar(5, 99999, fieldValidator.Ans()),
		23: AsciiFixed(253, fieldValidator.None()),
		24: AsciiVar(2, 28, fieldValidator.Ans()),
		25: AsciiVar(4, 9999, fieldValidator.Ans()),
		26: AsciiVar(2, 20, fieldValidator.Ans()),
		27: AsciiFixed(1, fieldValidator.Ans()),
		28: AsciiFixed(4, fieldValidator.N()),
		29: BinaryFixed(40),
		30: AsciiFixed(1, fieldValidator.Ans()),
		31: AsciiVar(2, 11, fieldValidator.Ans()),
		32: BinaryVar(2, 33, formatter.Ascii()),
		33: AsciiFixed(4, fieldValidator.N()),
		34: AsciiFixed(2, fieldValidator.N()),
		35: AsciiVar(2, 11, fieldValidator.Ans()),
		39: AsciiFixed(2, fieldValidator.An()),
	}

	body := TemplateDef{
		2:   AsciiVar(2, 19, fieldValidator.N()),
		3:   AsciiFixed(6, fieldValidator.N()),
		4:   AsciiNumeric(12),
		5:   AsciiNumeric(12),
		7:   AsciiFixed(10, fieldValidator.N()),
		9:   AsciiFixed(8, fieldValidator.N()),
		11:  AsciiFixed(6, fieldValidator.N()),
		12:  AsciiFixed(6, fieldValidator.N()),
		13:  AsciiFixed(4, fieldValidator.N()),
		14:  AsciiFixed(4, fieldValidator.N()),
		15:  AsciiFixed(4, fieldValidator.N()),
		16:  AsciiFixed(4, fieldValidator.N()),
		18:  AsciiFixed(4, fieldValidator.N()),
		22:  AsciiFixed(3, fieldValidator.N()),
		23:  AsciiFixed(3, fieldValidator.N()),
		25:  AsciiFixed(2, fieldValidator.N()),
		26:  AsciiFixed(2, fieldValidator.N()),
		27:  AsciiFixed(1, fieldValidator.N()),
		28:  AsciiFixed(9, fieldValidator.Rev87AmountValidator()),
		29:  AsciiFixed(9, fieldValidator.Rev87AmountValidator()),
		30:  AsciiFixed(9, fieldValidator.Rev87AmountValidator()),
		31:  AsciiFixed(9, fieldValidator.Rev87AmountValidator()),
		32:  AsciiVar(2, 11, fieldValidator.N()),
		33:  AsciiVar(2, 11, fieldValidator.N()),
		35:  AsciiVar(2, 37, fieldValidator.Ans()),
		37:  AsciiFixed(12, fieldValidator.N()),
		38:  AsciiFixed(6, fieldValidator.An()),
		39:  AsciiFixed(2, fieldValidator.An()),
		40:  AsciiFixed(3, fieldValidator.N()),
		41:  AsciiFixed(8, fieldValidator.Ans()),
		42:  AsciiFixed(15, fieldValidator.Ans()),
		43:  AsciiFixed(40, fieldValidator.Ans()),
		44:  AsciiVar(2, 25, fieldValidator.Ans()),
		45:  AsciiVar(2, 76, fieldValidator.Ans()),
		48:  AsciiVar(3, 999, fieldValidator.None()),
		49:  AsciiFixed(3, fieldValidator.N()),
		50:  AsciiFixed(3, fieldValidator.N()),
		52:  BinaryFixed(16),
		53:  BinaryFixed(48),
		54:  AsciiVar(3, 120, fieldValidator.Ans()),
		55:  BinaryVar(2, 255, formatter.Binary()),
		56:  AsciiVar(3, 4, fieldValidator.N()),
		57:  AsciiVar(3, 3, fieldValidator.N()),
		58:  AsciiVar(3, 11, fieldValidator.Anp()),
		59:  AsciiVar(3, 255, fieldValidator.None()),
		66:  AsciiFixed(1, fieldValidator.N()),
		67:  AsciiFixed(2, fieldValidator.N()),
		70:  AsciiFixed(3, fieldValidator.N()),
		73:  AsciiFixed(6, fieldValidator.N()),
		74:  AsciiFixed(10, fieldValidator.N()),
		75:  AsciiFixed(10, fieldValidator.N()),
		76:  AsciiFixed(10, fieldValidator.N()),
		77:  AsciiFixed(10, fieldValidator.N()),
		78:  AsciiFixed(10, fieldValidator.N()),
		79:  AsciiFixed(10, fieldValidator.N()),
		80:  AsciiFixed(10, fieldValidator.N()),
		81:  AsciiFixed(10, fieldValidator.N()),
		82:  AsciiFixed(12, fieldValidator.N()),
		83:  AsciiFixed(12, fieldValidator.N()),
		84:  AsciiFixed(12, fieldValidator.N()),
		85:  AsciiFixed(12, fieldValidator.N()),
		86:  AsciiFixed(16, fieldValidator.N()),
		87:  AsciiFixed(16, fieldValidator.N()),
		88:  AsciiFixed(16, fieldValidator.N()),
		89:  AsciiFixed(16, fieldValidator.N()),
		90:  AsciiFixed(42, fieldValidator.N()),
		91:  AsciiFixed(1, fieldValidator.An()),
		95:  AsciiFixed(42, fieldValidator.Ans()),
		97:  AsciiFixed(17, fieldValidator.Rev87AmountValidator()),
		98:  AsciiFixed(25, fieldValidator.Ans()),
		100: AsciiVar(2, 11, fieldValidator.Ansp()),
		101: AsciiVar(2, 17, fieldValidator.Ans()),
		102: AsciiVar(2, 28, fieldValidator.Ans()),
		103: AsciiVar(2, 28, fieldValidator.Ans()),
		118: AsciiVar(3, 30, fieldValidator.N()),
		119: AsciiVar(3, 10, fieldValidator.N()),
		123: AsciiVar(3, 15, fieldValidator.None()),
		125: AsciiVar(3, 40, fieldValidator.None()),
		127: CompositeField(6, 999999, 16, NewTemplate(field127Template)),
	}
	return body
}

func KenSwitchHeader() TemplateDef {

	header := TemplateDef{
		0: AsciiFixed(4, fieldValidator.N()),
	}
	return header
}

func kenSwitchEcho() map[string]string{
	output := make(map[string]string)
	output["H0"] = "0800"
	output["7"] = "#MMDDhhmmss#"
	output["11"] = "#STAN6"
	output["12"] = "#hhmm#"
	output["13"] = "#MMDD#"
	output["70"] = "001"
	return output
}
