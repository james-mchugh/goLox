package main

import "fmt"

type tokenType = string

//goland:noinspection GoUnusedConst
const (
	// Single-character tokens
	LeftParen  tokenType = "LeftParen"
	RightParen tokenType = "RightParen"
	LeftBrace  tokenType = "LeftBrace"
	RightBrace tokenType = "RightBrace"
	Comma      tokenType = "Comma"
	Dot        tokenType = "Dot"
	Minus      tokenType = "Minus"
	Plus       tokenType = "Plus"
	Semicolon  tokenType = "Semicolon"
	Slash      tokenType = "Slash"
	Star       tokenType = "Star"

	// One or two character tokens
	Bang         tokenType = "Bang"
	BangEqual    tokenType = "BangEqual"
	Equal        tokenType = "Equal"
	EqualEqual   tokenType = "EqualEqual"
	Greater      tokenType = "Greater"
	GreaterEqual tokenType = "GreaterEqual"
	Less         tokenType = "Less"
	LessEqual    tokenType = "LessEqual"

	// Literals
	Identifier tokenType = "Identifier"
	String     tokenType = "String"
	Number     tokenType = "Number"

	// Keywords
	And    tokenType = "And"
	Class  tokenType = "Class"
	Else   tokenType = "Else"
	False  tokenType = "False"
	Fun    tokenType = "Fun"
	For    tokenType = "For"
	If     tokenType = "If"
	Nil    tokenType = "Nil"
	Or     tokenType = "Or"
	Print  tokenType = "Print"
	Return tokenType = "Return"
	Super  tokenType = "Super"
	This   tokenType = "This"
	True   tokenType = "True"
	Var    tokenType = "Var"
	While  tokenType = "While"

	EOF tokenType = "EOF"
)

type Token struct {
	type_   tokenType
	lexeme  string
	literal interface{}
	line    int
}

func Init(type_ tokenType, lexeme string, literal interface{}, line int) *Token {
	return &Token{type_, lexeme, literal, line}
}

func (token Token) ToString() string {
	return fmt.Sprintf("%s %s %v", token.type_, token.lexeme, token.literal)
}
