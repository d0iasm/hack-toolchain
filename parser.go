package main

import (
	"strings"
)

type COMMAND_TYPE int

const (
	A_COMMAND COMMAND_TYPE = iota
	C_COMMAND COMMAND_TYPE = iota
	L_COMMAND COMMAND_TYPE = iota
	IGNORE    COMMAND_TYPE = iota
)

func commandType(s string) COMMAND_TYPE {
	switch {
	case len(s) == 0 || strings.HasPrefix(s, "//"):
		return IGNORE
	case strings.Contains(s, "@"):
		return A_COMMAND
	case strings.Contains(s, "(") && strings.Contains(s, ")"):
		return L_COMMAND
	default:
		return C_COMMAND
	}
}
