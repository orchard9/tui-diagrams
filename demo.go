package main

import (
	"fmt"

	"github.com/orchard9/tui-diagrams/pkg/diagrams"
)

func main() {
	printHeader("TUI DIAGRAMS - Quick Demo")
	fmt.Println("Terminal-friendly diagram rendering for Go")
	fmt.Println()

	// 1. Flowchart Example
	printSection("1. FLOWCHART (Vertical Layout)")
	flowchart := diagrams.NewFlowchart(diagrams.TopToBottom)

	flowchart.AddNode("start", "Start Process", diagrams.ShapeRounded).
		AddNode("input", "Get User Input", diagrams.ShapeBox).
		AddNode("validate", "Valid Input?", diagrams.ShapeDiamond).
		AddNode("process", "Process Data", diagrams.ShapeBox).
		AddNode("error", "Show Error", diagrams.ShapeBox).
		AddNode("done", "Complete", diagrams.ShapeCircle)

	flowchart.AddEdge("start", "input", "").
		AddEdge("input", "validate", "").
		AddEdge("validate", "process", "Yes").
		AddEdge("validate", "error", "No").
		AddEdge("process", "done", "").
		AddEdge("error", "input", "Retry")

	fmt.Println(flowchart.Render())
	fmt.Println()

	// 2. Sequence Diagram Example
	printSection("2. SEQUENCE DIAGRAM (Actor Interactions)")
	sequence := diagrams.NewSequenceDiagram()

	sequence.AddActor("client", "Client").
		AddActor("api", "API").
		AddActor("db", "Database")

	sequence.AddMessage("client", "api", "POST /users", diagrams.MessageSync).
		AddMessage("api", "db", "INSERT user", diagrams.MessageSync).
		AddMessage("db", "api", "User ID", diagrams.MessageReturn).
		AddMessage("api", "client", "201 Created", diagrams.MessageReturn)

	fmt.Println(sequence.Render())
	fmt.Println()

	// 3. Bar Chart Example (Horizontal)
	printSection("3. BAR CHART (Horizontal with Colors)")
	chart := diagrams.NewBarChart("Team Performance", diagrams.Horizontal)

	chart.AddBarWithColor("Frontend", 85, "\x1b[32m").  // Green
		AddBarWithColor("Backend", 92, "\x1b[34m").     // Blue
		AddBarWithColor("DevOps", 78, "\x1b[33m").      // Yellow
		AddBarWithColor("QA", 88, "\x1b[36m").          // Cyan
		SetWidth(40)

	fmt.Println(chart.Render())
	fmt.Println()

	// 4. Bar Chart Example (Vertical)
	printSection("4. BAR CHART (Vertical)")
	vertChart := diagrams.NewBarChart("Monthly Sales", diagrams.Vertical)

	vertChart.AddBarWithColor("Jan", 45, "\x1b[34m").  // Blue
		AddBarWithColor("Feb", 52, "\x1b[34m").
		AddBarWithColor("Mar", 68, "\x1b[32m").        // Green
		AddBarWithColor("Apr", 71, "\x1b[32m").
		SetWidth(6).
		SetHeight(12)

	fmt.Println(vertChart.Render())
	fmt.Println()

	// 5. Mermaid Parser Example
	printSection("5. MERMAID PARSER (Flowchart)")
	mermaidCode := `graph TD
    A[Login] --> B{Authenticated?}
    B -->|Yes| C[Dashboard]
    B -->|No| D[Error Page]`

	parsedFlow, err := diagrams.ParseMermaidFlowchart(mermaidCode)
	if err != nil {
		fmt.Printf("Error parsing Mermaid: %v\n", err)
	} else {
		fmt.Println(parsedFlow.Render())
	}
	fmt.Println()

	// 6. Mermaid Parser Example (Sequence)
	printSection("6. MERMAID PARSER (Sequence Diagram)")
	mermaidSeq := `sequenceDiagram
    User->>System: Login Request
    System->>Database: Check Credentials
    Database-->>System: Valid
    System-->>User: Access Token`

	parsedSeq, err := diagrams.ParseMermaidSequence(mermaidSeq)
	if err != nil {
		fmt.Printf("Error parsing Mermaid: %v\n", err)
	} else {
		fmt.Println(parsedSeq.Render())
	}
	fmt.Println()

	// Usage instructions
	printFooter()
}

func printHeader(title string) {
	width := 70
	fmt.Println(repeat("=", width))
	fmt.Printf("  %s\n", title)
	fmt.Println(repeat("=", width))
}

func printSection(title string) {
	fmt.Println(title)
	fmt.Println(repeat("-", len(title)))
	fmt.Println()
}

func printFooter() {
	fmt.Println(repeat("=", 70))
	fmt.Println("USAGE:")
	fmt.Println()
	fmt.Println("  go get github.com/orchard9/tui-diagrams")
	fmt.Println()
	fmt.Println("  import \"github.com/orchard9/tui-diagrams/pkg/diagrams\"")
	fmt.Println()
	fmt.Println("EXAMPLES:")
	fmt.Println("  • Flowcharts:  examples/flowchart/main.go")
	fmt.Println("  • Sequences:   examples/sequence/main.go")
	fmt.Println("  • Bar Charts:  examples/barchart/main.go")
	fmt.Println("  • Full Demo:   cmd/demo/main.go")
	fmt.Println()
	fmt.Println("MERMAID FILES:")
	fmt.Println("  go run cmd/mermaid-render/main.go examples/authentication.mmd")
	fmt.Println("  go run cmd/mermaid-render/main.go example.md")
	fmt.Println()
	fmt.Println("DOCS:")
	fmt.Println("  https://github.com/orchard9/tui-diagrams")
	fmt.Println("  https://pkg.go.dev/github.com/orchard9/tui-diagrams")
	fmt.Println(repeat("=", 70))
}

func repeat(s string, count int) string {
	result := ""
	for i := 0; i < count; i++ {
		result += s
	}
	return result
}
