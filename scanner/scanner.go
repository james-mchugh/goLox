package scanner

import (
	"fmt"
	"github.com/james-mchugh/goLox/errorReporting"
	"github.com/james-mchugh/goLox/token"
	"strings"
)

type Scanner struct {
	source  string
	tokens  []token.Token
	reader  *strings.Reader
	start   int
	current int
	line    int
}

func (scanner *Scanner) Init(source string) {
	scanner.source = source
	scanner.line = 1
	scanner.reader = strings.NewReader(source)
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
	currentChar, err := scanner.advance()
	if err != nil {
		return err
	}
	switch currentChar {
	case '(':
		scanner.addToken(token.LeftParen)
	case ')':
		scanner.addToken(token.RightParen)
	case '{':
		scanner.addToken(token.LeftBrace)
	case '}':
		scanner.addToken(token.RightBrace)
	case ',':
		scanner.addToken(token.Comma)
	case '.':
		scanner.addToken(token.Dot)
	case '-':
		scanner.addToken(token.Minus)
	case '+':
		scanner.addToken(token.Plus)
	case ';':
		scanner.addToken(token.Semicolon)
	case '*':
		scanner.addToken(token.Star)
	default:
		errorReporting.Error(scanner.line, fmt.Sprintf("Unexpected character: %c", currentChar))
	}

	return nil
}

func (scanner *Scanner) advance() (byte, error) {
	scanner.current++
	char, err := scanner.reader.ReadByte()
	if err != nil {
		return 0, err
	}
	return char, nil
}

func (scanner *Scanner) addToken(tokenType token.TokenType) {
	text := scanner.source[scanner.start:scanner.current]
	scanner.tokens = append(scanner.tokens, *token.NewToken(tokenType, text, nil, scanner.line))
}
