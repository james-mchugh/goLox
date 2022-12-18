package main

import (
	"bufio"
	"fmt"
	"github.com/james-mchugh/goLox/errorReporting"
	"github.com/james-mchugh/goLox/scanner"
	"os"
)

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

	err = run(string(fileContents))

	if errorReporting.ErrorOccurred() {
		os.Exit(65)
	}
	return err
}

func runPrompt() error {
	reader := bufio.NewReader(os.Stdin)
	for {
		_, _ = fmt.Fprintf(os.Stdout, "> ")
		line, _, _ := reader.ReadLine()
		if line == nil {
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
	s := scanner.Scanner{}

	s.Init(loxScript)

	tokens := s.ScanTokens()

	for _, token := range tokens {
		fmt.Printf("%s\n", token.ToString())
	}

	return nil
}
