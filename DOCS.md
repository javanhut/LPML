# LPML Documentation
## Lazy Page Maker Language

LPML is a simple, human-readable markup language that compiles to HTML. It's designed for people who want to create web pages without dealing with HTML/CSS complexity.

---

## Table of Contents

1. [Getting Started](#getting-started)
2. [Basic Syntax](#basic-syntax)
3. [Page Structure](#page-structure)
4. [Elements](#elements)
5. [Styling](#styling)
6. [Text Formatting](#text-formatting)
7. [Lists](#lists)
8. [Code Blocks](#code-blocks)
9. [Links & Images](#links--images)
10. [Tables](#tables)
11. [Forms](#forms)
12. [Variables & References](#variables--references)
13. [Complete Example](#complete-example)

---

## Getting Started

### Installation

Build the LPML compiler:

```bash
go build -o lpml .
```

### Usage

```bash
# Convert .lpml to .html (outputs same filename with .html extension)
./lpml mypage.lpml

# Specify output filename
./lpml mypage.lpml output.html
```

### Your First LPML File

Create a file called `hello.lpml`:

```
[top-of-page-start]
  [h-start]
    contains = "Hello World!"
  [h-end]
[top-of-page-end]

[mid-page-start]
  [p-start]
    contains = "Welcome to LPML!"
  [p-end]
[mid-page-end]
```

Compile it:

```bash
./lpml hello.lpml
```

Open `hello.html` in your browser!

---

## Basic Syntax

### Elements

All elements follow this pattern:

```
[element-start]
  property = "value"
  another_property = "another value"
[element-end]
```

### Properties

Properties are assigned using `=` with quoted values:

```
property = "value"
```

### Arrays

Some properties accept arrays:

```
items = [1, 2, 3, 4, 5]
items = ["apple", "banana", "cherry"]
format_with = ["bold", "italic"]
```

### Code Blocks

Multi-line content uses curly braces:

```
syntax = {
let x = 10;
console.log(x);
}
```

---

## Page Structure

Every LPML page has three sections:

```
[top-of-page-start]
  <!-- Header content -->
[top-of-page-end]

[mid-page-start]
  <!-- Main content -->
[mid-page-end]

[bottom-of-page-start]
  <!-- Footer content -->
[bottom-of-page-end]
```

Each section creates a `<div>` with the corresponding class (`top-of-page`, `mid-page`, `bottom-of-page`).

---

## Elements

### Headings

```
[h-start]
  contains = "Your Heading Text"
  level = "1"              # 1-6, defaults to 1
  size = "32px"            # Custom size (optional)
[h-end]
```

### Paragraphs

```
[p-start]
  contains = "Your paragraph text goes here."
[p-end]
```

### Divisions (Containers)

```
[divide-start]
  label = "my_container"   # Optional ID
  class = "my-class"       # Optional CSS class

  [p-start]
    contains = "Content inside the div"
  [p-end]
[divide-end]
```

### Buttons

```
[btn-start]
  contains = "Click Me"
  label = "submit_btn"
[btn-end]
```

---

## Styling

LPML makes styling easy with human-readable properties. No CSS knowledge required!

### Colors

| Property | Description | Example |
|----------|-------------|---------|
| `color` | Text color | `"red"`, `"#333"`, `"rgb(0,0,0)"` |
| `text_color` | Same as color | `"blue"` |
| `background` | Background color/gradient | `"#f5f5f5"` |
| `bg_color` | Same as background | `"white"` |

### Text Size

| Property | Description |
|----------|-------------|
| `text_size` | Font size |

**Friendly values:** `tiny`, `small`, `normal`, `medium`, `large`, `huge`, `giant`

Or use exact values: `"16px"`, `"1.5em"`, `"2rem"`

| Name | Size |
|------|------|
| `tiny` | 10px |
| `small` | 12px |
| `normal` | 16px |
| `medium` | 20px |
| `large` | 24px |
| `huge` | 32px |
| `giant` | 48px |

### Spacing

| Property | Description |
|----------|-------------|
| `padding` | Inside spacing |
| `margin` | Outside spacing |

**Friendly values:** `none`, `tiny`, `small`, `medium`, `large`, `huge`

| Name | Size |
|------|------|
| `none` | 0 |
| `tiny` | 4px |
| `small` | 8px |
| `medium` | 16px |
| `large` | 24px |
| `huge` | 32px |

### Borders

| Property | Description |
|----------|-------------|
| `border` | Border style |
| `rounded` | Corner radius |

**Border values:** `none`, `thin`, `medium`, `thick`

| Name | Style |
|------|-------|
| `thin` | 1px solid #ccc |
| `medium` | 2px solid #999 |
| `thick` | 3px solid #333 |

**Rounded values:** `none`, `small`, `medium`, `large`, `full`, `circle`

| Name | Radius |
|------|--------|
| `none` | 0 |
| `small` | 4px |
| `medium` | 8px |
| `large` | 16px |
| `full` | 9999px (pill shape) |
| `circle` | 50% |

### Shadows

| Property | Description |
|----------|-------------|
| `shadow` | Box shadow |

**Values:** `none`, `small`, `medium`, `large`, `huge`

### Layout

| Property | Description | Example |
|----------|-------------|---------|
| `align` | Text alignment | `"left"`, `"center"`, `"right"` |
| `width` | Element width | `"100%"`, `"300px"` |
| `height` | Element height | `"200px"`, `"auto"` |
| `center_content` | Center children | `"true"` |

### Font

| Property | Description | Example |
|----------|-------------|---------|
| `font` | Font family | `"Arial"`, `"Georgia"` |
| `line_spacing` | Line height | `"1.5"`, `"2em"` |

### Complete Styling Example

```
[divide-start]
  background = "#f0f0f0"
  padding = "large"
  margin = "medium"
  border = "thin"
  rounded = "large"
  shadow = "medium"

  [p-start]
    contains = "Beautifully styled content!"
    color = "#333"
    text_size = "large"
    align = "center"
    font = "Georgia"
  [p-end]
[divide-end]
```

### Gradients

Use CSS gradient syntax in the `background` property:

```
[divide-start]
  background = "linear-gradient(to right, #667eea, #764ba2)"
  padding = "huge"
  rounded = "large"
[divide-end]
```

---

## Text Formatting

Use the `format_with` property to apply text formatting:

```
[p-start]
  contains = "This text is bold and italic!"
  format_with = ["bold", "italic"]
[p-end]
```

### Available Formats

| Format | HTML Output | Description |
|--------|-------------|-------------|
| `"bold"` | `<strong>` | Bold text |
| `"italic"` | `<em>` | Italic text |
| `"underline"` | `<u>` | Underlined text |
| `"strike"` | `<s>` | Strikethrough text |
| `"code"` | `<code>` | Inline code |
| `"mark"` | `<mark>` | Highlighted text |

### Combining Formats

```
[p-start]
  contains = "Important highlighted text"
  format_with = ["bold", "mark", "underline"]
[p-end]
```

---

## Lists

### Ordered List (Numbered)

```
[lst-ord]
  items = [1, 2, 3, 4, 5]
[lst-end]

[lst-ord]
  items = ["First item", "Second item", "Third item"]
[lst-end]
```

### Unordered List (Bullets)

```
[lst-unord]
  items = ["Apple", "Banana", "Cherry"]
[lst-end]
```

### Manual List Items

```
[list-start]
  [item-start]
    contains = "First item"
  [item-end]
  [item-start]
    contains = "Second item"
  [item-end]
[list-end]
```

---

## Code Blocks

### Inline Code with Syntax Highlighting

```
[code-start]
  file_type = "javascript"
  syntax = {
function hello(name) {
    return `Hello, ${name}!`;
}
console.log(hello("World"));
}
[code-end]
```

The `file_type` adds a CSS class for syntax highlighters like Prism.js or highlight.js.

**Supported file types:** `javascript`, `python`, `go`, `html`, `css`, `json`, `bash`, `sql`, etc.

### Linking External Files

```
[code-start]
  linked_file = "../../src/index.js"
[code-end]
```

This adds a comment showing the file path (useful for documentation).

---

## Links & Images

### Links

```
[link-start]
  contains = "Click here to visit"
  link_url = "https://example.com"
[link-end]
```

### Images

```
[img-start]
  src = "path/to/image.jpg"
  alt = "Description of image"
  label = "my_image"
[img-end]
```

---

## Tables

```
[table-start]
  label = "my_table"

  [row-start]
    [cell-start]
      contains = "Header 1"
    [cell-end]
    [cell-start]
      contains = "Header 2"
    [cell-end]
  [row-end]

  [row-start]
    [cell-start]
      contains = "Data 1"
    [cell-end]
    [cell-start]
      contains = "Data 2"
    [cell-end]
  [row-end]

[table-end]
```

---

## Forms

```
[form-start]
  action = "/submit"
  label = "contact_form"

  [input-start]
    type = "text"
    name = "username"
    label = "user_input"
  [input-end]

  [input-start]
    type = "email"
    name = "email"
  [input-end]

  [btn-start]
    contains = "Submit"
  [btn-end]

[form-end]
```

### Input Types

Use standard HTML input types: `text`, `email`, `password`, `number`, `date`, `checkbox`, `radio`, etc.

---

## Variables & References

### Labels

Add a `label` to any element to reference it later:

```
[h-start]
  label = "main_title"
  contains = "Welcome!"
[h-end]
```

### Variable References

Reference other elements using `$label_name`:

```
[divide-start]
  linked_header = $main_title

  [p-start]
    contains = "Content linked to the header above"
  [p-end]
[divide-end]
```

---

## Complete Example

```
[top-of-page-start]
  [divide-start]
    background = "#2c3e50"
    padding = "large"

    [h-start]
      contains = "My Awesome Website"
      color = "white"
      align = "center"
      text_size = "giant"
    [h-end]
  [divide-end]
[top-of-page-end]

[mid-page-start]
  [divide-start]
    background = "white"
    padding = "large"
    margin = "medium"
    rounded = "large"
    shadow = "medium"

    [h-start]
      contains = "About Us"
      level = "2"
      color = "#2c3e50"
    [h-end]

    [p-start]
      contains = "We make web development easy with LPML!"
      text_size = "medium"
      color = "#666"
    [p-end]

    [p-start]
      contains = "No more wrestling with HTML and CSS."
      format_with = ["bold", "italic"]
      color = "#3498db"
    [p-end]
  [divide-end]

  [divide-start]
    background = "linear-gradient(135deg, #667eea 0%, #764ba2 100%)"
    padding = "huge"
    rounded = "large"
    margin = "medium"
    center_content = "true"

    [h-start]
      contains = "Features"
      color = "white"
      align = "center"
    [h-end]

    [lst-unord]
      items = ["Easy syntax", "No CSS needed", "Fast compilation", "Beautiful output"]
    [lst-end]
  [divide-end]

  [divide-start]
    background = "#f8f9fa"
    padding = "large"
    margin = "medium"
    rounded = "medium"

    [h-start]
      contains = "Sample Code"
      level = "2"
    [h-end]

    [code-start]
      file_type = "javascript"
      syntax = {
// LPML makes coding examples easy!
const greeting = "Hello, LPML!";
console.log(greeting);
}
    [code-end]
  [divide-end]
[mid-page-end]

[bottom-of-page-start]
  [divide-start]
    background = "#2c3e50"
    padding = "medium"

    [p-start]
      contains = "Made with LPML - The Lazy Page Maker Language"
      color = "white"
      align = "center"
      text_size = "small"
    [p-end]
  [divide-end]
[bottom-of-page-end]
```

---

## Quick Reference

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
| `[item-start]...[item-end]` | List item |
| `[table-start]...[table-end]` | Table |
| `[row-start]...[row-end]` | Table row |
| `[cell-start]...[cell-end]` | Table cell |
| `[form-start]...[form-end]` | Form |
| `[input-start]...[input-end]` | Input field |
| `[btn-start]...[btn-end]` | Button |
| `[code-start]...[code-end]` | Code block |

### Common Properties

| Property | Used On | Description |
|----------|---------|-------------|
| `contains` | Text elements | The text content |
| `label` | Any element | ID for referencing |
| `format_with` | Text elements | Text formatting array |
| `link_url` | Links | URL destination |
| `src` | Images | Image source path |
| `alt` | Images | Alt text |
| `items` | Lists | Array of list items |
| `file_type` | Code | Programming language |
| `syntax` | Code | Code content block |
| `action` | Forms | Form submission URL |
| `type` | Inputs | Input type |
| `name` | Inputs | Input name |

### All Style Properties

| Property | Values |
|----------|--------|
| `color` / `text_color` | Any color value |
| `background` / `bg_color` | Color or gradient |
| `text_size` | tiny/small/normal/medium/large/huge/giant or px |
| `font` | Font family name |
| `align` | left/center/right |
| `padding` | none/tiny/small/medium/large/huge or px |
| `margin` | none/tiny/small/medium/large/huge or px |
| `border` | none/thin/medium/thick |
| `rounded` | none/small/medium/large/full/circle |
| `shadow` | none/small/medium/large/huge |
| `width` | Any CSS width |
| `height` | Any CSS height |
| `center_content` | "true" to center children |
| `line_spacing` | Line height value |

---

## Tips

1. **Start simple** - Begin with just headings and paragraphs
2. **Use containers** - Wrap related content in `[divide-start]...[divide-end]`
3. **Friendly sizes** - Use words like `"large"` instead of memorizing pixel values
4. **Combine styles** - Mix colors, padding, shadows for beautiful cards
5. **Gradients are easy** - Just paste CSS gradient syntax into `background`

---

*LPML - Because life's too short for HTML boilerplate.*
