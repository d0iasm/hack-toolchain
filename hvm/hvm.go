package main

import (
	"fmt"
	"os"
	"regexp"
)

const (
	SP   int = 0 // RAM[0] Stack pointer: points to the next topmost location in the stack.
	LCL  int = 1 // RAM[1] Points to the base of the current VM function's local segment.
	ARG  int = 2 // RAM[2] Points to the base of the current VM function 's argument segment.
	THIS int = 3 // RAM[3] Points to the base of the current this segment (within the heap).
	THAT int = 4 // RAM[4] Points to the base of the current that segment (within the heap).
	TEMP int = 5 // RAM[5-12] Hold the contents of the temp segment
// RAM[13-15] Can be used by the VM implementation as general-purpose registers.
)

type Stack struct {
	constant []int
}

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

func remove(s string) string {
	comment := regexp.MustCompile(`//.*`)
	return comment.ReplaceAllString(s, "")
}

func main() {
	filename := os.Args[1]
	cw := createCodeWriter(filename)

	for cw.scanner.Scan() {
		text := remove(cw.scanner.Text())
		token := tokenize(text)
		fmt.Fprintln(cw.writer, token)

		fmt.Println(cw.scanner.Text())
		fmt.Println(token)
	}
	cw.close()
}
