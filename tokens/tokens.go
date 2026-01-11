package tokens

type TokenType string

const (
	// Special tokens
	ILLEGAL TokenType = "ILLEGAL"
	EOF     TokenType = "EOF"

	// Structural tokens
	LBRACKET TokenType = "["  // [
	RBRACKET TokenType = "]"  // ]
	LBRACE   TokenType = "{"  // { for code blocks
	RBRACE   TokenType = "}"  // } for code blocks
	EQUALS   TokenType = "="  // =
	DOLLAR   TokenType = "$"  // $ for variable references
	COMMA    TokenType = ","  // , for array items
	NEWLINE  TokenType = "NEWLINE"

	// Literals
	STRING    TokenType = "STRING"    // "quoted string"
	NUMBER    TokenType = "NUMBER"    // numeric literal
	IDENT     TokenType = "IDENT"     // identifier (property names, labels)
	CODEBLOCK TokenType = "CODEBLOCK" // code content inside { }

	// Page section tags
	TOP_OF_PAGE_START    TokenType = "TOP_OF_PAGE_START"
	TOP_OF_PAGE_END      TokenType = "TOP_OF_PAGE_END"
	MID_PAGE_START       TokenType = "MID_PAGE_START"
	MID_PAGE_END         TokenType = "MID_PAGE_END"
	BOTTOM_OF_PAGE_START TokenType = "BOTTOM_OF_PAGE_START"
	BOTTOM_OF_PAGE_END   TokenType = "BOTTOM_OF_PAGE_END"

	// Element tags - opening
	DIVIDE_START   TokenType = "DIVIDE_START"
	P_START        TokenType = "P_START"
	H_START        TokenType = "H_START"
	LINK_START     TokenType = "LINK_START"
	IMG_START      TokenType = "IMG_START"
	LIST_START     TokenType = "LIST_START"
	LIST_ORD_START TokenType = "LIST_ORD_START"   // [lst-ord]
	LIST_UNORD_START TokenType = "LIST_UNORD_START" // [lst-unord]
	ITEM_START     TokenType = "ITEM_START"
	TABLE_START  TokenType = "TABLE_START"
	ROW_START    TokenType = "ROW_START"
	CELL_START   TokenType = "CELL_START"
	FORM_START   TokenType = "FORM_START"
	INPUT_START  TokenType = "INPUT_START"
	BTN_START    TokenType = "BTN_START"
	BOLD_START   TokenType = "BOLD_START"
	ITALIC_START TokenType = "ITALIC_START"
	CODE_START   TokenType = "CODE_START"

	// Element tags - closing
	DIVIDE_END     TokenType = "DIVIDE_END"
	P_END          TokenType = "P_END"
	H_END          TokenType = "H_END"
	LINK_END       TokenType = "LINK_END"
	IMG_END        TokenType = "IMG_END"
	LIST_END       TokenType = "LIST_END"
	LIST_ORD_END   TokenType = "LIST_ORD_END"   // [lst-end] for ordered
	LIST_UNORD_END TokenType = "LIST_UNORD_END" // [lst-end] for unordered
	ITEM_END       TokenType = "ITEM_END"
	TABLE_END  TokenType = "TABLE_END"
	ROW_END    TokenType = "ROW_END"
	CELL_END   TokenType = "CELL_END"
	FORM_END   TokenType = "FORM_END"
	INPUT_END  TokenType = "INPUT_END"
	BTN_END    TokenType = "BTN_END"
	BOLD_END   TokenType = "BOLD_END"
	ITALIC_END TokenType = "ITALIC_END"
	CODE_END   TokenType = "CODE_END"
)

// Token represents a lexical token
type Token struct {
	Type    TokenType
	Literal string
	Line    int
	Column  int
}

// keywords maps tag names to token types
var keywords = map[string]TokenType{
	// Page sections
	"top-of-page-start":    TOP_OF_PAGE_START,
	"top-of-page-end":      TOP_OF_PAGE_END,
	"mid-page-start":       MID_PAGE_START,
	"mid-page-end":         MID_PAGE_END,
	"bottom-of-page-start": BOTTOM_OF_PAGE_START,
	"bottom-of-page-end":   BOTTOM_OF_PAGE_END,

	// Element opening tags
	"divide-start": DIVIDE_START,
	"p-start":      P_START,
	"h-start":      H_START,
	"link-start":   LINK_START,
	"img-start":    IMG_START,
	"list-start":   LIST_START,
	"lst-ord":      LIST_ORD_START,
	"lst-unord":    LIST_UNORD_START,
	"item-start":   ITEM_START,
	"table-start":  TABLE_START,
	"row-start":    ROW_START,
	"cell-start":   CELL_START,
	"form-start":   FORM_START,
	"input-start":  INPUT_START,
	"btn-start":    BTN_START,
	"bold-start":   BOLD_START,
	"italic-start": ITALIC_START,
	"code-start":   CODE_START,

	// Element closing tags
	"divide-end": DIVIDE_END,
	"p-end":      P_END,
	"h-end":      H_END,
	"link-end":   LINK_END,
	"img-end":    IMG_END,
	"list-end":   LIST_END,
	"lst-end":    LIST_ORD_END, // shared closing tag for both list types
	"item-end":   ITEM_END,
	"table-end":  TABLE_END,
	"row-end":    ROW_END,
	"cell-end":   CELL_END,
	"form-end":   FORM_END,
	"input-end":  INPUT_END,
	"btn-end":    BTN_END,
	"bold-end":   BOLD_END,
	"italic-end": ITALIC_END,
	"code-end":   CODE_END,
}

// LookUpIdent checks if an identifier is a keyword and returns its token type
func LookUpIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}

// IsOpeningTag returns true if the token type is an opening tag
func IsOpeningTag(t TokenType) bool {
	switch t {
	case TOP_OF_PAGE_START, MID_PAGE_START, BOTTOM_OF_PAGE_START,
		DIVIDE_START, P_START, H_START, LINK_START, IMG_START,
		LIST_START, LIST_ORD_START, LIST_UNORD_START, ITEM_START,
		TABLE_START, ROW_START, CELL_START,
		FORM_START, INPUT_START, BTN_START, BOLD_START, ITALIC_START,
		CODE_START:
		return true
	}
	return false
}

// IsClosingTag returns true if the token type is a closing tag
func IsClosingTag(t TokenType) bool {
	switch t {
	case TOP_OF_PAGE_END, MID_PAGE_END, BOTTOM_OF_PAGE_END,
		DIVIDE_END, P_END, H_END, LINK_END, IMG_END,
		LIST_END, LIST_ORD_END, LIST_UNORD_END, ITEM_END,
		TABLE_END, ROW_END, CELL_END,
		FORM_END, INPUT_END, BTN_END, BOLD_END, ITALIC_END,
		CODE_END:
		return true
	}
	return false
}

// GetMatchingClose returns the closing tag type for an opening tag
func GetMatchingClose(open TokenType) TokenType {
	switch open {
	case TOP_OF_PAGE_START:
		return TOP_OF_PAGE_END
	case MID_PAGE_START:
		return MID_PAGE_END
	case BOTTOM_OF_PAGE_START:
		return BOTTOM_OF_PAGE_END
	case DIVIDE_START:
		return DIVIDE_END
	case P_START:
		return P_END
	case H_START:
		return H_END
	case LINK_START:
		return LINK_END
	case IMG_START:
		return IMG_END
	case LIST_START:
		return LIST_END
	case LIST_ORD_START:
		return LIST_ORD_END
	case LIST_UNORD_START:
		return LIST_UNORD_END
	case ITEM_START:
		return ITEM_END
	case TABLE_START:
		return TABLE_END
	case ROW_START:
		return ROW_END
	case CELL_START:
		return CELL_END
	case FORM_START:
		return FORM_END
	case INPUT_START:
		return INPUT_END
	case BTN_START:
		return BTN_END
	case BOLD_START:
		return BOLD_END
	case ITALIC_START:
		return ITALIC_END
	case CODE_START:
		return CODE_END
	}
	return ILLEGAL
}
