package token

import "fmt"

type TokenType = string

//goland:noinspection GoUnusedConst
const (
	// Single-character tokens
	LeftParen  TokenType = "LeftParen"
	RightParen TokenType = "RightParen"
	LeftBrace  TokenType = "LeftBrace"
	RightBrace TokenType = "RightBrace"
	Comma      TokenType = "Comma"
	Dot        TokenType = "Dot"
	Minus      TokenType = "Minus"
	Plus       TokenType = "Plus"
	Semicolon  TokenType = "Semicolon"
	Slash      TokenType = "Slash"
	Star       TokenType = "Star"

	// One or two character tokens
	Bang         TokenType = "Bang"
	BangEqual    TokenType = "BangEqual"
	Equal        TokenType = "Equal"
	EqualEqual   TokenType = "EqualEqual"
	Greater      TokenType = "Greater"
	GreaterEqual TokenType = "GreaterEqual"
	Less         TokenType = "Less"
	LessEqual    TokenType = "LessEqual"

	// Literals
	Identifier TokenType = "Identifier"
	String     TokenType = "String"
	Number     TokenType = "Number"

	// Keywords
	And    TokenType = "And"
	Class  TokenType = "Class"
	Else   TokenType = "Else"
	False  TokenType = "False"
	Fun    TokenType = "Fun"
	For    TokenType = "For"
	If     TokenType = "If"
	Nil    TokenType = "Nil"
	Or     TokenType = "Or"
	Print  TokenType = "Print"
	Return TokenType = "Return"
	Super  TokenType = "Super"
	This   TokenType = "This"
	True   TokenType = "True"
	Var    TokenType = "Var"
	While  TokenType = "While"

	EOF TokenType = "EOF"
)

type Token struct {
	type_   TokenType
	lexeme  string
	literal interface{}
	line    int
}

func NewToken(type_ TokenType, lexeme string, literal interface{}, line int) *Token {
	return &Token{type_, lexeme, literal, line}
}

func (token Token) ToString() string {
	return fmt.Sprintf("%s %s %v", token.type_, token.lexeme, token.literal)
}
