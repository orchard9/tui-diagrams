# Integrating tui-diagrams with TUI Framework

This guide explains how to use `tui-diagrams` with the `github.com/orchard9/tui` framework.

## Installation

```bash
go get github.com/orchard9/tui
go get github.com/orchard9/tui-diagrams
```

## Basic Integration

### Example 1: Static Diagram in TUI

```go
package main

import (
	"log"

	"github.com/orchard9/tui/pkg/tui"
	"github.com/orchard9/tui-diagrams/pkg/diagrams"
)

type model struct {
	diagram diagrams.Diagram
}

func (m model) Init() tui.Cmd {
	// Create a flowchart
	flow := diagrams.NewFlowchart(diagrams.TopToBottom)
	flow.AddNode("start", "Start", diagrams.ShapeRounded).
		AddNode("process", "Process", diagrams.ShapeBox).
		AddNode("end", "End", diagrams.ShapeRounded)

	flow.AddEdge("start", "process", "").
		AddEdge("process", "end", "")

	m.diagram = flow
	return nil
}

func (m model) Update(msg tui.Msg) (tui.Model, tui.Cmd) {
	switch msg := msg.(type) {
	case tui.KeyMsg:
		if msg.String() == "q" {
			return m, tui.Quit
		}
	}
	return m, nil
}

func (m model) View() string {
	header := "Flowchart Demo - Press 'q' to quit\n\n"
	return header + m.diagram.Render()
}

func main() {
	p := tui.NewProgram(model{})
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
```

### Example 2: Interactive Diagram Switcher

```go
package main

import (
	"log"

	"github.com/orchard9/tui/pkg/tui"
	"github.com/orchard9/tui-diagrams/pkg/diagrams"
)

type diagramType int

const (
	flowchartView diagramType = iota
	sequenceView
	chartView
)

type model struct {
	currentView diagramType
	flowchart   *diagrams.Flowchart
	sequence    *diagrams.SequenceDiagram
	chart       *diagrams.BarChart
}

func initialModel() model {
	// Create flowchart
	flow := diagrams.NewFlowchart(diagrams.TopToBottom)
	flow.AddNode("A", "Start", diagrams.ShapeRounded).
		AddNode("B", "Process", diagrams.ShapeBox).
		AddNode("C", "End", diagrams.ShapeRounded)
	flow.AddEdge("A", "B", "").AddEdge("B", "C", "")

	// Create sequence diagram
	seq := diagrams.NewSequenceDiagram()
	seq.AddActor("user", "User").
		AddActor("api", "API").
		AddMessage("user", "api", "Request", diagrams.MessageSync).
		AddMessage("api", "user", "Response", diagrams.MessageReturn)

	// Create bar chart
	chart := diagrams.NewBarChart("Stats", diagrams.Horizontal)
	chart.AddBar("Tasks Done", 45).
		AddBar("Tasks Pending", 23).
		SetWidth(30)

	return model{
		currentView: flowchartView,
		flowchart:   flow,
		sequence:    seq,
		chart:       chart,
	}
}

func (m model) Init() tui.Cmd {
	return nil
}

func (m model) Update(msg tui.Msg) (tui.Model, tui.Cmd) {
	switch msg := msg.(type) {
	case tui.KeyMsg:
		switch msg.String() {
		case "q":
			return m, tui.Quit
		case "1":
			m.currentView = flowchartView
		case "2":
			m.currentView = sequenceView
		case "3":
			m.currentView = chartView
		}
	}
	return m, nil
}

func (m model) View() string {
	header := `
╔════════════════════════════════════════════════════════════╗
║  Diagram Viewer - Press 1/2/3 to switch, q to quit        ║
╠════════════════════════════════════════════════════════════╣
║  1: Flowchart  |  2: Sequence  |  3: Bar Chart            ║
╚════════════════════════════════════════════════════════════╝

`

	var diagram string
	switch m.currentView {
	case flowchartView:
		diagram = m.flowchart.Render()
	case sequenceView:
		diagram = m.sequence.Render()
	case chartView:
		diagram = m.chart.Render()
	}

	return header + diagram
}

func main() {
	p := tui.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
```

### Example 3: Dynamic Data Visualization

```go
package main

import (
	"log"
	"time"

	"github.com/orchard9/tui/pkg/tui"
	"github.com/orchard9/tui-diagrams/pkg/diagrams"
)

type tickMsg time.Time

type model struct {
	count int
	chart *diagrams.BarChart
}

func (m model) Init() tui.Cmd {
	return tui.Tick(1 * time.Second)
}

func (m model) Update(msg tui.Msg) (tui.Model, tui.Cmd) {
	switch msg := msg.(type) {
	case tui.KeyMsg:
		if msg.String() == "q" {
			return m, tui.Quit
		}
	case tui.TickMsg:
		m.count++

		// Update chart with new data
		m.chart = diagrams.NewBarChart("Real-time Stats", diagrams.Horizontal)
		m.chart.AddBarWithColor("Active", m.count*3, "\x1b[32m").
			AddBarWithColor("Pending", m.count*2, "\x1b[33m").
			AddBarWithColor("Failed", m.count, "\x1b[31m").
			SetWidth(40)

		return m, tui.Tick(1 * time.Second)
	}
	return m, nil
}

func (m model) View() string {
	header := "Live Data Dashboard - Press 'q' to quit\n\n"
	if m.chart == nil {
		return header + "Loading..."
	}
	return header + m.chart.Render()
}

func main() {
	p := tui.NewProgram(model{})
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
```

### Example 4: Mermaid File Viewer

```go
package main

import (
	"log"
	"os"

	"github.com/orchard9/tui/pkg/tui"
	"github.com/orchard9/tui-diagrams/pkg/diagrams"
)

type model struct {
	diagram diagrams.Diagram
	error   error
}

func (m model) Init() tui.Cmd {
	return nil
}

func (m model) Update(msg tui.Msg) (tui.Model, tui.Cmd) {
	switch msg := msg.(type) {
	case tui.KeyMsg:
		if msg.String() == "q" {
			return m, tui.Quit
		}
	}
	return m, nil
}

func (m model) View() string {
	if m.error != nil {
		return "Error: " + m.error.Error() + "\n\nPress 'q' to quit"
	}

	if m.diagram == nil {
		return "No diagram loaded\n\nPress 'q' to quit"
	}

	header := "Mermaid Diagram Viewer - Press 'q' to quit\n\n"
	return header + m.diagram.Render()
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: go run main.go <file.mmd>")
	}

	// Load .mmd file
	diagram, err := diagrams.ParseMmdFile(os.Args[1])

	p := tui.NewProgram(model{
		diagram: diagram,
		error:   err,
	})

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
```

## API Reference

### Diagram Interface

All diagrams implement the `diagrams.Diagram` interface:

```go
type Diagram interface {
    Render() string
}
```

This makes it easy to swap diagrams in your TUI views.

### Creating Diagrams

**Flowchart:**
```go
flow := diagrams.NewFlowchart(diagrams.TopToBottom)
flow.AddNode(id, label, shape)
flow.AddEdge(from, to, label)
```

**Sequence Diagram:**
```go
seq := diagrams.NewSequenceDiagram()
seq.AddActor(id, name)
seq.AddMessage(from, to, label, msgType)
```

**Bar Chart:**
```go
chart := diagrams.NewBarChart(title, orientation)
chart.AddBar(label, value)
chart.AddBarWithColor(label, value, ansiColor)
```

**From Mermaid:**
```go
diagram, err := diagrams.ParseMmdFile("diagram.mmd")
diagrams, err := diagrams.ParseMermaidFromFile("README.md")
flow, err := diagrams.ParseMermaidFlowchart(mermaidText)
seq, err := diagrams.ParseMermaidSequence(mermaidText)
```

## Best Practices

1. **Static Diagrams**: Create once in `Init()`, store in model
2. **Dynamic Diagrams**: Recreate in `Update()` when data changes
3. **Large Diagrams**: Consider pagination or scrolling
4. **Colors**: Use ANSI colors for better visual hierarchy
5. **Error Handling**: Always handle parse errors from Mermaid

## Performance

- Diagram rendering is fast (<1ms for typical diagrams)
- Re-rendering on every frame is acceptable
- For very large diagrams (>100 nodes), cache the rendered string

## Examples

See `examples/tui-integration/` for complete working examples.

## Troubleshooting

**Import issues:**
```bash
go mod tidy
```

**Rendering issues:**
- Ensure terminal supports Unicode box-drawing characters
- Ensure terminal supports ANSI colors (for colored charts)

## Resources

- TUI Framework: https://github.com/orchard9/tui
- TUI Diagrams: https://github.com/orchard9/tui-diagrams
- TUI Styles: https://github.com/orchard9/tui-styles
