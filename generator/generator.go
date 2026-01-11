package generator

import (
	"fmt"
	"lpml/ast"
	"strings"
)

// Generator converts AST to HTML
type Generator struct {
	labels map[string]*ast.Element // Store labeled elements for variable resolution
	indent int
}

// New creates a new Generator
func New() *Generator {
	return &Generator{
		labels: make(map[string]*ast.Element),
		indent: 0,
	}
}

// Generate produces HTML from the AST
func (g *Generator) Generate(doc *ast.Document) string {
	var sb strings.Builder

	// First pass: collect all labeled elements
	g.collectLabels(doc)

	// Write HTML document structure
	sb.WriteString("<!DOCTYPE html>\n")
	sb.WriteString("<html>\n")
	sb.WriteString("<head>\n")
	sb.WriteString("  <title>LPML Document</title>\n")
	sb.WriteString("  <style>\n")
	sb.WriteString("    .top-of-page { }\n")
	sb.WriteString("    .mid-page { }\n")
	sb.WriteString("    .bottom-of-page { }\n")
	sb.WriteString("  </style>\n")
	sb.WriteString("</head>\n")
	sb.WriteString("<body>\n")

	// Generate each section
	for _, section := range doc.Sections {
		sb.WriteString(g.generateSection(section))
	}

	sb.WriteString("</body>\n")
	sb.WriteString("</html>\n")

	return sb.String()
}

// collectLabels finds all elements with labels for variable resolution
func (g *Generator) collectLabels(doc *ast.Document) {
	for _, section := range doc.Sections {
		for _, child := range section.Children {
			g.collectElementLabels(child)
		}
	}
}

// collectElementLabels recursively collects labels from elements
func (g *Generator) collectElementLabels(node ast.Node) {
	elem, ok := node.(*ast.Element)
	if !ok {
		return
	}

	if labelVal, exists := elem.Properties["label"]; exists {
		if sv, ok := labelVal.(*ast.StringValue); ok {
			g.labels[sv.Value] = elem
		}
	}

	for _, child := range elem.Children {
		g.collectElementLabels(child)
	}
}

// generateSection generates HTML for a page section
func (g *Generator) generateSection(section *ast.PageSection) string {
	var sb strings.Builder
	className := section.Type + "-of-page"
	if section.Type == "mid" {
		className = "mid-page"
	}

	sb.WriteString(fmt.Sprintf("  <div class=\"%s\">\n", className))

	g.indent = 2
	for _, child := range section.Children {
		sb.WriteString(g.generateNode(child))
	}

	sb.WriteString("  </div>\n")

	return sb.String()
}

// generateNode generates HTML for any AST node
func (g *Generator) generateNode(node ast.Node) string {
	elem, ok := node.(*ast.Element)
	if !ok {
		return ""
	}

	return g.generateElement(elem)
}

// generateElement generates HTML for an element
func (g *Generator) generateElement(elem *ast.Element) string {
	var sb strings.Builder
	indent := strings.Repeat("  ", g.indent)

	switch elem.TagType {
	case "divide":
		sb.WriteString(g.generateDiv(elem, indent))
	case "p":
		sb.WriteString(g.generateParagraph(elem, indent))
	case "h":
		sb.WriteString(g.generateHeading(elem, indent))
	case "link":
		sb.WriteString(g.generateLink(elem, indent))
	case "img":
		sb.WriteString(g.generateImage(elem, indent))
	case "list":
		sb.WriteString(g.generateList(elem, indent, false))
	case "olist":
		sb.WriteString(g.generateList(elem, indent, true))
	case "lst-ord":
		sb.WriteString(g.generateList(elem, indent, true))
	case "lst-unord":
		sb.WriteString(g.generateList(elem, indent, false))
	case "item":
		sb.WriteString(g.generateListItem(elem, indent))
	case "table":
		sb.WriteString(g.generateTable(elem, indent))
	case "row":
		sb.WriteString(g.generateRow(elem, indent))
	case "cell":
		sb.WriteString(g.generateCell(elem, indent))
	case "form":
		sb.WriteString(g.generateForm(elem, indent))
	case "input":
		sb.WriteString(g.generateInput(elem, indent))
	case "btn":
		sb.WriteString(g.generateButton(elem, indent))
	case "bold":
		sb.WriteString(g.generateBold(elem, indent))
	case "italic":
		sb.WriteString(g.generateItalic(elem, indent))
	case "code":
		sb.WriteString(g.generateCode(elem, indent))
	}

	return sb.String()
}

// generateDiv generates a <div> element
func (g *Generator) generateDiv(elem *ast.Element, indent string) string {
	var sb strings.Builder

	id := g.getStringProp(elem, "label")
	class := g.getStringProp(elem, "class")
	styleAttr := g.buildStyleAttr(elem)

	sb.WriteString(indent + "<div")
	if id != "" {
		sb.WriteString(fmt.Sprintf(" id=\"%s\"", id))
	}
	if class != "" {
		sb.WriteString(fmt.Sprintf(" class=\"%s\"", class))
	}
	sb.WriteString(styleAttr)
	sb.WriteString(">\n")

	g.indent++
	for _, child := range elem.Children {
		sb.WriteString(g.generateNode(child))
	}
	g.indent--

	sb.WriteString(indent + "</div>\n")
	return sb.String()
}

// generateParagraph generates a <p> element
func (g *Generator) generateParagraph(elem *ast.Element, indent string) string {
	content := g.getStringProp(elem, "contains")
	id := g.getStringProp(elem, "label")

	// Apply formatting from format_with property
	content = g.applyFormatting(elem, content)

	idAttr := ""
	if id != "" {
		idAttr = fmt.Sprintf(" id=\"%s\"", id)
	}

	styleAttr := g.buildStyleAttr(elem)

	return fmt.Sprintf("%s<p%s%s>%s</p>\n", indent, idAttr, styleAttr, content)
}

// applyFormatting wraps content with formatting tags based on format_with property
func (g *Generator) applyFormatting(elem *ast.Element, content string) string {
	if formatVal, exists := elem.Properties["format_with"]; exists {
		if arr, ok := formatVal.(*ast.ArrayValue); ok {
			for _, item := range arr.Values {
				format := g.resolveValue(item)
				switch format {
				case "bold":
					content = "<strong>" + content + "</strong>"
				case "italic":
					content = "<em>" + content + "</em>"
				case "underline":
					content = "<u>" + content + "</u>"
				case "strike":
					content = "<s>" + content + "</s>"
				case "code":
					content = "<code>" + content + "</code>"
				case "mark":
					content = "<mark>" + content + "</mark>"
				}
			}
		}
	}
	return content
}

// buildStyleAttr builds inline CSS from friendly property names
func (g *Generator) buildStyleAttr(elem *ast.Element) string {
	var styles []string

	// Text color
	if v := g.getStringProp(elem, "text_color"); v != "" {
		styles = append(styles, fmt.Sprintf("color: %s", v))
	}
	if v := g.getStringProp(elem, "color"); v != "" {
		styles = append(styles, fmt.Sprintf("color: %s", v))
	}

	// Background
	if v := g.getStringProp(elem, "bg_color"); v != "" {
		styles = append(styles, fmt.Sprintf("background-color: %s", v))
	}
	if v := g.getStringProp(elem, "background"); v != "" {
		// Use 'background' for gradients, 'background-color' for solid colors
		if strings.Contains(v, "gradient") || strings.Contains(v, "url(") {
			styles = append(styles, fmt.Sprintf("background: %s", v))
		} else {
			styles = append(styles, fmt.Sprintf("background-color: %s", v))
		}
	}

	// Font size - support friendly names
	if v := g.getStringProp(elem, "text_size"); v != "" {
		styles = append(styles, fmt.Sprintf("font-size: %s", g.resolveFontSize(v)))
	}

	// Font family
	if v := g.getStringProp(elem, "font"); v != "" {
		styles = append(styles, fmt.Sprintf("font-family: %s", v))
	}

	// Text alignment
	if v := g.getStringProp(elem, "align"); v != "" {
		styles = append(styles, fmt.Sprintf("text-align: %s", v))
	}

	// Padding - support friendly names
	if v := g.getStringProp(elem, "padding"); v != "" {
		styles = append(styles, fmt.Sprintf("padding: %s", g.resolveSpacing(v)))
	}

	// Margin
	if v := g.getStringProp(elem, "margin"); v != "" {
		styles = append(styles, fmt.Sprintf("margin: %s", g.resolveSpacing(v)))
	}

	// Border - friendly syntax
	if v := g.getStringProp(elem, "border"); v != "" {
		styles = append(styles, fmt.Sprintf("border: %s", g.resolveBorder(v)))
	}

	// Border radius (rounded corners)
	if v := g.getStringProp(elem, "rounded"); v != "" {
		styles = append(styles, fmt.Sprintf("border-radius: %s", g.resolveRounded(v)))
	}

	// Box shadow
	if v := g.getStringProp(elem, "shadow"); v != "" {
		styles = append(styles, fmt.Sprintf("box-shadow: %s", g.resolveShadow(v)))
	}

	// Width
	if v := g.getStringProp(elem, "width"); v != "" {
		styles = append(styles, fmt.Sprintf("width: %s", v))
	}

	// Height
	if v := g.getStringProp(elem, "height"); v != "" {
		styles = append(styles, fmt.Sprintf("height: %s", v))
	}

	// Line height / spacing
	if v := g.getStringProp(elem, "line_spacing"); v != "" {
		styles = append(styles, fmt.Sprintf("line-height: %s", v))
	}

	// Display
	if v := g.getStringProp(elem, "display"); v != "" {
		styles = append(styles, fmt.Sprintf("display: %s", v))
	}

	// Flex centering shortcut
	if v := g.getStringProp(elem, "center_content"); v == "true" {
		styles = append(styles, "display: flex", "justify-content: center", "align-items: center")
	}

	if len(styles) == 0 {
		return ""
	}
	return fmt.Sprintf(" style=\"%s;\"", strings.Join(styles, "; "))
}

// resolveFontSize converts friendly size names to CSS
func (g *Generator) resolveFontSize(size string) string {
	switch size {
	case "tiny":
		return "10px"
	case "small":
		return "12px"
	case "normal":
		return "16px"
	case "medium":
		return "20px"
	case "large":
		return "24px"
	case "huge":
		return "32px"
	case "giant":
		return "48px"
	default:
		return size // assume it's already a valid CSS value
	}
}

// resolveSpacing converts friendly spacing names to CSS
func (g *Generator) resolveSpacing(spacing string) string {
	switch spacing {
	case "none":
		return "0"
	case "tiny":
		return "4px"
	case "small":
		return "8px"
	case "medium":
		return "16px"
	case "large":
		return "24px"
	case "huge":
		return "32px"
	default:
		return spacing
	}
}

// resolveBorder converts friendly border syntax
func (g *Generator) resolveBorder(border string) string {
	switch border {
	case "thin":
		return "1px solid #ccc"
	case "medium":
		return "2px solid #999"
	case "thick":
		return "3px solid #333"
	case "none":
		return "none"
	default:
		return border
	}
}

// resolveRounded converts friendly rounded corner names
func (g *Generator) resolveRounded(rounded string) string {
	switch rounded {
	case "none":
		return "0"
	case "small":
		return "4px"
	case "medium":
		return "8px"
	case "large":
		return "16px"
	case "full":
		return "9999px"
	case "circle":
		return "50%"
	default:
		return rounded
	}
}

// resolveShadow converts friendly shadow names
func (g *Generator) resolveShadow(shadow string) string {
	switch shadow {
	case "none":
		return "none"
	case "small":
		return "0 1px 3px rgba(0,0,0,0.12), 0 1px 2px rgba(0,0,0,0.24)"
	case "medium":
		return "0 3px 6px rgba(0,0,0,0.15), 0 2px 4px rgba(0,0,0,0.12)"
	case "large":
		return "0 10px 20px rgba(0,0,0,0.15), 0 3px 6px rgba(0,0,0,0.10)"
	case "huge":
		return "0 15px 25px rgba(0,0,0,0.15), 0 5px 10px rgba(0,0,0,0.05)"
	default:
		return shadow
	}
}

// generateHeading generates <h1>-<h6> based on size
func (g *Generator) generateHeading(elem *ast.Element, indent string) string {
	content := g.getStringProp(elem, "contains")
	id := g.getStringProp(elem, "label")
	size := g.getStringProp(elem, "size")
	level := g.getStringProp(elem, "level")

	// Apply formatting
	content = g.applyFormatting(elem, content)

	// Determine heading level - default to h1
	if level == "" {
		level = "1"
	}

	// Build style - include size if specified
	styleAttr := g.buildStyleAttr(elem)
	if size != "" && styleAttr == "" {
		styleAttr = fmt.Sprintf(" style=\"font-size: %s;\"", size)
	} else if size != "" {
		// Append size to existing styles
		styleAttr = strings.TrimSuffix(styleAttr, "\"")
		styleAttr = styleAttr + fmt.Sprintf("; font-size: %s;\"", size)
	}

	idAttr := ""
	if id != "" {
		idAttr = fmt.Sprintf(" id=\"%s\"", id)
	}

	return fmt.Sprintf("%s<h%s%s%s>%s</h%s>\n", indent, level, idAttr, styleAttr, content, level)
}

// generateLink generates an <a> element
func (g *Generator) generateLink(elem *ast.Element, indent string) string {
	content := g.getStringProp(elem, "contains")
	// Support both link_url (LPML way) and href (legacy)
	href := g.getStringProp(elem, "link_url")
	if href == "" {
		href = g.getStringProp(elem, "href")
	}
	id := g.getStringProp(elem, "label")

	idAttr := ""
	if id != "" {
		idAttr = fmt.Sprintf(" id=\"%s\"", id)
	}

	return fmt.Sprintf("%s<a href=\"%s\"%s>%s</a>\n", indent, href, idAttr, content)
}

// generateImage generates an <img> element
func (g *Generator) generateImage(elem *ast.Element, indent string) string {
	src := g.getStringProp(elem, "src")
	alt := g.getStringProp(elem, "alt")
	id := g.getStringProp(elem, "label")

	idAttr := ""
	if id != "" {
		idAttr = fmt.Sprintf(" id=\"%s\"", id)
	}

	return fmt.Sprintf("%s<img src=\"%s\" alt=\"%s\"%s>\n", indent, src, alt, idAttr)
}

// generateList generates <ul> or <ol>
func (g *Generator) generateList(elem *ast.Element, indent string, ordered bool) string {
	var sb strings.Builder

	tag := "ul"
	if ordered {
		tag = "ol"
	}

	// Check for type property to override ordered/unordered
	listType := g.getStringProp(elem, "type")
	if listType == "ordered" {
		tag = "ol"
	} else if listType == "unordered" {
		tag = "ul"
	}

	id := g.getStringProp(elem, "label")
	idAttr := ""
	if id != "" {
		idAttr = fmt.Sprintf(" id=\"%s\"", id)
	}

	sb.WriteString(fmt.Sprintf("%s<%s%s>\n", indent, tag, idAttr))

	g.indent++
	childIndent := strings.Repeat("  ", g.indent)

	// Check if there's an items array property
	if itemsVal, exists := elem.Properties["items"]; exists {
		if arr, ok := itemsVal.(*ast.ArrayValue); ok {
			for _, item := range arr.Values {
				itemContent := g.resolveValue(item)
				sb.WriteString(fmt.Sprintf("%s<li>%s</li>\n", childIndent, itemContent))
			}
		}
	}

	// Also process any child elements
	for _, child := range elem.Children {
		sb.WriteString(g.generateNode(child))
	}
	g.indent--

	sb.WriteString(fmt.Sprintf("%s</%s>\n", indent, tag))
	return sb.String()
}

// generateListItem generates <li>
func (g *Generator) generateListItem(elem *ast.Element, indent string) string {
	content := g.getStringProp(elem, "contains")
	return fmt.Sprintf("%s<li>%s</li>\n", indent, content)
}

// generateTable generates <table>
func (g *Generator) generateTable(elem *ast.Element, indent string) string {
	var sb strings.Builder

	id := g.getStringProp(elem, "label")
	idAttr := ""
	if id != "" {
		idAttr = fmt.Sprintf(" id=\"%s\"", id)
	}

	sb.WriteString(fmt.Sprintf("%s<table%s>\n", indent, idAttr))

	g.indent++
	for _, child := range elem.Children {
		sb.WriteString(g.generateNode(child))
	}
	g.indent--

	sb.WriteString(fmt.Sprintf("%s</table>\n", indent))
	return sb.String()
}

// generateRow generates <tr>
func (g *Generator) generateRow(elem *ast.Element, indent string) string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("%s<tr>\n", indent))

	g.indent++
	for _, child := range elem.Children {
		sb.WriteString(g.generateNode(child))
	}
	g.indent--

	sb.WriteString(fmt.Sprintf("%s</tr>\n", indent))
	return sb.String()
}

// generateCell generates <td>
func (g *Generator) generateCell(elem *ast.Element, indent string) string {
	content := g.getStringProp(elem, "contains")
	return fmt.Sprintf("%s<td>%s</td>\n", indent, content)
}

// generateForm generates <form>
func (g *Generator) generateForm(elem *ast.Element, indent string) string {
	var sb strings.Builder

	action := g.getStringProp(elem, "action")
	id := g.getStringProp(elem, "label")

	idAttr := ""
	if id != "" {
		idAttr = fmt.Sprintf(" id=\"%s\"", id)
	}

	sb.WriteString(fmt.Sprintf("%s<form action=\"%s\"%s>\n", indent, action, idAttr))

	g.indent++
	for _, child := range elem.Children {
		sb.WriteString(g.generateNode(child))
	}
	g.indent--

	sb.WriteString(fmt.Sprintf("%s</form>\n", indent))
	return sb.String()
}

// generateInput generates <input>
func (g *Generator) generateInput(elem *ast.Element, indent string) string {
	inputType := g.getStringProp(elem, "type")
	name := g.getStringProp(elem, "name")
	id := g.getStringProp(elem, "label")

	idAttr := ""
	if id != "" {
		idAttr = fmt.Sprintf(" id=\"%s\"", id)
	}

	if inputType == "" {
		inputType = "text"
	}

	return fmt.Sprintf("%s<input type=\"%s\" name=\"%s\"%s>\n", indent, inputType, name, idAttr)
}

// generateButton generates <button>
func (g *Generator) generateButton(elem *ast.Element, indent string) string {
	content := g.getStringProp(elem, "contains")
	id := g.getStringProp(elem, "label")

	idAttr := ""
	if id != "" {
		idAttr = fmt.Sprintf(" id=\"%s\"", id)
	}

	return fmt.Sprintf("%s<button%s>%s</button>\n", indent, idAttr, content)
}

// generateBold generates <strong>
func (g *Generator) generateBold(elem *ast.Element, indent string) string {
	content := g.getStringProp(elem, "contains")
	return fmt.Sprintf("%s<strong>%s</strong>\n", indent, content)
}

// generateItalic generates <em>
func (g *Generator) generateItalic(elem *ast.Element, indent string) string {
	content := g.getStringProp(elem, "contains")
	return fmt.Sprintf("%s<em>%s</em>\n", indent, content)
}

// generateCode generates <pre><code> block
func (g *Generator) generateCode(elem *ast.Element, indent string) string {
	var sb strings.Builder

	// Check for linked_file property
	linkedFile := g.getStringProp(elem, "linked_file")
	fileType := g.getStringProp(elem, "file_type")

	// Get the syntax/code content
	var codeContent string
	if syntaxVal, exists := elem.Properties["syntax"]; exists {
		if cb, ok := syntaxVal.(*ast.CodeBlockValue); ok {
			codeContent = cb.Content
		}
	}

	// Determine language class for syntax highlighting
	langClass := ""
	if fileType != "" {
		langClass = fmt.Sprintf(" class=\"language-%s\"", fileType)
	}

	sb.WriteString(fmt.Sprintf("%s<pre><code%s>", indent, langClass))

	if linkedFile != "" {
		// If it's a linked file, add a comment showing the file path
		sb.WriteString(fmt.Sprintf("/* File: %s */\n", linkedFile))
		// Note: In a real implementation, you might read the file contents here
	}

	if codeContent != "" {
		// Escape HTML entities in code
		escaped := escapeHTML(codeContent)
		sb.WriteString(escaped)
	}

	sb.WriteString("</code></pre>\n")

	return sb.String()
}

// escapeHTML escapes HTML special characters
func escapeHTML(s string) string {
	s = strings.ReplaceAll(s, "&", "&amp;")
	s = strings.ReplaceAll(s, "<", "&lt;")
	s = strings.ReplaceAll(s, ">", "&gt;")
	s = strings.ReplaceAll(s, "\"", "&quot;")
	return s
}

// getStringProp gets a string property value from an element
func (g *Generator) getStringProp(elem *ast.Element, name string) string {
	if val, exists := elem.Properties[name]; exists {
		return g.resolveValue(val)
	}
	return ""
}

// resolveValue converts any Value to a string
func (g *Generator) resolveValue(val ast.Value) string {
	switch v := val.(type) {
	case *ast.StringValue:
		return v.Value
	case *ast.NumberValue:
		return v.Value
	case *ast.VariableRef:
		// Resolve variable reference
		if refElem, exists := g.labels[v.Name]; exists {
			// Get the contains of the referenced element
			return g.getStringProp(refElem, "contains")
		}
		return "$" + v.Name // Return as-is if not found
	case *ast.ArrayValue:
		// For arrays, join values with comma (for display purposes)
		var parts []string
		for _, item := range v.Values {
			parts = append(parts, g.resolveValue(item))
		}
		return strings.Join(parts, ", ")
	}
	return ""
}
