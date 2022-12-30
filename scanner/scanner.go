package scanner

import (
	"fmt"
	"github.com/james-mchugh/goLox/errorReporting"
	"github.com/james-mchugh/goLox/token"
	"strconv"
)

type Scanner struct {
	source  string
	tokens  []token.Token
	start   int
	current int
	line    int
}

func (scanner *Scanner) Init(source string) {
	scanner.source = source
	scanner.line = 1
}

func (scanner *Scanner) ScanTokens() []token.Token {
	for !scanner.finishedScanning() {
		scanner.start = scanner.current
		err := scanner.ScanToken()
		if err != nil {
			return nil
		}
	}

	scanner.tokens = append(scanner.tokens, *token.NewToken(token.EOF, "", nil, scanner.line))
	return scanner.tokens

}

func (scanner *Scanner) finishedScanning() bool {
	return len(scanner.source) <= scanner.current
}

func (scanner *Scanner) ScanToken() error {
	var tokenType token.TokenType
	switch currentChar := scanner.advance(); currentChar {
	case '(':
		scanner.addToken(token.LeftParen, nil)
	case ')':
		scanner.addToken(token.RightParen, nil)
	case '{':
		scanner.addToken(token.LeftBrace, nil)
	case '}':
		scanner.addToken(token.RightBrace, nil)
	case ',':
		scanner.addToken(token.Comma, nil)
	case '.':
		scanner.addToken(token.Dot, nil)
	case '-':
		scanner.addToken(token.Minus, nil)
	case '+':
		scanner.addToken(token.Plus, nil)
	case ';':
		scanner.addToken(token.Semicolon, nil)
	case '*':
		scanner.addToken(token.Star, nil)
	case '!':
		if scanner.match('=') {
			tokenType = token.BangEqual
		} else {
			tokenType = token.Bang
		}
		scanner.addToken(tokenType, nil)
	case '=':
		if scanner.match('=') {
			tokenType = token.EqualEqual
		} else {
			tokenType = token.Equal
		}
		scanner.addToken(tokenType, nil)
	case '<':
		if scanner.match('=') {
			tokenType = token.LessEqual
		} else {
			tokenType = token.Less
		}
		scanner.addToken(tokenType, nil)
	case '>':
		if scanner.match('=') {
			tokenType = token.GreaterEqual
		} else {
			tokenType = token.Greater
		}
		scanner.addToken(tokenType, nil)
	case '/':
		if scanner.match('/') {
			// If a comment symbol is found ('//'), consume the line until the end
			for currentChar := scanner.peek(); currentChar != '\n' && !scanner.finishedScanning(); currentChar = scanner.peek() {
				scanner.advance()
			}

		} else {
			tokenType = token.Slash
		}
		scanner.addToken(tokenType, nil)
	case '"':
		scanner.consumeString()
	case '\n':
		scanner.line++
	case '\r':
	case '\t':
	case ' ':
	default:
		if scanner.isDigit(currentChar) {
			scanner.consumeNumber()
		} else if scanner.isAlpha(currentChar) {
			scanner.consumeIdentifier()
		} else {
			errorReporting.Error(scanner.line, fmt.Sprintf("Unexpected character: %c", currentChar))

		}
	}

	return nil
}

func (scanner *Scanner) match(char byte) bool {
	if scanner.finishedScanning() {
		return false
	}
	if scanner.source[scanner.current] != char {
		return false
	}
	scanner.current++
	return true
}

func (scanner *Scanner) advance() byte {
	char := scanner.source[scanner.current]
	scanner.current++
	return char
}

func (scanner *Scanner) peek() byte {
	if scanner.finishedScanning() {
		return '\000'
	}
	char := scanner.source[scanner.current]
	return char
}

func (scanner *Scanner) peekNext() byte {
	if scanner.current+1 >= len(scanner.source) {
		return '\000'
	}
	char := scanner.source[scanner.current+1]
	return char
}

func (scanner *Scanner) consumeString() {

	// Consume everything up to the closing '"'
	for currentChar := scanner.peek(); currentChar != '"' && !scanner.finishedScanning(); currentChar = scanner.peek() {
		if currentChar == '\n' {
			scanner.line++
		}
		scanner.advance()
	}

	// If finished scanning without encountering the ending quote, report an error
	if scanner.finishedScanning() {
		errorReporting.Error(scanner.line, "Unterminated string.")
		return
	}

	// Consume the closing '"'
	scanner.advance()

	// Drop the quotes from the lexeme
	tokenLiteral := scanner.source[scanner.start+1 : scanner.current-1]
	scanner.addToken(token.String, tokenLiteral)

}

func (scanner *Scanner) consumeNumber() {
	for scanner.isDigit(scanner.peek()) {
		scanner.advance()
	}

	if scanner.peek() == '.' && scanner.isDigit(scanner.peekNext()) {
		scanner.advance()
		for scanner.isDigit(scanner.peek()) {
			scanner.advance()
		}
	}

	value, err := strconv.ParseFloat(scanner.source[scanner.start:scanner.current], 32)

	if err != nil {
		errorReporting.Error(scanner.line, "Attempted to parse invalid number.")
		return
	}

	scanner.addToken(token.Number, value)
}

func (scanner *Scanner) consumeIdentifier() {
	for scanner.isAlphaNumeric(scanner.peek()) {
		scanner.advance()
	}

	var tokenType token.TokenType
	switch identifier := scanner.source[scanner.start:scanner.current]; identifier {
	case "and":
		tokenType = token.And
	case "or":
		tokenType = token.Or
	case "class":
		tokenType = token.Class
	case "if":
		tokenType = token.If
	case "else":
		tokenType = token.Else
	case "while":
		tokenType = token.While
	case "for":
		tokenType = token.For
	case "fun":
		tokenType = token.Fun
	case "true":
		tokenType = token.True
	case "false":
		tokenType = token.False
	case "super":
		tokenType = token.Super
	case "print":
		tokenType = token.Print
	case "var":
		tokenType = token.Var
	case "this":
		tokenType = token.This
	case "nil":
		tokenType = token.Nil
	case "return":
		tokenType = token.Return
	default:
		tokenType = token.Identifier

	}

	scanner.addToken(tokenType, nil)

}

func (scanner *Scanner) addToken(tokenType token.TokenType, value interface{}) {
	text := scanner.source[scanner.start:scanner.current]
	scanner.tokens = append(scanner.tokens, *token.NewToken(tokenType, text, value, scanner.line))
}

func (scanner *Scanner) isDigit(char byte) bool {
	return '0' <= char && char <= '9'
}

func (scanner *Scanner) isAlpha(char byte) bool {
	return ('a' <= char && char <= 'z') || ('A' <= char && char <= 'Z') || (char == '_')
}

func (scanner *Scanner) isAlphaNumeric(char byte) bool {
	return scanner.isAlpha(char) || scanner.isDigit(char)
}
