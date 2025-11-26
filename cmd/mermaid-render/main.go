package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/orchard9/tui-diagrams/pkg/diagrams"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run cmd/mermaid-render/main.go <file>")
		fmt.Println()
		fmt.Println("Renders Mermaid diagrams from Markdown (.md) or Mermaid (.mmd) files")
		fmt.Println()
		fmt.Println("Examples:")
		fmt.Println("  go run cmd/mermaid-render/main.go example.md              # Markdown with ```mermaid blocks")
		fmt.Println("  go run cmd/mermaid-render/main.go examples/auth.mmd       # Standalone .mmd file")
		fmt.Println("  go run cmd/mermaid-render/main.go README.md               # Any markdown file")
		os.Exit(1)
	}

	filename := os.Args[1]

	// Check if file exists
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		log.Fatalf("File not found: %s", filename)
	}

	ext := strings.ToLower(filepath.Ext(filename))

	fmt.Printf("Rendering Mermaid diagram(s) from: %s\n", filename)
	fmt.Println("================================================================================")

	var err error

	switch ext {
	case ".mmd":
		// Standalone Mermaid file
		err = diagrams.RenderMmdFile(filename)
	case ".md", ".markdown":
		// Markdown file with ```mermaid blocks
		err = diagrams.RenderMarkdownFile(filename)
	default:
		log.Fatalf("Unsupported file type: %s (use .md or .mmd)", ext)
	}

	if err != nil {
		log.Fatalf("Error rendering diagrams: %v", err)
	}

	fmt.Println("================================================================================")
	fmt.Println("Done!")
}
