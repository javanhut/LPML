# LPML - Lazy Page Maker Language

A simple, human-readable markup language that compiles to HTML. Build beautiful web pages without wrestling with HTML tags or CSS properties.

## Why LPML?

- **Simple syntax** - No angle brackets, no closing tag confusion
- **Easy styling** - Use words like `"large"` and `"medium"` instead of pixel values
- **Fast compilation** - Instantly generates clean HTML
- **Zero dependencies** - Single binary, works anywhere

## Quick Example

**Input (`hello.lpml`):**
```
[mid-page-start]
  [divide-start]
    background = "#f5f5f5"
    padding = "large"
    rounded = "medium"
    shadow = "medium"

    [h-start]
      contains = "Hello World!"
      color = "navy"
      align = "center"
    [h-end]

    [p-start]
      contains = "This is so much easier than HTML!"
      format_with = ["bold", "italic"]
      color = "#666"
    [p-end]
  [divide-end]
[mid-page-end]
```

**Output (`hello.html`):**
```html
<div class="mid-page">
  <div style="background-color: #f5f5f5; padding: 24px; border-radius: 8px; box-shadow: 0 3px 6px rgba(0,0,0,0.15);">
    <h1 style="color: navy; text-align: center;">Hello World!</h1>
    <p style="color: #666;"><strong><em>This is so much easier than HTML!</em></strong></p>
  </div>
</div>
```

## Installation

```bash
git clone https://github.com/yourusername/lpml.git
cd lpml
go build -o lpml .
```

## Usage

```bash
# Basic usage
./lpml mypage.lpml

# Specify output file
./lpml mypage.lpml output.html
```

## Features

### Page Structure
```
[top-of-page-start]
  <!-- Header -->
[top-of-page-end]

[mid-page-start]
  <!-- Main content -->
[mid-page-end]

[bottom-of-page-start]
  <!-- Footer -->
[bottom-of-page-end]
```

### Easy Styling

No CSS knowledge required! Use friendly property names:

| Property | Example Values |
|----------|---------------|
| `color` | `"red"`, `"#333"`, `"rgb(0,0,0)"` |
| `background` | `"#f5f5f5"`, `"linear-gradient(...)"` |
| `text_size` | `"tiny"`, `"small"`, `"medium"`, `"large"`, `"huge"` |
| `padding` | `"none"`, `"small"`, `"medium"`, `"large"` |
| `rounded` | `"small"`, `"medium"`, `"large"`, `"full"` |
| `shadow` | `"small"`, `"medium"`, `"large"` |
| `align` | `"left"`, `"center"`, `"right"` |

### Text Formatting

```
[p-start]
  contains = "Bold and italic text!"
  format_with = ["bold", "italic"]
[p-end]
```

Formats: `bold`, `italic`, `underline`, `strike`, `code`, `mark`

### Lists with Array Syntax

```
[lst-ord]
  items = [1, 2, 3, 4, 5]
[lst-end]

[lst-unord]
  items = ["Apple", "Banana", "Cherry"]
[lst-end]
```

### Code Blocks

```
[code-start]
  file_type = "javascript"
  syntax = {
function hello() {
    console.log("Hello!");
}
}
[code-end]
```

### All Elements

| Element | Description |
|---------|-------------|
| `[h-start]...[h-end]` | Heading |
| `[p-start]...[p-end]` | Paragraph |
| `[divide-start]...[divide-end]` | Container/div |
| `[link-start]...[link-end]` | Hyperlink |
| `[img-start]...[img-end]` | Image |
| `[lst-ord]...[lst-end]` | Ordered list |
| `[lst-unord]...[lst-end]` | Unordered list |
| `[table-start]...[table-end]` | Table |
| `[form-start]...[form-end]` | Form |
| `[code-start]...[code-end]` | Code block |
| `[btn-start]...[btn-end]` | Button |

## Examples

Check out the `examples/` folder:

- `landing.lpml` - Product landing page
- `portfolio.lpml` - Developer portfolio

Compile examples:
```bash
./lpml examples/landing.lpml examples/landing.html
./lpml examples/portfolio.lpml examples/portfolio.html
```

## Documentation

See [DOCS.md](DOCS.md) for complete documentation.

## Project Structure

```
lpml/
├── main.go              # CLI entry point
├── tokens/tokens.go     # Token definitions
├── lexer/lexer.go       # Tokenizer
├── ast/ast.go           # AST node definitions
├── parser/parser.go     # Parser
├── generator/generator.go # HTML generator
├── examples/            # Example LPML files
├── DOCS.md              # Full documentation
└── README.md            # This file
```

## Contributing

Contributions welcome! Feel free to:

- Report bugs
- Suggest new features
- Add more styling options
- Improve documentation

## License

MIT License - do whatever you want with it.

---

*LPML - Because life's too short for HTML boilerplate.*
