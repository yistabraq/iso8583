package spec

import (
	fieldValidator "iso8583/FieldValidator"
	formatter "iso8583/Formatter"
)

func BPCH2HBody() TemplateDef {

	body := TemplateDef{
		2:   AsciiVar(2, 24, fieldValidator.N()),
		3:   AsciiFixed(6, fieldValidator.N()),
		4:   AsciiNumeric(12),
		5:   AsciiNumeric(12),
		6:   AsciiNumeric(12),
		9:   AsciiFixed(8, fieldValidator.N()),
		11:  AsciiFixed(6, fieldValidator.N()),
		12:  AsciiFixed(12, fieldValidator.N()),
		14:  AsciiFixed(4, fieldValidator.N()),
		15:  AsciiFixed(6, fieldValidator.N()),
		18:  AsciiFixed(4, fieldValidator.N()),
		19:  AsciiFixed(3, fieldValidator.N()),
		21:  AsciiFixed(3, fieldValidator.N()),
		22:  AsciiFixed(12, fieldValidator.Ans()),
		23:  AsciiFixed(2, fieldValidator.N()),
		24:  AsciiFixed(3, fieldValidator.N()),
		32:  AsciiVar(2, 11, fieldValidator.N()),
		35:  AsciiVar(2, 37, fieldValidator.Ans()),
		37:  AsciiFixed(12, fieldValidator.Ans()),
		38:  AsciiFixed(6, fieldValidator.Ans()),
		39:  AsciiFixed(3, fieldValidator.An()),
		41:  AsciiFixed(8, fieldValidator.Ans()),
		42:  AsciiFixed(15, fieldValidator.Ans()),
		43:  AsciiVar(3, 143, fieldValidator.None()),
		45:  AsciiVar(2, 76, fieldValidator.Ans()),
		48:  AsciiVar(3, 999, fieldValidator.None()),
		49:  AsciiFixed(3, fieldValidator.Ans()),
		50:  AsciiFixed(3, fieldValidator.Ans()),
		51:  AsciiFixed(3, fieldValidator.Ans()),
		52:  AsciiVar(2, 16, fieldValidator.An()),
		53:  AsciiVar(2, 16, fieldValidator.N()),
		54:  AsciiVar(3, 999, fieldValidator.An()),
		55:  BinaryVar(3, 255, formatter.Binary()),
		61:  AsciiVar(3, 999, fieldValidator.Ans()),
		64:  BinaryFixed(16),
		91:  AsciiFixed(1, fieldValidator.An()),
		95:  AsciiVar(3, 999, fieldValidator.An()),
		96:  AsciiVar(2, 99, fieldValidator.Ans()),
		100: AsciiVar(2, 99, fieldValidator.N()),
		102: AsciiVar(2, 32, fieldValidator.Ans()),
		103: AsciiVar(2, 32, fieldValidator.Ans()),
		125: AsciiVar(2, 16, fieldValidator.Ans()),
		128: BinaryFixed(16),
	}
	return body
}

func BPCH2HHeader() TemplateDef {

	header := TemplateDef{
		0: AsciiFixed(4, fieldValidator.An()),
	}
	return header
}

func BPCH2HEcho() map[string]string{
	output := make(map[string]string)
	output["H0"] = "0800"
	output["7"] = "#MMDDhhmmss#"
	output["11"] = "#STAN6"
	output["12"] = "#hhmm#"
	output["13"] = "#MMDD#"
	output["70"] = "001"
	return output
}