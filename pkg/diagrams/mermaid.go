package diagrams

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

// ParseMermaidFlowchart parses Mermaid flowchart syntax and returns a Flowchart
//
// Supports syntax like:
//
//	graph TD
//	A[Start] --> B{Decision}
//	B -->|Yes| C[OK]
//	B -->|No| D[End]
func ParseMermaidFlowchart(mermaidText string) (*Flowchart, error) {
	lines := strings.Split(mermaidText, "\n")
	if len(lines) == 0 {
		return nil, fmt.Errorf("empty mermaid text")
	}

	// Parse direction from first line
	direction := TopToBottom
	firstLine := strings.TrimSpace(lines[0])
	if strings.HasPrefix(firstLine, "graph ") || strings.HasPrefix(firstLine, "flowchart ") {
		dirStr := strings.Fields(firstLine)
		if len(dirStr) > 1 {
			switch dirStr[1] {
			case "LR", "RL":
				direction = LeftToRight
			case "TD", "TB":
				direction = TopToBottom
			}
		}
	}

	flow := NewFlowchart(direction)
	nodes := make(map[string]bool)

	// Regex patterns
	// Match nodes: A[text], A(text), A{text}, A((text))
	nodeRegex := regexp.MustCompile(`([A-Za-z0-9_]+)(\(\(|[\[\(\{])([^\]\)\}]+)(\)\)|[\]\)\}])`)
	// Match edges: A --> B or A[text] --> B[text] or A -->|label| B
	// Allows optional shape syntax after node IDs
	edgeRegex := regexp.MustCompile(`([A-Za-z0-9_]+)(?:\[[^\]]+\]|\([^)]+\)|\{[^}]+\}|\(\([^)]+\)\))?\s*--+>\s*(?:\|([^|]+)\|)?\s*([A-Za-z0-9_]+)(?:\[[^\]]+\]|\([^)]+\)|\{[^}]+\}|\(\([^)]+\)\))?`)

	for i := 1; i < len(lines); i++ {
		line := strings.TrimSpace(lines[i])
		if line == "" || strings.HasPrefix(line, "%%") {
			continue
		}

		// Extract nodes
		nodeMatches := nodeRegex.FindAllStringSubmatch(line, -1)
		for _, match := range nodeMatches {
			if len(match) >= 4 {
				id := match[1]
				label := match[3]
				shapeStart := match[2]

				if !nodes[id] {
					shape := ShapeBox
					switch shapeStart {
					case "[":
						shape = ShapeBox
					case "(":
						shape = ShapeRounded
					case "((":
						shape = ShapeCircle
					case "{":
						shape = ShapeDiamond
					}

					flow.AddNode(id, label, shape)
					nodes[id] = true
				}
			}
		}

		// Extract edges
		edgeMatches := edgeRegex.FindAllStringSubmatch(line, -1)
		for _, match := range edgeMatches {
			if len(match) >= 4 {
				from := match[1]
				label := ""
				if len(match) > 2 && match[2] != "" {
					label = match[2]
				}
				to := match[3]
				flow.AddEdge(from, to, label)
			}
		}
	}

	return flow, nil
}

// ParseMermaidSequence parses Mermaid sequence diagram syntax
func ParseMermaidSequence(mermaidText string) (*SequenceDiagram, error) {
	lines := strings.Split(mermaidText, "\n")
	if len(lines) == 0 {
		return nil, fmt.Errorf("empty mermaid text")
	}

	seq := NewSequenceDiagram()
	actors := make(map[string]bool)

	msgRegex := regexp.MustCompile(`([A-Za-z0-9_]+)\s*(-->>|->>|-->|->)\s*([A-Za-z0-9_]+)\s*:\s*(.+)`)

	for i, line := range lines {
		line = strings.TrimSpace(line)

		if i == 0 || line == "" || strings.HasPrefix(line, "%%") {
			continue
		}

		// Parse participant declarations
		if strings.HasPrefix(line, "participant ") {
			parts := strings.Fields(line)
			if len(parts) >= 2 {
				id := parts[1]
				name := id
				if len(parts) > 3 && parts[2] == "as" {
					name = strings.Join(parts[3:], " ")
				}
				if !actors[id] {
					seq.AddActor(id, name)
					actors[id] = true
				}
			}
			continue
		}

		// Parse messages
		matches := msgRegex.FindStringSubmatch(line)
		if len(matches) >= 5 {
			from := matches[1]
			arrow := matches[2]
			to := matches[3]
			label := matches[4]

			if !actors[from] {
				seq.AddActor(from, from)
				actors[from] = true
			}
			if !actors[to] {
				seq.AddActor(to, to)
				actors[to] = true
			}

			msgType := MessageSync
			switch arrow {
			case "->>", "->":
				msgType = MessageSync
			case "-->>", "-->":
				msgType = MessageReturn
			}

			seq.AddMessage(from, to, label, msgType)
		}
	}

	return seq, nil
}

// MermaidBlock represents a Mermaid diagram found in Markdown
type MermaidBlock struct {
	Type    string  // "flowchart", "sequenceDiagram", etc.
	Content string  // The mermaid code
	Diagram Diagram // Parsed diagram (if successful)
}

// ExtractMermaidFromMarkdown extracts all ```mermaid blocks from Markdown content
func ExtractMermaidFromMarkdown(markdown string) ([]MermaidBlock, error) {
	var blocks []MermaidBlock
	lines := strings.Split(markdown, "\n")

	inMermaidBlock := false
	var currentBlock []string
	var blockType string

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)

		// Check for start of mermaid block
		if strings.HasPrefix(trimmed, "```mermaid") {
			inMermaidBlock = true
			currentBlock = []string{}
			continue
		}

		// Check for end of code block
		if inMermaidBlock && strings.HasPrefix(trimmed, "```") {
			// Parse the accumulated mermaid code
			content := strings.Join(currentBlock, "\n")

			// Determine diagram type from first line
			firstLine := strings.TrimSpace(currentBlock[0])
			if strings.HasPrefix(firstLine, "graph ") || strings.HasPrefix(firstLine, "flowchart ") {
				blockType = "flowchart"
			} else if strings.HasPrefix(firstLine, "sequenceDiagram") {
				blockType = "sequenceDiagram"
			} else {
				blockType = "unknown"
			}

			block := MermaidBlock{
				Type:    blockType,
				Content: content,
			}

			// Try to parse the diagram
			var diagram Diagram
			var err error

			switch blockType {
			case "flowchart":
				diagram, err = ParseMermaidFlowchart(content)
			case "sequenceDiagram":
				diagram, err = ParseMermaidSequence(content)
			}

			if err == nil && diagram != nil {
				block.Diagram = diagram
			}

			blocks = append(blocks, block)
			inMermaidBlock = false
			currentBlock = nil
			continue
		}

		// Accumulate lines in mermaid block
		if inMermaidBlock {
			currentBlock = append(currentBlock, line)
		}
	}

	return blocks, nil
}

// RenderMarkdownFile reads a Markdown file and renders all Mermaid diagrams
func RenderMarkdownFile(filename string) error {
	// Read file
	content, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	// Extract mermaid blocks
	blocks, err := ExtractMermaidFromMarkdown(string(content))
	if err != nil {
		return fmt.Errorf("failed to extract mermaid blocks: %w", err)
	}

	if len(blocks) == 0 {
		fmt.Println("No mermaid diagrams found in file")
		return nil
	}

	// Render each diagram
	for i, block := range blocks {
		fmt.Printf("\n=== Diagram %d (%s) ===\n\n", i+1, block.Type)

		if block.Diagram != nil {
			fmt.Println(block.Diagram.Render())
		} else {
			fmt.Printf("Unable to parse %s diagram\n", block.Type)
			fmt.Println("Content:")
			fmt.Println(block.Content)
		}
		fmt.Println()
	}

	return nil
}

// ParseMermaidFromFile is a convenience function to extract and parse diagrams from a file
func ParseMermaidFromFile(filename string) ([]Diagram, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	blocks, err := ExtractMermaidFromMarkdown(string(content))
	if err != nil {
		return nil, err
	}

	var diagrams []Diagram
	for _, block := range blocks {
		if block.Diagram != nil {
			diagrams = append(diagrams, block.Diagram)
		}
	}

	return diagrams, nil
}

// ParseMmdFile reads a standalone .mmd file (Mermaid diagram file) and parses it
func ParseMmdFile(filename string) (Diagram, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	mermaidText := string(content)
	lines := strings.Split(mermaidText, "\n")
	if len(lines) == 0 {
		return nil, fmt.Errorf("empty file")
	}

	firstLine := strings.TrimSpace(lines[0])

	// Detect diagram type from first line
	if strings.HasPrefix(firstLine, "graph ") || strings.HasPrefix(firstLine, "flowchart ") {
		return ParseMermaidFlowchart(mermaidText)
	}

	if strings.HasPrefix(firstLine, "sequenceDiagram") {
		return ParseMermaidSequence(mermaidText)
	}

	return nil, fmt.Errorf("unsupported mermaid diagram type: %s", firstLine)
}

// RenderMmdFile reads and renders a standalone .mmd file
func RenderMmdFile(filename string) error {
	diagram, err := ParseMmdFile(filename)
	if err != nil {
		return err
	}

	fmt.Println(diagram.Render())
	return nil
}
