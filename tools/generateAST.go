package main

import (
	"bufio"
	"fmt"
	"os"
	path2 "path"
	"strings"
)

func main() {

	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Println("Usage: generateAST <output directory>")
		os.Exit(64)
	}

	outputDir := args[0]
	err := defineAST(outputDir, "Expr", []string{"Binary : left Expr, operator token.Token, right Expr", "Unary : operator token.Token, right Expr", "Grouping : expr Expr", "Literal : value interface{}"})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func defineAST(outputDir string, baseName string, types []string) error {
	path := path2.Join(outputDir, baseName) + ".go"

	outputFile, err := os.Create(path)

	if err != nil {
		return err
	}

	writer := bufio.NewWriter(outputFile)

	writer.WriteString("package syntaxTree\n")
	writer.WriteString("\n")
	writer.WriteString("import \"github.com/james-mchugh/goLox/token\"")
	writer.WriteString("\n")
	writer.WriteString("type " + baseName + " interface {\n")
	writer.WriteString("}\n")

	for _, t := range types {
		className := strings.TrimSpace(strings.Split(t, ":")[0])
		fields := strings.TrimSpace(strings.Split(t, ":")[1])
		defineType(writer, baseName, className, fields)
	}

	writer.Flush()

	return nil
}

func defineType(writer *bufio.Writer, baseName string, className string, fieldList string) {

	// struct
	writer.WriteString("type " + baseName + className + " struct {\n")

	fields := strings.Split(fieldList, ", ")
	for _, field := range fields {
		writer.WriteString("    " + field + "\n")
	}
	writer.WriteString("}\n")

	// init receiver
	writer.WriteString("func (this *" + baseName + className + ") Init(" + fieldList + ") {\n")
	for _, field := range fields {
		fieldName := strings.Split(field, " ")[0]
		writer.WriteString("    this." + fieldName + " = " + fieldName + "\n")
	}
	writer.WriteString("}\n")
}
