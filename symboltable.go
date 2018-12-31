package main

type ST struct {
	table   map[string]int
	varAddr int
}

func initST() ST {
	st := ST{make(map[string]int), 16}
	st.table["SP"] = 0
	st.table["LCL"] = 1
	st.table["ARG"] = 2
	st.table["THIS"] = 3
	st.table["THAT"] = 4
	st.table["R0"] = 0
	st.table["R1"] = 1
	st.table["R2"] = 2
	st.table["R3"] = 3
	st.table["R4"] = 4
	st.table["R5"] = 5
	st.table["R6"] = 6
	st.table["R7"] = 7
	st.table["R8"] = 8
	st.table["R9"] = 9
	st.table["R10"] = 10
	st.table["R11"] = 11
	st.table["R12"] = 12
	st.table["R13"] = 13
	st.table["R14"] = 14
	st.table["R15"] = 15
	st.table["SCREEN"] = 16384
	st.table["KBD"] = 24576
	return st
}

func (st ST) addEntry(sym string, addr int) {
	st.table[sym] = addr
}

func (st ST) getAddress(sym string) (int, bool) {
	v, ok := st.table[sym]
	return v, ok
}
