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
		line := scanner.Text()
		switch t := commandType(line); t {
		case A_COMMAND, C_COMMAND, L_COMMAND:
			fmt.Fprintln(writer, line+"!")
			break
		case IGNORE:
			break
		}
	}
	err = scanner.Err()
	check(err)
	writer.Flush()
}
