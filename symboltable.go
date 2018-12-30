package main

type STable map[string]int

func initSTable() STable {
	st := make(STable)
	st["SP"] = 0
	st["LCL"] = 1
	st["ARG"] = 2
	st["THIS"] = 3
	st["THAT"] = 4
	st["R0"] = 0
	st["R1"] = 1
	st["R2"] = 2
	st["R3"] = 3
	st["R4"] = 4
	st["R5"] = 5
	st["R6"] = 6
	st["R7"] = 7
	st["R8"] = 8
	st["R9"] = 9
	st["R10"] = 10
	st["R11"] = 11
	st["R12"] = 12
	st["R13"] = 13
	st["R14"] = 14
	st["R15"] = 15
	st["SCREEN"] = 16384
	st["KBD"] = 24576
	return st
}

func (st STable) addEntry(sym string, addr int) {
	st[sym] = addr
}

func (st STable) getAddress(sym string) (int, bool) {
	v, ok := st[sym]
	return v, ok
}
