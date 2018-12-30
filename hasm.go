package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func debug(x interface{}) {
	fmt.Printf("%#v\n", x)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func assemble(s string) (string, bool) {
	switch t := commandType(s); t {
	case A_COMMAND, L_COMMAND:
		return s + " : " + symbol(s), true
	case C_COMMAND:
		return s + " : " + dest(s), true
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

	scanner := bufio.NewScanner(rfile)
	writer := bufio.NewWriter(wfile)
	for scanner.Scan() {
		bin, ok := assemble(scanner.Text())
		if ok {
			fmt.Fprintln(writer, bin)
		}
	}
	err = scanner.Err()
	check(err)
	writer.Flush()
}
