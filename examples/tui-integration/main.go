package main

import (
	"log"

	"github.com/orchard9/tui-diagrams/pkg/diagrams"
	"github.com/orchard9/tui/pkg/tui"
)

type diagramType int

const (
	flowchartDiagram diagramType = iota
	sequenceDiagram
	barchartDiagram
)

type model struct {
	currentDiagram diagramType
}

func (m model) Init() tui.Cmd {
	return nil
}

func (m model) Update(msg tui.Msg) (tui.Model, tui.Cmd) {
	switch msg := msg.(type) {
	case tui.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tui.Quit
		case "1":
			m.currentDiagram = flowchartDiagram
		case "2":
			m.currentDiagram = sequenceDiagram
		case "3":
			m.currentDiagram = barchartDiagram
		}
	}
	return m, nil
}

func (m model) View() string {
	header := `
╔═══════════════════════════════════════════════════════════╗
║       TUI Diagrams - Interactive Diagram Viewer           ║
╠═══════════════════════════════════════════════════════════╣
║  Press 1: Flowchart  |  2: Sequence  |  3: Bar Chart     ║
║  Press q: Quit                                            ║
╚═══════════════════════════════════════════════════════════╝

`

	var diagram string
	switch m.currentDiagram {
	case flowchartDiagram:
		diagram = renderFlowchart()
	case sequenceDiagram:
		diagram = renderSequence()
	case barchartDiagram:
		diagram = renderBarchart()
	}

	return header + diagram
}

func renderFlowchart() string {
	flow := diagrams.NewFlowchart(diagrams.TopToBottom)

	flow.AddNode("start", "Start", diagrams.ShapeRounded).
		AddNode("check", "Has Data?", diagrams.ShapeDiamond).
		AddNode("process", "Process", diagrams.ShapeBox).
		AddNode("save", "Save Result", diagrams.ShapeBox).
		AddNode("end", "End", diagrams.ShapeRounded)

	flow.AddEdge("start", "check", "").
		AddEdge("check", "process", "yes").
		AddEdge("check", "end", "no").
		AddEdge("process", "save", "").
		AddEdge("save", "end", "")

	return "\nFlowchart: Data Processing Pipeline\n" +
		"====================================\n\n" +
		flow.Render()
}

func renderSequence() string {
	seq := diagrams.NewSequenceDiagram()

	seq.AddActor("user", "User").
		AddActor("app", "App").
		AddActor("api", "API")

	seq.AddMessage("user", "app", "Click Button", diagrams.MessageSync).
		AddMessage("app", "api", "GET /data", diagrams.MessageSync).
		AddMessage("api", "app", "JSON", diagrams.MessageReturn).
		AddMessage("app", "app", "Parse", diagrams.MessageSync).
		AddMessage("app", "user", "Display", diagrams.MessageReturn)

	return "\nSequence: API Data Fetch\n" +
		"========================\n\n" +
		seq.Render()
}

func renderBarchart() string {
	chart := diagrams.NewBarChart("Project Progress", diagrams.Horizontal)

	chart.AddBarWithColor("Planning", 100, "\x1b[32m").
		AddBarWithColor("Development", 75, "\x1b[33m").
		AddBarWithColor("Testing", 45, "\x1b[36m").
		AddBarWithColor("Deployment", 20, "\x1b[31m")

	chart.SetWidth(35)

	return "\n" + chart.Render()
}

func main() {
	p := tui.NewProgram(model{currentDiagram: flowchartDiagram})
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
