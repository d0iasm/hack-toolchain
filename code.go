package main

func dest(s string) string {
	switch s {
	case "M":
		return "001"
	case "D":
		return "010"
	case "MD":
		return "011"
	case "A":
		return "100"
	case "AM":
		return "101"
	case "AD":
		return "110"
	case "AMD":
		return "111"
	default:
		return "000"
	}
}

func comp(s string) string {
	return compMnemonic(s)
}

func jump(s string) string {
	return jumpMnemonic(s)
}
