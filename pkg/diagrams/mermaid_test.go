package diagrams

import (
	"strings"
	"testing"
)

func TestParseMermaidFlowchart_TopDown(t *testing.T) {
	mermaid := `graph TD
    A[Start] --> B[End]`

	flow, err := ParseMermaidFlowchart(mermaid)
	if err != nil {
		t.Fatalf("ParseMermaidFlowchart failed: %v", err)
	}

	if flow.Direction != TopToBottom {
		t.Errorf("Expected TopToBottom, got %v", flow.Direction)
	}

	if len(flow.Nodes) != 2 {
		t.Errorf("Expected 2 nodes, got %d", len(flow.Nodes))
	}

	if len(flow.Edges) != 1 {
		t.Errorf("Expected 1 edge, got %d", len(flow.Edges))
	}
}

func TestParseMermaidFlowchart_LeftRight(t *testing.T) {
	mermaid := `graph LR
    A[Start] --> B[End]`

	flow, err := ParseMermaidFlowchart(mermaid)
	if err != nil {
		t.Fatalf("ParseMermaidFlowchart failed: %v", err)
	}

	if flow.Direction != LeftToRight {
		t.Errorf("Expected LeftToRight, got %v", flow.Direction)
	}
}

func TestParseMermaidFlowchart_NodeShapes(t *testing.T) {
	mermaid := `graph TD
    A[Box]
    B(Rounded)
    C{Diamond}
    D((Circle))`

	flow, err := ParseMermaidFlowchart(mermaid)
	if err != nil {
		t.Fatalf("ParseMermaidFlowchart failed: %v", err)
	}

	if len(flow.Nodes) != 4 {
		t.Fatalf("Expected 4 nodes, got %d", len(flow.Nodes))
	}

	// Check shapes
	shapes := map[string]NodeShape{
		"A": ShapeBox,
		"B": ShapeRounded,
		"C": ShapeDiamond,
		"D": ShapeCircle,
	}

	for _, node := range flow.Nodes {
		expectedShape, ok := shapes[node.ID]
		if !ok {
			t.Errorf("Unexpected node ID: %s", node.ID)
			continue
		}

		if node.Shape != expectedShape {
			t.Errorf("Node %s: expected shape %v, got %v", node.ID, expectedShape, node.Shape)
		}
	}
}

func TestParseMermaidFlowchart_EdgeLabels(t *testing.T) {
	mermaid := `graph TD
    A[Start] -->|Yes| B[End]
    B --> C[Done]`

	flow, err := ParseMermaidFlowchart(mermaid)
	if err != nil {
		t.Fatalf("ParseMermaidFlowchart failed: %v", err)
	}

	if len(flow.Edges) != 2 {
		t.Fatalf("Expected 2 edges, got %d", len(flow.Edges))
	}

	// First edge should have label
	if flow.Edges[0].Label != "Yes" {
		t.Errorf("Expected edge label 'Yes', got '%s'", flow.Edges[0].Label)
	}

	// Second edge should have no label
	if flow.Edges[1].Label != "" {
		t.Errorf("Expected empty edge label, got '%s'", flow.Edges[1].Label)
	}
}

func TestParseMermaidSequence(t *testing.T) {
	mermaid := `sequenceDiagram
    Alice->>Bob: Hello
    Bob-->>Alice: Hi`

	seq, err := ParseMermaidSequence(mermaid)
	if err != nil {
		t.Fatalf("ParseMermaidSequence failed: %v", err)
	}

	if len(seq.Actors) != 2 {
		t.Errorf("Expected 2 actors, got %d", len(seq.Actors))
	}

	if len(seq.Messages) != 2 {
		t.Errorf("Expected 2 messages, got %d", len(seq.Messages))
	}

	// Check first message
	msg1 := seq.Messages[0]
	if msg1.From != "Alice" {
		t.Errorf("Expected from 'Alice', got '%s'", msg1.From)
	}
	if msg1.To != "Bob" {
		t.Errorf("Expected to 'Bob', got '%s'", msg1.To)
	}
	if msg1.Label != "Hello" {
		t.Errorf("Expected label 'Hello', got '%s'", msg1.Label)
	}
	if msg1.Type != MessageSync {
		t.Errorf("Expected MessageSync, got %v", msg1.Type)
	}

	// Check second message (return)
	msg2 := seq.Messages[1]
	if msg2.Type != MessageReturn {
		t.Errorf("Expected MessageReturn, got %v", msg2.Type)
	}
}

func TestParseMermaidSequence_Participants(t *testing.T) {
	mermaid := `sequenceDiagram
    participant A as Alice
    participant B as Bob
    A->>B: Message`

	seq, err := ParseMermaidSequence(mermaid)
	if err != nil {
		t.Fatalf("ParseMermaidSequence failed: %v", err)
	}

	if len(seq.Actors) != 2 {
		t.Fatalf("Expected 2 actors, got %d", len(seq.Actors))
	}

	// Check actor names
	if seq.Actors[0].Name != "Alice" {
		t.Errorf("Expected actor name 'Alice', got '%s'", seq.Actors[0].Name)
	}

	if seq.Actors[1].Name != "Bob" {
		t.Errorf("Expected actor name 'Bob', got '%s'", seq.Actors[1].Name)
	}
}

func TestExtractMermaidFromMarkdown(t *testing.T) {
	markdown := `# Example

Some text

` + "```mermaid" + `
graph TD
    A[Start] --> B[End]
` + "```" + `

More text

` + "```mermaid" + `
sequenceDiagram
    A->>B: Hello
` + "```" + `
`

	blocks, err := ExtractMermaidFromMarkdown(markdown)
	if err != nil {
		t.Fatalf("ExtractMermaidFromMarkdown failed: %v", err)
	}

	if len(blocks) != 2 {
		t.Fatalf("Expected 2 blocks, got %d", len(blocks))
	}

	// Check first block
	if blocks[0].Type != "flowchart" {
		t.Errorf("Expected type 'flowchart', got '%s'", blocks[0].Type)
	}

	if blocks[0].Diagram == nil {
		t.Error("Expected diagram to be parsed, got nil")
	}

	// Check second block
	if blocks[1].Type != "sequenceDiagram" {
		t.Errorf("Expected type 'sequenceDiagram', got '%s'", blocks[1].Type)
	}

	if blocks[1].Diagram == nil {
		t.Error("Expected diagram to be parsed, got nil")
	}
}

func TestExtractMermaidFromMarkdown_Empty(t *testing.T) {
	markdown := `# No diagrams here

Just some regular markdown text.
`

	blocks, err := ExtractMermaidFromMarkdown(markdown)
	if err != nil {
		t.Fatalf("ExtractMermaidFromMarkdown failed: %v", err)
	}

	if len(blocks) != 0 {
		t.Errorf("Expected 0 blocks, got %d", len(blocks))
	}
}

func TestParseMermaidFlowchart_Render(t *testing.T) {
	mermaid := `graph TD
    A[Start] --> B[End]`

	flow, err := ParseMermaidFlowchart(mermaid)
	if err != nil {
		t.Fatalf("ParseMermaidFlowchart failed: %v", err)
	}

	output := flow.Render()

	if output == "" {
		t.Error("Expected non-empty render output")
	}

	// Should contain node labels
	if !strings.Contains(output, "Start") {
		t.Error("Expected output to contain 'Start'")
	}

	if !strings.Contains(output, "End") {
		t.Error("Expected output to contain 'End'")
	}
}

func TestParseMermaidSequence_Render(t *testing.T) {
	mermaid := `sequenceDiagram
    Alice->>Bob: Hello`

	seq, err := ParseMermaidSequence(mermaid)
	if err != nil {
		t.Fatalf("ParseMermaidSequence failed: %v", err)
	}

	output := seq.Render()

	if output == "" {
		t.Error("Expected non-empty render output")
	}

	// Should contain actor names
	if !strings.Contains(output, "Alice") {
		t.Error("Expected output to contain 'Alice'")
	}

	if !strings.Contains(output, "Bob") {
		t.Error("Expected output to contain 'Bob'")
	}

	// Should contain message
	if !strings.Contains(output, "Hello") {
		t.Error("Expected output to contain 'Hello'")
	}
}
