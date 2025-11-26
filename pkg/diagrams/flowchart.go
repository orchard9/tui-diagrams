package diagrams

import (
	"fmt"
	"strings"
)

// NodeShape defines the visual style of a flowchart node
type NodeShape int

const (
	// ShapeBox is a rectangular box
	ShapeBox NodeShape = iota
	// ShapeRounded is a box with rounded corners
	ShapeRounded
	// ShapeDiamond is a diamond shape (for decisions)
	ShapeDiamond
	// ShapeCircle is a circular shape
	ShapeCircle
)

// Node represents a single node in a flowchart
type Node struct {
	ID    string
	Label string
	Shape NodeShape
}

// Edge represents a connection between two nodes
type Edge struct {
	From  string
	To    string
	Label string // Optional label for the edge
}

// Flowchart represents a flowchart diagram
type Flowchart struct {
	Direction Direction
	Nodes     []Node
	Edges     []Edge
}

// NewFlowchart creates a new flowchart with the given direction
func NewFlowchart(direction Direction) *Flowchart {
	return &Flowchart{
		Direction: direction,
		Nodes:     []Node{},
		Edges:     []Edge{},
	}
}

// AddNode adds a node to the flowchart
func (f *Flowchart) AddNode(id, label string, shape NodeShape) *Flowchart {
	f.Nodes = append(f.Nodes, Node{
		ID:    id,
		Label: label,
		Shape: shape,
	})
	return f
}

// AddEdge adds an edge between two nodes
func (f *Flowchart) AddEdge(from, to, label string) *Flowchart {
	f.Edges = append(f.Edges, Edge{
		From:  from,
		To:    to,
		Label: label,
	})
	return f
}

// Render converts the flowchart to ASCII art
func (f *Flowchart) Render() string {
	if f.Direction == TopToBottom {
		return f.renderVertical()
	}
	return f.renderHorizontal()
}

func (f *Flowchart) renderVertical() string {
	var output strings.Builder

	// Simple vertical rendering
	for i, node := range f.Nodes {
		// Render the node
		output.WriteString(renderNode(node))

		// Render edges from this node
		if i < len(f.Nodes)-1 {
			output.WriteString("\n")
			// Find edge connecting this node to next
			for _, edge := range f.Edges {
				if edge.From == node.ID {
					output.WriteString(renderVerticalEdge(edge))
					break
				}
			}
			output.WriteString("\n")
		}
	}

	return output.String()
}

func (f *Flowchart) renderHorizontal() string {
	var output strings.Builder

	for i, node := range f.Nodes {
		if i > 0 {
			// Find edge connecting previous node to this one
			prevNode := f.Nodes[i-1]
			for _, edge := range f.Edges {
				if edge.From == prevNode.ID && edge.To == node.ID {
					output.WriteString(" ")
					output.WriteString(renderHorizontalEdge(edge))
					output.WriteString(" ")
					break
				}
			}
		}
		output.WriteString(renderNodeInline(node))
	}

	return output.String()
}

func renderNode(node Node) string {
	switch node.Shape {
	case ShapeBox:
		return renderBox(node.Label)
	case ShapeRounded:
		return renderRounded(node.Label)
	case ShapeDiamond:
		return renderDiamond(node.Label)
	case ShapeCircle:
		return renderCircle(node.Label)
	default:
		return renderBox(node.Label)
	}
}

func renderBox(label string) string {
	width := len(label) + 4
	var b strings.Builder

	// Top border
	b.WriteString(BoxTopLeft)
	b.WriteString(strings.Repeat(BoxHorizontal, width-2))
	b.WriteString(BoxTopRight)
	b.WriteString("\n")

	// Content
	b.WriteString(BoxVertical)
	b.WriteString(" ")
	b.WriteString(label)
	b.WriteString(" ")
	b.WriteString(BoxVertical)
	b.WriteString("\n")

	// Bottom border
	b.WriteString(BoxBottomLeft)
	b.WriteString(strings.Repeat(BoxHorizontal, width-2))
	b.WriteString(BoxBottomRight)

	return b.String()
}

func renderRounded(label string) string {
	width := len(label) + 4
	var b strings.Builder

	// Top border
	b.WriteString("╭")
	b.WriteString(strings.Repeat(BoxHorizontal, width-2))
	b.WriteString("╮")
	b.WriteString("\n")

	// Content
	b.WriteString(BoxVertical)
	b.WriteString(" ")
	b.WriteString(label)
	b.WriteString(" ")
	b.WriteString(BoxVertical)
	b.WriteString("\n")

	// Bottom border
	b.WriteString("╰")
	b.WriteString(strings.Repeat(BoxHorizontal, width-2))
	b.WriteString("╯")

	return b.String()
}

func renderDiamond(label string) string {
	width := len(label) + 4
	var b strings.Builder

	// Top point
	padding := width / 2
	b.WriteString(strings.Repeat(" ", padding))
	b.WriteString("◆")
	b.WriteString("\n")

	// Middle with label
	b.WriteString("< ")
	b.WriteString(label)
	b.WriteString(" >")
	b.WriteString("\n")

	// Bottom point
	b.WriteString(strings.Repeat(" ", padding))
	b.WriteString("◆")

	return b.String()
}

func renderCircle(label string) string {
	var b strings.Builder

	// Simple circle representation
	b.WriteString("( ")
	b.WriteString(label)
	b.WriteString(" )")

	return b.String()
}

func renderNodeInline(node Node) string {
	switch node.Shape {
	case ShapeBox:
		return fmt.Sprintf("[%s]", node.Label)
	case ShapeRounded:
		return fmt.Sprintf("(%s)", node.Label)
	case ShapeDiamond:
		return fmt.Sprintf("<%s>", node.Label)
	case ShapeCircle:
		return fmt.Sprintf("((%s))", node.Label)
	default:
		return fmt.Sprintf("[%s]", node.Label)
	}
}

func renderVerticalEdge(edge Edge) string {
	var b strings.Builder

	if edge.Label != "" {
		// Edge with label
		b.WriteString("    ")
		b.WriteString(BoxVertical)
		b.WriteString(" ")
		b.WriteString(edge.Label)
		b.WriteString("\n")
	}

	b.WriteString("    ")
	b.WriteString(ArrowDown)

	return b.String()
}

func renderHorizontalEdge(edge Edge) string {
	if edge.Label != "" {
		return fmt.Sprintf("%s[%s]%s", BoxHorizontal, edge.Label, ArrowRight)
	}
	return fmt.Sprintf("%s%s%s", BoxHorizontal, BoxHorizontal, ArrowRight)
}
