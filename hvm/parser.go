package main

import (
	"strconv"
	"strings"
)

type COMMAND_TYPE int

const (
	IGNORE COMMAND_TYPE = iota
	C_ARITHMETIC
	C_PUSH
	C_POP
)

type Token struct {
	ct      COMMAND_TYPE // int
	command string
	arg1    string
	arg2    int
}

func tokenize(s string) *Token {
	t := strings.Split(s, " ")
	if len(t) == 1 {
		return &Token{commandType(t[0]), t[0], "", 0}
	} else if len(t) == 3 {
		arg2, _ := strconv.Atoi(t[2])
		return &Token{commandType(t[0]), t[0], t[1], arg2}
	}
	return &Token{IGNORE, t[0], "", 0}
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
		return C_ARITHMETIC
	case s == "push":
		return C_PUSH
	case s == "pop":
		return C_POP
	default:
		return IGNORE
	}
	return IGNORE
}
