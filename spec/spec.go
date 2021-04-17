package spec

func NewSpec(name string) (AMessage, Header, map[string]string) {

	switch name {
	case "umoja":
		return *NewAMessage(NewTemplate(UmojaBody())), *NewHeader(NewTemplate(UmojaHeader())), UmojaEcho()
	case "ghipps":
		return *NewAMessage(NewTemplate(GhippsBody())), *NewHeader(NewTemplate(GhippsHeader())), GhippsEcho()
	case "gim":
		return *NewAMessage(NewTemplate(GimBody())), *NewHeader(NewTemplate(GimHeader())), GimEcho()
	case "bpch2h":
		return *NewAMessage(NewTemplate(BPCH2HBody())), *NewHeader(NewTemplate(BPCH2HHeader())), BPCH2HEcho()
	case "bpccbs":
		return *NewAMessage(NewTemplate(BPCCBSBody())), *NewHeader(NewTemplate(BPCCBSHeader())), BPCCBSEcho()
	case "bpcepos":
		return *NewAMessage(NewTemplate(BPCEPOSBody())), *NewHeader(NewTemplate(BPCEPOSHeader())), BPCEPOSEcho()
	case "kenswitch":
		return *NewAMessage(NewTemplate(KenSwitchBody())), *NewHeader(NewTemplate(KenSwitchHeader())), kenSwitchEcho()
	default:
		return *NewAMessage(NewTemplate(UmojaBody())), *NewHeader(NewTemplate(UmojaHeader())), nil
	}

}
