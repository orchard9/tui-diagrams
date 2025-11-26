package main

import (
	"fmt"
	"strings"

	"github.com/orchard9/tui-diagrams/pkg/diagrams"
)

func main() {
	printHeader("TUI-DIAGRAMS SHOWCASE")
	fmt.Println("Terminal-friendly diagram rendering for Go")
	fmt.Println()

	// Flowchart: User Authentication Flow
	printSection("FLOWCHART: User Authentication Flow")
	authFlow := diagrams.NewFlowchart(diagrams.TopToBottom)
	authFlow.AddNode("start", "User Login", diagrams.ShapeRounded).
		AddNode("input", "Enter Credentials", diagrams.ShapeBox).
		AddNode("validate", "Valid Format?", diagrams.ShapeDiamond).
		AddNode("check", "Check Database", diagrams.ShapeBox).
		AddNode("exists", "User Exists?", diagrams.ShapeDiamond).
		AddNode("verify", "Verify Password", diagrams.ShapeBox).
		AddNode("match", "Password Match?", diagrams.ShapeDiamond).
		AddNode("token", "Generate JWT", diagrams.ShapeBox).
		AddNode("error", "Show Error", diagrams.ShapeBox).
		AddNode("success", "Login Success", diagrams.ShapeRounded)

	authFlow.AddEdge("start", "input", "").
		AddEdge("input", "validate", "").
		AddEdge("validate", "check", "yes").
		AddEdge("validate", "error", "no").
		AddEdge("check", "exists", "").
		AddEdge("exists", "verify", "yes").
		AddEdge("exists", "error", "no").
		AddEdge("verify", "match", "").
		AddEdge("match", "token", "yes").
		AddEdge("match", "error", "no").
		AddEdge("token", "success", "").
		AddEdge("error", "input", "retry")

	fmt.Println(authFlow.Render())
	fmt.Println()

	// Horizontal Flowchart: CI/CD Pipeline
	printSection("FLOWCHART: CI/CD Pipeline (Horizontal)")
	cicdFlow := diagrams.NewFlowchart(diagrams.LeftToRight)
	cicdFlow.AddNode("commit", "Commit", diagrams.ShapeRounded).
		AddNode("build", "Build", diagrams.ShapeBox).
		AddNode("test", "Test", diagrams.ShapeBox).
		AddNode("deploy", "Deploy", diagrams.ShapeBox).
		AddNode("done", "Live", diagrams.ShapeCircle)

	cicdFlow.AddEdge("commit", "build", "").
		AddEdge("build", "test", "").
		AddEdge("test", "deploy", "").
		AddEdge("deploy", "done", "")

	fmt.Println(cicdFlow.Render())
	fmt.Println()

	// Sequence Diagram: Microservices Communication
	printSection("SEQUENCE DIAGRAM: E-Commerce Order Processing")
	seq := diagrams.NewSequenceDiagram()
	seq.AddActor("client", "Client").
		AddActor("api", "API Gateway").
		AddActor("order", "Order Service").
		AddActor("payment", "Payment").
		AddActor("inventory", "Inventory")

	seq.AddMessage("client", "api", "POST /orders", diagrams.MessageSync).
		AddMessage("api", "order", "Create Order", diagrams.MessageSync).
		AddMessage("order", "payment", "Charge Card", diagrams.MessageSync).
		AddMessage("payment", "order", "Success", diagrams.MessageReturn).
		AddMessage("order", "inventory", "Reserve Items", diagrams.MessageSync).
		AddMessage("inventory", "order", "Reserved", diagrams.MessageReturn).
		AddMessage("order", "order", "Save to DB", diagrams.MessageSync).
		AddMessage("order", "api", "Order Created", diagrams.MessageReturn).
		AddMessage("api", "client", "201 Created", diagrams.MessageReturn)

	fmt.Println(seq.Render())
	fmt.Println()

	// Horizontal Bar Chart: Programming Language Popularity
	printSection("BAR CHART: Programming Language Popularity 2024")
	langChart := diagrams.NewBarChart("", diagrams.Horizontal)
	langChart.AddBar("Python", 100).
		AddBar("JavaScript", 98).
		AddBar("TypeScript", 87).
		AddBar("Java", 82).
		AddBar("C#", 68).
		AddBar("Go", 65).
		AddBar("Rust", 54).
		AddBar("Swift", 41).
		SetWidth(45)

	fmt.Println(langChart.Render())
	fmt.Println()

	// Vertical Bar Chart with Colors: Team Sprint Velocity
	printSection("BAR CHART: Team Sprint Velocity (Colored)")
	velocityChart := diagrams.NewBarChart("Story Points Completed per Sprint", diagrams.Vertical)
	velocityChart.AddBarWithColor("Sprint 1", 23, "\x1b[31m"). // Red
									AddBarWithColor("Sprint 2", 28, "\x1b[33m"). // Yellow
									AddBarWithColor("Sprint 3", 31, "\x1b[33m"). // Yellow
									AddBarWithColor("Sprint 4", 35, "\x1b[32m"). // Green
									AddBarWithColor("Sprint 5", 42, "\x1b[32m"). // Green
									AddBarWithColor("Sprint 6", 38, "\x1b[32m"). // Green
									SetWidth(9).
									SetHeight(15)

	fmt.Println(velocityChart.Render())
	fmt.Println()

	// Horizontal Bar Chart: System Resource Usage
	printSection("BAR CHART: System Resource Usage")
	resourceChart := diagrams.NewBarChart("Current Resource Utilization", diagrams.Horizontal)
	resourceChart.AddBarWithColor("CPU Usage", 67, "\x1b[33m"). // Yellow
									AddBarWithColor("Memory Usage", 84, "\x1b[31m"). // Red (high)
									AddBarWithColor("Disk I/O", 45, "\x1b[32m").     // Green
									AddBarWithColor("Network", 23, "\x1b[32m").      // Green
									AddBarWithColor("GPU Usage", 91, "\x1b[35m").    // Magenta (very high)
									SetWidth(40)

	fmt.Println(resourceChart.Render())
	fmt.Println()

	// Comparison Chart: Before/After Optimization
	printSection("BAR CHART: Performance Optimization Results")
	perfChart := diagrams.NewBarChart("", diagrams.Horizontal)
	perfChart.AddBarWithColor("Response Time (Before)", 850, "\x1b[31m"). // Red
										AddBarWithColor("Response Time (After)", 125, "\x1b[32m"). // Green
										AddBarWithColor("Memory Usage (Before)", 512, "\x1b[31m"). // Red
										AddBarWithColor("Memory Usage (After)", 89, "\x1b[32m").   // Green
										AddBarWithColor("Error Rate (Before)", 12, "\x1b[31m").    // Red
										AddBarWithColor("Error Rate (After)", 1, "\x1b[32m").      // Green
										SetWidth(35)

	fmt.Println(perfChart.Render())
	fmt.Println()

	// Project Status Dashboard
	printSection("BAR CHART: Project Status Dashboard")
	projectChart := diagrams.NewBarChart("", diagrams.Horizontal)
	projectChart.AddBarWithColor("Completed Tasks", 142, "\x1b[32m"). // Green
										AddBarWithColor("In Progress", 23, "\x1b[33m").    // Yellow
										AddBarWithColor("Blocked", 8, "\x1b[31m").         // Red
										AddBarWithColor("Pending Review", 15, "\x1b[36m"). // Cyan
										AddBarWithColor("Ready to Start", 34, "\x1b[34m"). // Blue
										SetWidth(35)

	fmt.Println(projectChart.Render())
	fmt.Println()

	// Team Distribution Chart
	printSection("BAR CHART: Team Distribution by Department")
	teamChart := diagrams.NewBarChart("Headcount by Department", diagrams.Vertical)
	teamChart.AddBarWithColor("Engineering", 45, "\x1b[34m"). // Blue
									AddBarWithColor("Product", 12, "\x1b[35m").   // Magenta
									AddBarWithColor("Design", 8, "\x1b[36m").     // Cyan
									AddBarWithColor("Sales", 28, "\x1b[32m").     // Green
									AddBarWithColor("Marketing", 15, "\x1b[33m"). // Yellow
									AddBarWithColor("Support", 18, "\x1b[31m").   // Red
									SetWidth(11).
									SetHeight(14)

	fmt.Println(teamChart.Render())
	fmt.Println()

	// Mermaid Syntax Examples
	printSection("MERMAID: Flowchart from Syntax")

	mermaidFlow := `graph TD
    Start[Begin] --> Check{Ready?}
    Check -->|Yes| Process[Execute]
    Check -->|No| Wait[Wait]
    Process --> Done[Complete]
    Wait --> Check`

	parsedFlow, err := diagrams.ParseMermaidFlowchart(mermaidFlow)
	if err == nil {
		fmt.Println(parsedFlow.Render())
	}
	fmt.Println()

	printSection("MERMAID: Sequence Diagram from Syntax")

	mermaidSeq := `sequenceDiagram
    User->>System: Login Request
    System->>Database: Validate Credentials
    Database-->>System: User Data
    System->>System: Generate Session
    System-->>User: Access Token`

	parsedSeq, err := diagrams.ParseMermaidSequence(mermaidSeq)
	if err == nil {
		fmt.Println(parsedSeq.Render())
	}
	fmt.Println()

	printSection("MERMAID: From .mmd File")
	fmt.Println("To render .mmd files, use:")
	fmt.Println("  go run cmd/mermaid-render/main.go examples/authentication.mmd")
	fmt.Println("  go run cmd/mermaid-render/main.go examples/api-flow.mmd")
	fmt.Println()
	fmt.Println("Example .mmd files in examples/ directory")
	fmt.Println()

	printFooter()
}

func printHeader(title string) {
	width := 80
	padding := (width - len(title) - 2) / 2

	fmt.Println(strings.Repeat("=", width))
	fmt.Printf("%s %s %s\n", strings.Repeat("=", padding), title, strings.Repeat("=", padding))
	fmt.Println(strings.Repeat("=", width))
}

func printSection(title string) {
	fmt.Println(title)
	fmt.Println(strings.Repeat("-", len(title)))
	fmt.Println()
}

func printFooter() {
	fmt.Println(strings.Repeat("=", 80))
	fmt.Println("github.com/orchard9/tui-diagrams - Terminal diagram rendering for Go")
	fmt.Println(strings.Repeat("=", 80))
}
