package main

type COMMAND_TYPE int

const (
	IGNORE = iota
	C_ARTHMETIC
	C_PUSH
	C_POP
)

var arthmetic = []string{"add", "sub", "neg", "eq", "gt", "lt", "and", "or", "not"}

func contains(val string, arr []string) bool {
	for _, elm := range arr {
		if elm == val {
			return true
		}
	}
	return false
}

func commandType(tokens []string) COMMAND_TYPE {
	switch {
	case contains(tokens[0], arthmetic):
		return C_ARTHMETIC
	case tokens[0] == "push":
		return C_PUSH
	case tokens[0] == "pop":
		return C_POP
	default:
		return IGNORE
	}
	return IGNORE
}

