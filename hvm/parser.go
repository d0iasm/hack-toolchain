package main

import (
	"strconv"
	"strings"
)

type COMMAND_TYPE int
type ARG1 string
type ARG2 int

const (
	IGNORE COMMAND_TYPE = iota
	C_ARTHMETIC
	C_PUSH
	C_POP
)

const (
	NONE ARG2 = iota
	CONSTANT
)

type Token struct {
	ct    COMMAND_TYPE // int
	arg1  ARG1         // string
	arg2  ARG2         // int
	immed int
}

func createToken(ct COMMAND_TYPE, arg1 ARG1, arg2 ARG2, immed int) *Token {
	return &Token{ct, arg1, arg2, immed}
}

func tokenize(s string) *Token {
	t := strings.Split(s, " ")
	if len(t) == 1 {
		return createToken(
			commandType(t[0]),
			ARG1(t[0]), NONE, 0,
		)
	} else if len(t) == 3 {
		immed, _ := strconv.Atoi(t[2])
		return createToken(
			commandType(t[0]),
			ARG1(t[0]), arg2(t[1]), immed,
		)
	}
	return createToken(IGNORE, "", NONE, 0)
}

func arg2(s string) ARG2 {
	switch s {
	case "constant":
		return CONSTANT
	default:
		return NONE
	}
}
func isArithmetic(s string) bool {
	a := []string{"add", "sub", "neg", "eq", "gt", "lt", "and", "or", "not"}
	for _, e := range a {
		if e == s {
			return true
		}
	}
	return false
}

func commandType(s string) COMMAND_TYPE {
	switch {
	case isArithmetic(s):
		return C_ARTHMETIC
	case s == "push":
		return C_PUSH
	case s == "pop":
		return C_POP
	default:
		return IGNORE
	}
	return IGNORE
}
