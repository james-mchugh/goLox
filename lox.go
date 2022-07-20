package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"text/scanner"
)

//type logLevel = string

func main() {

	args := os.Args[1:]

	if 1 < len(args) {
		fmt.Println("Usage: jlox [script]")
		os.Exit(64)
	}
	var err error = nil
	if len(args) == 1 {
		err = runFile(args[0])

	} else {
		err = runPrompt()
	}

	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "jlox: %v\n", err)
		os.Exit(1)
	}

}

func runFile(filePath string) error {
	fileContents, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}
	return run(string(fileContents))
}

func runPrompt() error {
	reader := bufio.NewReader(os.Stdin)
	for {
		_, _ = fmt.Fprintf(os.Stdout, "> ")
		line, _, _ := reader.ReadLine()
		if string(line) == "" {
			break
		}
		err := run(string(line))
		if err != nil {
			return err
		}
	}

	return nil

}

func run(loxScript string) error {
	reader := strings.NewReader(loxScript)
	s := scanner.Scanner{}

	s.Init(reader)

	for currentToken := s.Scan(); currentToken != scanner.EOF; currentToken = s.Scan() {
		fmt.Printf("%s: %s\n", s.Pos(), s.TokenText())
	}

	return nil
}

//func reportError(line int, message string) {
//	report(line, "", message, "Error")
//}
//
//func report(line int, s string, message string, level logLevel) {
//	_, _ = fmt.Fprintf(os.Stderr, "[line %d] %s %s: %s\n", line, s, message, level)
//}
