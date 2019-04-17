package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

func (cw *CodeWriter) writeArithmetic(arithmetic string) {
	if isArithmetic(arithmetic) {
		panic(arithmetic + " is not arithmetic opcode")
	}

	// Stack pointer should be decremented
	// because current SP indicates next empty place.
	asm := "@SP\n"
	asm += "M=M-1\n" // sp--
	// Get the topmost data in the stack and
	// store it into D register.
	asm += "A=M\n" // A = sp
	asm += "D=M\n" // D = RAM[sp]
	switch arithmetic {
	case "add":
		asm += "@SP\n"
		asm += "M=M-1\n" // sp--
		asm += "A=M\n"   // A = sp
		asm += "M=M+D\n" // RAM[sp] = RAM[sp] + D
	case "sub":
		asm += "M=M-D"
	case "neg":
		asm += "M=-M"
	case "eq":
		asm += "JEQ"
	case "gt":
		asm += "JGT"
	case "lt":
		asm += "JLT"
	case "and":
		asm += "M=D&M"
	case "or":
		asm += "M=D|A"
	case "not":
		asm += "M=!M"
	}
	fmt.Fprintln(cw.writer, asm)
}

func (cw *CodeWriter) writePushPop(ct COMMAND_TYPE, s *Stack, seg string, idx int) {
	// Stack pointer(SP) is hold at RAM[0].
	// Stack base starts from RAM[256].
	asm := ""
	val := s.s[seg][idx]

	switch ct {
	case C_PUSH:
		asm += "@" + strconv.Itoa(val) + "\n"
		asm += "D=A\n" // D = stack[idx]
		asm += "@SP\n"
		asm += "A=M\n"
		asm += "M=D\n" // RAM[sp] = D
		asm += "@SP\n"
		asm += "M=M+1\n" // RAM[sp]++
	case C_POP:
		// TODO: Where should I store D?
		asm += "@SP\n"
		asm += "A=M\n"
		asm += "D=M\n" // D = RAM[sp]
	}
	if asm != "" {
		fmt.Fprintln(cw.writer, asm)
	}
}
