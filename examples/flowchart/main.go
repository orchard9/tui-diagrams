package main

import (
	"fmt"

	"github.com/orchard9/tui-diagrams/pkg/diagrams"
)

func main() {
	// Create a simple decision flowchart
	flow := diagrams.NewFlowchart(diagrams.TopToBottom)

	flow.AddNode("start", "Start", diagrams.ShapeRounded).
		AddNode("input", "Get Input", diagrams.ShapeBox).
		AddNode("validate", "Valid?", diagrams.ShapeDiamond).
		AddNode("process", "Process", diagrams.ShapeBox).
		AddNode("error", "Show Error", diagrams.ShapeBox).
		AddNode("end", "End", diagrams.ShapeRounded)

	flow.AddEdge("start", "input", "").
		AddEdge("input", "validate", "").
		AddEdge("validate", "process", "yes").
		AddEdge("validate", "error", "no").
		AddEdge("process", "end", "").
		AddEdge("error", "input", "retry")

	fmt.Println("Decision Flowchart (Vertical)")
	fmt.Println("=============================")
	fmt.Println(flow.Render())

	fmt.Println()

	// Create a horizontal process flow
	flow2 := diagrams.NewFlowchart(diagrams.LeftToRight)

	flow2.AddNode("req", "Request", diagrams.ShapeRounded).
		AddNode("auth", "Authenticate", diagrams.ShapeBox).
		AddNode("proc", "Process", diagrams.ShapeBox).
		AddNode("resp", "Response", diagrams.ShapeRounded)

	flow2.AddEdge("req", "auth", "").
		AddEdge("auth", "proc", "").
		AddEdge("proc", "resp", "")

	fmt.Println("Process Flow (Horizontal)")
	fmt.Println("=========================")
	fmt.Println(flow2.Render())
}
