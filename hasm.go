package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func debug(x ...interface{}) {
	for i, v := range x {
		fmt.Printf("%v: %#v\n", i, v)
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func assemble(s string) (string, bool) {
	switch t := commandType(s); t {
	case A_COMMAND:
		sym := symbol(s)
		return "0" + value(sym), true
	case C_COMMAND:
		dmn := destMnemonic(s)
		cmn := compMnemonic(s)
		jmn := jumpMnemonic(s)
		debug("dest: " + dmn + "comp: " + cmn + "jump: " + jmn)
		return "111" + comp(cmn) + dest(dmn) + jump(jmn), true
	case L_COMMAND:
	case IGNORE:
	}
	return "", false
}

func buildTable(st STable, s string, addr int) int {
	switch t := commandType(s); t {
	case L_COMMAND:
		st.addEntry(symbol(s), addr+1)
		return addr
	case A_COMMAND, C_COMMAND:
		return addr + 1
	case IGNORE:
	}
	return addr
}

func main() {
	fname := os.Args[1]
	fsplit := strings.Split(fname, ".")

	rfile, err := os.Open(fname)
	check(err)
	defer rfile.Close()

	wfile, err := os.Create(fsplit[0] + ".hack")
	check(err)
	defer wfile.Close()

	st := initSTable()
	addr := 0
	text := ""
	scanner := bufio.NewScanner(rfile)
	for scanner.Scan() {
		text = remove(scanner.Text())
		addr = buildTable(st, text, addr)
	}
	err = scanner.Err()
	check(err)

	// debug(addr)
	// debug(st)

	scanner = bufio.NewScanner(rfile)
	writer := bufio.NewWriter(wfile)
	for scanner.Scan() {
		text = remove(scanner.Text())
		debug(text)
		bin, isPrint := assemble(text)
		if isPrint {
			fmt.Fprintln(writer, bin)
		}
	}
	err = scanner.Err()
	check(err)
	writer.Flush()
}
