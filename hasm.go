package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func debug(x ...interface{}) {
	fmt.Println("----- debug start -----")
	for i := 0; i < len(x); i++ {
		fmt.Printf("%#v\n", x[i])
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
		return "111" + comp(cmn) + dest(dmn) + jump(jmn), true
	case L_COMMAND:
		// TODO: Implement to add a symbol into a symbol table
	case IGNORE:
	}
	return "", false
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
	st.addEntry("testtesttest", 99999999999)
	debug(st.getAddress("R5"))
	debug(st.getAddress("hoge"))
	debug(st)

	scanner := bufio.NewScanner(rfile)
	writer := bufio.NewWriter(wfile)
	for scanner.Scan() {
		bin, isPrint := assemble(scanner.Text())
		if isPrint {
			fmt.Fprintln(writer, bin)
		}
	}
	err = scanner.Err()
	check(err)
	writer.Flush()
}
