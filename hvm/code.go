package main

func writeArithmetic(c string) string {
	switch c {
	case arithmetic[0]: // add
		return "M=D+M"
	case arithmetic[1]: // sub
		return "M=M-D"
	case arithmetic[2]: // neg
		return "M=-M"
	case arithmetic[3]: // eq
		return "JEQ"
	case arithmetic[4]: // gt
		return "JGT"
	case arithmetic[5]: // lt
		return "JLT"
	case arithmetic[6]: // and
		return "M=D&M"
	case arithmetic[7]: // or
		return "M=D|A"
	case arithmetic[8]: // not
		return "M=!M"
	default:
		return ""
	}
}

func writePushPop(c COMMAND_TYPE, seg string, idx int) {
	// Stack pointer is hold at RAM[0].
	// Stack base starts from RAM[256].
	switch seg {
	case "push":
		return `@0
  M=D
  M=M+1
`
	case "pop":
		return `@0
  D=M
  M=M-1
  `
	}
}
