package main

import (
	"fmt"

	"github.com/orchard9/tui-diagrams/pkg/diagrams"
)

func main() {
	// Horizontal bar chart
	chart := diagrams.NewBarChart("Monthly Sales", diagrams.Horizontal)

	chart.AddBar("January", 45).
		AddBar("February", 67).
		AddBar("March", 89).
		AddBar("April", 72).
		AddBar("May", 95).
		AddBar("June", 83)

	chart.SetWidth(40)

	fmt.Println(chart.Render())

	fmt.Println()

	// Vertical bar chart with colors
	chart2 := diagrams.NewBarChart("Team Performance", diagrams.Vertical)

	chart2.AddBarWithColor("Frontend", 85, "\x1b[32m"). // Green
								AddBarWithColor("Backend", 92, "\x1b[34m"). // Blue
								AddBarWithColor("DevOps", 78, "\x1b[36m").  // Cyan
								AddBarWithColor("QA", 88, "\x1b[35m")       // Magenta

	chart2.SetWidth(8).SetHeight(12)

	fmt.Println(chart2.Render())

	fmt.Println()

	// Simple comparison chart
	chart3 := diagrams.NewBarChart("", diagrams.Horizontal)

	chart3.AddBar("Tasks Complete", 156).
		AddBar("Tasks Pending", 43).
		AddBar("Tasks Blocked", 12)

	chart3.SetWidth(30).SetShowValues(true)

	fmt.Println(chart3.Render())
}
