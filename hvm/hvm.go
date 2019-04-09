package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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

func remove(s string) string {
	comment := regexp.MustCompile(`//.*`)
	return comment.ReplaceAllString(s, "")
}

func tokenize(s string) []string {
	return strings.Split(s, " ")
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
        tokens := []string{}
	for scanner.Scan() {
		text = remove(scanner.Text())
                tokens = tokenize(text)
		fmt.Fprintln(writer, tokens)
		fmt.Println(commandType(tokens))
	}
	err = scanner.Err()
	check(err)
	writer.Flush()
}

