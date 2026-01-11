package lexer

import (
	"lpml/tokens"
)

// Lexer tokenizes LPML input
type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
	line         int  // current line number
	column       int  // current column number
}

// New creates a new Lexer for the given input
func New(input string) *Lexer {
	l := &Lexer{input: input, line: 1, column: 0}
	l.readChar()
	return l
}

// readChar reads the next character and advances positions
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 // ASCII code for NUL
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
	l.column++

	if l.ch == '\n' {
		l.line++
		l.column = 0
	}
}

// peekChar returns the next character without advancing
func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}

// NextToken returns the next token from the input
func (l *Lexer) NextToken() tokens.Token {
	var tok tokens.Token

	l.skipWhitespace()

	tok.Line = l.line
	tok.Column = l.column

	switch l.ch {
	case '[':
		// Check if next char suggests this is an array or a tag
		if l.isArrayStart() {
			tok = newToken(tokens.LBRACKET, l.ch, l.line, l.column)
			l.readChar()
		} else {
			tok = l.readTag()
		}
	case ']':
		tok = newToken(tokens.RBRACKET, l.ch, l.line, l.column)
		l.readChar()
	case '=':
		tok = newToken(tokens.EQUALS, l.ch, l.line, l.column)
		l.readChar()
	case ',':
		tok = newToken(tokens.COMMA, l.ch, l.line, l.column)
		l.readChar()
	case '{':
		tok = l.readCodeBlock()
	case '$':
		tok = l.readVariableReference()
	case '"':
		tok.Type = tokens.STRING
		tok.Literal = l.readString()
		tok.Line = l.line
		tok.Column = l.column
	case '\n':
		tok = newToken(tokens.NEWLINE, l.ch, l.line, l.column)
		l.readChar()
	case 0:
		tok.Type = tokens.EOF
		tok.Literal = ""
		tok.Line = l.line
		tok.Column = l.column
	default:
		if isDigit(l.ch) {
			tok.Line = l.line
			tok.Column = l.column
			tok.Literal = l.readNumber()
			tok.Type = tokens.NUMBER
			return tok
		} else if isLetter(l.ch) || l.ch == '_' {
			tok.Line = l.line
			tok.Column = l.column
			tok.Literal = l.readIdentifier()
			tok.Type = tokens.IDENT
			return tok
		} else {
			tok = newToken(tokens.ILLEGAL, l.ch, l.line, l.column)
			l.readChar()
		}
	}

	return tok
}

// isArrayStart checks if '[' is start of array (not a tag)
// Arrays start with [ followed by number, $, ", or ]
func (l *Lexer) isArrayStart() bool {
	next := l.peekChar()
	return isDigit(next) || next == '$' || next == '"' || next == ']' || next == ' ' || next == '\n' || next == '\t'
}

// readNumber reads a numeric literal
func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) || l.ch == '.' {
		l.readChar()
	}
	return l.input[position:l.position]
}

// readCodeBlock reads content between { and } for code blocks
func (l *Lexer) readCodeBlock() tokens.Token {
	line := l.line
	col := l.column

	l.readChar() // consume '{'

	// Skip initial whitespace/newlines after {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}

	position := l.position
	braceCount := 1

	// Read until matching closing brace
	for braceCount > 0 && l.ch != 0 {
		if l.ch == '{' {
			braceCount++
		} else if l.ch == '}' {
			braceCount--
			if braceCount == 0 {
				break
			}
		}
		l.readChar()
	}

	content := l.input[position:l.position]

	// Trim trailing whitespace from content
	content = trimTrailingWhitespace(content)

	if l.ch == '}' {
		l.readChar() // consume '}'
	}

	return tokens.Token{
		Type:    tokens.CODEBLOCK,
		Literal: content,
		Line:    line,
		Column:  col,
	}
}

// trimTrailingWhitespace removes trailing whitespace from a string
func trimTrailingWhitespace(s string) string {
	end := len(s)
	for end > 0 && (s[end-1] == ' ' || s[end-1] == '\t' || s[end-1] == '\n' || s[end-1] == '\r') {
		end--
	}
	return s[:end]
}

// readTag reads a bracketed tag like [tag-name] or [tag-name]
func (l *Lexer) readTag() tokens.Token {
	line := l.line
	col := l.column

	l.readChar() // consume '['

	// Read the tag name
	tagName := l.readTagName()

	// Skip to the closing bracket
	for l.ch != ']' && l.ch != 0 {
		l.readChar()
	}

	if l.ch == ']' {
		l.readChar() // consume ']'
	}

	// Look up if this is a known tag
	tokType := tokens.LookUpIdent(tagName)

	return tokens.Token{
		Type:    tokType,
		Literal: tagName,
		Line:    line,
		Column:  col,
	}
}

// readTagName reads the name inside brackets
func (l *Lexer) readTagName() string {
	position := l.position
	for isTagChar(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// readIdentifier reads an identifier (property name, label value, etc.)
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) || isDigit(l.ch) || l.ch == '_' {
		l.readChar()
	}
	return l.input[position:l.position]
}

// readString reads a quoted string
func (l *Lexer) readString() string {
	l.readChar() // consume opening quote
	position := l.position
	for l.ch != '"' && l.ch != 0 {
		l.readChar()
	}
	str := l.input[position:l.position]
	if l.ch == '"' {
		l.readChar() // consume closing quote
	}
	return str
}

// readVariableReference reads a variable reference like $label_name
func (l *Lexer) readVariableReference() tokens.Token {
	line := l.line
	col := l.column

	l.readChar() // consume '$'

	varName := l.readIdentifier()

	return tokens.Token{
		Type:    tokens.DOLLAR,
		Literal: varName,
		Line:    line,
		Column:  col,
	}
}

// skipWhitespace skips spaces and tabs (but not newlines)
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\r' || l.ch == '\n' {
		l.readChar()
	}
}

// newToken creates a new token
func newToken(tokenType tokens.TokenType, ch byte, line, col int) tokens.Token {
	return tokens.Token{
		Type:    tokenType,
		Literal: string(ch),
		Line:    line,
		Column:  col,
	}
}

// isLetter checks if character is a letter
func isLetter(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}

// isDigit checks if character is a digit
func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

// isTagChar checks if character can be part of a tag name (letters, digits, hyphen)
func isTagChar(ch byte) bool {
	return isLetter(ch) || isDigit(ch) || ch == '-' || ch == '_'
}
