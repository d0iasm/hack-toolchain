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
	s map[string][]int
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

func (s *Stack) initStacks() {
	for i := 0; i < 32767; i++ {
		s.s["constant"] = append(s.s["constant"], i)
	}
}

func main() {
	s := &Stack{make(map[string][]int)}
	s.initStacks()

	filename := os.Args[1]
	cw := createCodeWriter(filename)

	for cw.scan() {
		t := tokenize(remove(cw.text()))
		switch t.ct {
		case C_ARITHMETIC:
			cw.writeArithmetic(t.command)
		case C_PUSH:
		case C_POP:
			cw.writePushPop(t.ct, t.arg1, t.arg2)
		}

		fmt.Println(cw.text())
		fmt.Println(t)
	}

	cw.close()
}
