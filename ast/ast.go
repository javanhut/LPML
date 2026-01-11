package ast

import (
	"lpml/tokens"
)

// Node represents any node in the AST
type Node interface {
	TokenLiteral() string
}

// Document is the root node of the AST
type Document struct {
	Sections []*PageSection
}

func (d *Document) TokenLiteral() string {
	if len(d.Sections) > 0 {
		return d.Sections[0].TokenLiteral()
	}
	return ""
}

// PageSection represents a page section (top, mid, bottom)
type PageSection struct {
	Token    tokens.Token // TOP_OF_PAGE_START, MID_PAGE_START, BOTTOM_OF_PAGE_START
	Type     string       // "top", "mid", "bottom"
	Children []Node
}

func (ps *PageSection) TokenLiteral() string {
	return ps.Token.Literal
}

// Element represents an LPML element like divide, p, h, link, etc.
type Element struct {
	Token      tokens.Token     // The opening tag token
	TagType    string           // "divide", "p", "h", "link", etc.
	Properties map[string]Value // Property assignments
	Children   []Node           // Nested elements
}

func (e *Element) TokenLiteral() string {
	return e.Token.Literal
}

// Value represents a property value (string literal, number, variable reference, or array)
type Value interface {
	Node
	valueNode()
}

// StringValue represents a string literal value like "hello"
type StringValue struct {
	Token tokens.Token
	Value string
}

func (sv *StringValue) TokenLiteral() string { return sv.Token.Literal }
func (sv *StringValue) valueNode()           {}

// NumberValue represents a numeric literal like 123 or 3.14
type NumberValue struct {
	Token tokens.Token
	Value string
}

func (nv *NumberValue) TokenLiteral() string { return nv.Token.Literal }
func (nv *NumberValue) valueNode()           {}

// VariableRef represents a variable reference like $label_name
type VariableRef struct {
	Token tokens.Token
	Name  string
}

func (vr *VariableRef) TokenLiteral() string { return vr.Token.Literal }
func (vr *VariableRef) valueNode()           {}

// ArrayValue represents an array of values like [1, 2, 3] or [$ref1, $ref2]
type ArrayValue struct {
	Token  tokens.Token
	Values []Value
}

func (av *ArrayValue) TokenLiteral() string { return av.Token.Literal }
func (av *ArrayValue) valueNode()           {}

// CodeBlockValue represents code content inside { }
type CodeBlockValue struct {
	Token   tokens.Token
	Content string
}

func (cb *CodeBlockValue) TokenLiteral() string { return cb.Token.Literal }
func (cb *CodeBlockValue) valueNode()           {}

// Property represents a property assignment like label = "value"
type Property struct {
	Name  string
	Value Value
}

// GetTagName returns the element type from a token type
func GetTagName(t tokens.TokenType) string {
	switch t {
	case tokens.DIVIDE_START, tokens.DIVIDE_END:
		return "divide"
	case tokens.P_START, tokens.P_END:
		return "p"
	case tokens.H_START, tokens.H_END:
		return "h"
	case tokens.LINK_START, tokens.LINK_END:
		return "link"
	case tokens.IMG_START, tokens.IMG_END:
		return "img"
	case tokens.LIST_START, tokens.LIST_END:
		return "list"
	case tokens.LIST_ORD_START, tokens.LIST_ORD_END:
		return "lst-ord"
	case tokens.LIST_UNORD_START, tokens.LIST_UNORD_END:
		return "lst-unord"
	case tokens.ITEM_START, tokens.ITEM_END:
		return "item"
	case tokens.TABLE_START, tokens.TABLE_END:
		return "table"
	case tokens.ROW_START, tokens.ROW_END:
		return "row"
	case tokens.CELL_START, tokens.CELL_END:
		return "cell"
	case tokens.FORM_START, tokens.FORM_END:
		return "form"
	case tokens.INPUT_START, tokens.INPUT_END:
		return "input"
	case tokens.BTN_START, tokens.BTN_END:
		return "btn"
	case tokens.BOLD_START, tokens.BOLD_END:
		return "bold"
	case tokens.ITALIC_START, tokens.ITALIC_END:
		return "italic"
	case tokens.CODE_START, tokens.CODE_END:
		return "code"
	case tokens.TOP_OF_PAGE_START, tokens.TOP_OF_PAGE_END:
		return "top-of-page"
	case tokens.MID_PAGE_START, tokens.MID_PAGE_END:
		return "mid-page"
	case tokens.BOTTOM_OF_PAGE_START, tokens.BOTTOM_OF_PAGE_END:
		return "bottom-of-page"
	}
	return ""
}

// GetSectionType returns the section type from a token type
func GetSectionType(t tokens.TokenType) string {
	switch t {
	case tokens.TOP_OF_PAGE_START:
		return "top"
	case tokens.MID_PAGE_START:
		return "mid"
	case tokens.BOTTOM_OF_PAGE_START:
		return "bottom"
	}
	return ""
}

// IsPageSection returns true if the token is a page section start
func IsPageSection(t tokens.TokenType) bool {
	switch t {
	case tokens.TOP_OF_PAGE_START, tokens.MID_PAGE_START, tokens.BOTTOM_OF_PAGE_START:
		return true
	}
	return false
}
