package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"lpml/generator"
	"lpml/lexer"
	"lpml/parser"
)

func main() {
	fmt.Println("LAZY PAGE MAKER LANG")

	if len(os.Args) < 2 {
		fmt.Println("Usage: lpml <input.lpml> [output.html]")
		fmt.Println("  If output file is not specified, it will use the input filename with .html extension")
		os.Exit(1)
	}

	inputFile := os.Args[1]

	// Validate file extension
	if !checkFileType(inputFile) {
		log.Fatal("Invalid file type: needs to end in suffix .lpml")
	}

	// Determine output file
	outputFile := strings.TrimSuffix(inputFile, ".lpml") + ".html"
	if len(os.Args) >= 3 {
		outputFile = os.Args[2]
	}

	// Read input file
	content, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	// Lex the input
	l := lexer.New(string(content))

	// Parse into AST
	p := parser.New(l)
	doc := p.ParseDocument()

	// Check for parsing errors
	if len(p.Errors()) > 0 {
		fmt.Println("Parsing errors:")
		for _, e := range p.Errors() {
			fmt.Printf("  - %s\n", e)
		}
		os.Exit(1)
	}

	// Generate HTML
	gen := generator.New()
	html := gen.Generate(doc)

	// Write output file
	err = os.WriteFile(outputFile, []byte(html), 0644)
	if err != nil {
		log.Fatalf("Failed to write output file: %v", err)
	}

	fmt.Printf("Successfully generated: %s\n", outputFile)
}

func checkFileType(filename string) bool {
	return strings.HasSuffix(filename, ".lpml")
}
