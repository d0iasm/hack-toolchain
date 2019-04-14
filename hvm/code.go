package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type CodeWriter struct {
	input    *os.File
	output   *os.File
	basename string
	scanner  *bufio.Scanner
	writer   *bufio.Writer
}

func createCodeWriter(inname string) *CodeWriter {
	s := strings.Split(inname, ".")
	basename, extension := s[0], s[1]

	if extension != "vm" {
		panic("Input file should have .vm extension.")
	}

	input, err := os.Open(inname)
	check(err)

	outname := basename + ".asm"

	output, err := os.Create(outname)
	check(err)

	scanner := bufio.NewScanner(input)
	writer := bufio.NewWriter(output)

	return &CodeWriter{input, output, basename, scanner, writer}
}

func (cw *CodeWriter) close() {
	err := cw.scanner.Err()
	fmt.Println(cw)
	check(err)
	cw.writer.Flush()

	cw.input.Close()
	cw.output.Close()
}

func (cw *CodeWriter) text() string {
	return cw.scanner.Text()
}

func (cw *CodeWriter) scan() bool {
	return cw.scanner.Scan()
}

func (cw *CodeWriter) writeArithmetic(arg1 ARG1) {
	command := ""
	switch arg1 {
	case "add":
		command = "M=D+M"
	case "sub":
		command = "M=M-D"
	case "neg":
		command = "M=-M"
	case "eq":
		command = "JEQ"
	case "gt":
		command = "JGT"
	case "lt":
		command = "JLT"
	case "and":
		command = "M=D&M"
	case "or":
		command = "M=D|A"
	case "not":
		command = "M=!M"
	}
	if command != "" {
		fmt.Fprintln(cw.writer, command)
	}
}

func (cw *CodeWriter) writePushPop(ct COMMAND_TYPE, seg string, idx int) {
	// Stack pointer(SP) is hold at RAM[0].
	// Stack base starts from RAM[256].
	command := ""

	switch ct {
	case C_PUSH:
		command = `@SP
  A=M
  M=M+1
`
	case C_POP:
		command = `@0
  D=M
  M=M-1
  `
	}
	if command != "" {
		fmt.Fprintln(cw.writer, command)
	}
}
