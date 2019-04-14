package main

func writeArithmetic(arg1 ARG1) string {
	switch arg1 {
	case "add":
		return "M=D+M"
	case "sub":
		return "M=M-D"
	case "neg":
		return "M=-M"
	case "eq":
		return "JEQ"
	case "gt":
		return "JGT"
	case "lt":
		return "JLT"
	case "and":
		return "M=D&M"
	case "or":
		return "M=D|A"
	case "not":
		return "M=!M"
	default:
		return ""
	}
}

func writePushPop(c COMMAND_TYPE, seg string, idx int) string {
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
	default:
		return ""
	}
}
