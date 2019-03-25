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

func main() {
	fname := os.Args[1]
	fsplit := strings.Split(fname, ".")

	rfile, err := os.Open(fname)
	check(err)
	defer rfile.Close()

	wfile, err := os.Create(fsplit[0] + ".asm")
	check(err)
	defer wfile.Close()

        scanner := bufio.NewScanner(rfile)
	writer := bufio.NewWriter(wfile)
        text := ""
	for scanner.Scan() {
		text = scanner.Text()
		fmt.Fprintln(writer, text)
	}
	err = scanner.Err()
	check(err)
	writer.Flush()
}
