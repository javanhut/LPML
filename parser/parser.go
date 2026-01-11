package parser

import (
	"fmt"
	"lpml/ast"
	"lpml/lexer"
	"lpml/tokens"
)

// Parser parses LPML tokens into an AST
type Parser struct {
	l         *lexer.Lexer
	curToken  tokens.Token
	peekToken tokens.Token
	errors    []string
}

// New creates a new Parser
func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l, errors: []string{}}
	// Read two tokens to initialize curToken and peekToken
	p.nextToken()
	p.nextToken()
	return p
}

// Errors returns any parsing errors
func (p *Parser) Errors() []string {
	return p.errors
}

// nextToken advances to the next token
func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

// ParseDocument parses the entire document
func (p *Parser) ParseDocument() *ast.Document {
	doc := &ast.Document{Sections: []*ast.PageSection{}}

	for p.curToken.Type != tokens.EOF {
		if ast.IsPageSection(p.curToken.Type) {
			section := p.parsePageSection()
			if section != nil {
				doc.Sections = append(doc.Sections, section)
			}
		} else {
			p.nextToken()
		}
	}

	return doc
}

// parsePageSection parses a page section (top, mid, bottom)
func (p *Parser) parsePageSection() *ast.PageSection {
	section := &ast.PageSection{
		Token:    p.curToken,
		Type:     ast.GetSectionType(p.curToken.Type),
		Children: []ast.Node{},
	}

	closingTag := tokens.GetMatchingClose(p.curToken.Type)
	p.nextToken() // move past opening tag

	// Parse children until we hit the closing tag
	for p.curToken.Type != closingTag && p.curToken.Type != tokens.EOF {
		child := p.parseElement()
		if child != nil {
			section.Children = append(section.Children, child)
		}
	}

	if p.curToken.Type == closingTag {
		p.nextToken() // consume closing tag
	} else {
		p.addError(fmt.Sprintf("expected closing tag for section %s", section.Type))
	}

	return section
}

// parseElement parses an element like [p-start]...[p-end]
func (p *Parser) parseElement() *ast.Element {
	if !tokens.IsOpeningTag(p.curToken.Type) {
		p.nextToken()
		return nil
	}

	elem := &ast.Element{
		Token:      p.curToken,
		TagType:    ast.GetTagName(p.curToken.Type),
		Properties: make(map[string]ast.Value),
		Children:   []ast.Node{},
	}

	openingType := p.curToken.Type
	p.nextToken() // move past opening tag

	// Parse properties and children until we hit the closing tag
	for !p.isMatchingClose(openingType, p.curToken.Type) && p.curToken.Type != tokens.EOF {
		if p.curToken.Type == tokens.IDENT {
			// This is a property assignment
			p.parseProperty(elem)
		} else if tokens.IsOpeningTag(p.curToken.Type) {
			// This is a nested element
			child := p.parseElement()
			if child != nil {
				elem.Children = append(elem.Children, child)
			}
		} else {
			p.nextToken()
		}
	}

	if p.isMatchingClose(openingType, p.curToken.Type) {
		p.nextToken() // consume closing tag
	} else {
		p.addError(fmt.Sprintf("expected closing tag for element %s at line %d", elem.TagType, elem.Token.Line))
	}

	return elem
}

// isMatchingClose checks if the current token is a valid closing tag for the opening tag
func (p *Parser) isMatchingClose(open, close tokens.TokenType) bool {
	// Special case: lst-end closes both lst-ord and lst-unord
	if (open == tokens.LIST_ORD_START || open == tokens.LIST_UNORD_START) &&
		(close == tokens.LIST_ORD_END || close == tokens.LIST_UNORD_END) {
		return true
	}
	return close == tokens.GetMatchingClose(open)
}

// parseProperty parses a property assignment like label = "value" or linked = $ref or items = [1,2,3]
func (p *Parser) parseProperty(elem *ast.Element) {
	propName := p.curToken.Literal
	p.nextToken() // move past property name

	// Expect '='
	if p.curToken.Type != tokens.EQUALS {
		p.addError(fmt.Sprintf("expected '=' after property name %s, got %s", propName, p.curToken.Type))
		return
	}
	p.nextToken() // consume '='

	// Parse value (string, number, variable reference, or array)
	value := p.parseValue(propName)
	if value != nil {
		elem.Properties[propName] = value
	}
}

// parseValue parses a value (string, number, variable reference, array, or code block)
func (p *Parser) parseValue(propName string) ast.Value {
	switch p.curToken.Type {
	case tokens.STRING:
		value := &ast.StringValue{
			Token: p.curToken,
			Value: p.curToken.Literal,
		}
		p.nextToken()
		return value

	case tokens.NUMBER:
		value := &ast.NumberValue{
			Token: p.curToken,
			Value: p.curToken.Literal,
		}
		p.nextToken()
		return value

	case tokens.DOLLAR:
		value := &ast.VariableRef{
			Token: p.curToken,
			Name:  p.curToken.Literal,
		}
		p.nextToken()
		return value

	case tokens.LBRACKET:
		return p.parseArray()

	case tokens.CODEBLOCK:
		value := &ast.CodeBlockValue{
			Token:   p.curToken,
			Content: p.curToken.Literal,
		}
		p.nextToken()
		return value

	default:
		p.addError(fmt.Sprintf("expected value for property %s, got %s", propName, p.curToken.Type))
		return nil
	}
}

// parseArray parses an array like [1, 2, 3] or [$ref1, $ref2] or ["a", "b"]
func (p *Parser) parseArray() *ast.ArrayValue {
	arr := &ast.ArrayValue{
		Token:  p.curToken,
		Values: []ast.Value{},
	}

	p.nextToken() // consume '['

	// Parse array elements
	for p.curToken.Type != tokens.RBRACKET && p.curToken.Type != tokens.EOF {
		var val ast.Value

		switch p.curToken.Type {
		case tokens.STRING:
			val = &ast.StringValue{Token: p.curToken, Value: p.curToken.Literal}
			p.nextToken()
		case tokens.NUMBER:
			val = &ast.NumberValue{Token: p.curToken, Value: p.curToken.Literal}
			p.nextToken()
		case tokens.DOLLAR:
			val = &ast.VariableRef{Token: p.curToken, Name: p.curToken.Literal}
			p.nextToken()
		case tokens.COMMA:
			p.nextToken() // skip comma
			continue
		default:
			p.nextToken() // skip unknown
			continue
		}

		if val != nil {
			arr.Values = append(arr.Values, val)
		}
	}

	if p.curToken.Type == tokens.RBRACKET {
		p.nextToken() // consume ']'
	}

	return arr
}

// addError adds a parsing error
func (p *Parser) addError(msg string) {
	p.errors = append(p.errors, msg)
}
