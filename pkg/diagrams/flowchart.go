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
	if len(f.Nodes) == 0 {
		return ""
	}

	// Build graph structure
	nodeMap := make(map[string]Node)
	outgoing := make(map[string][]Edge)
	incomingCount := make(map[string]int)

	for _, node := range f.Nodes {
		nodeMap[node.ID] = node
		incomingCount[node.ID] = 0
	}

	for _, edge := range f.Edges {
		outgoing[edge.From] = append(outgoing[edge.From], edge)
		incomingCount[edge.To]++
	}

	// Find root nodes (nodes with no incoming edges)
	var roots []string
	for _, node := range f.Nodes {
		if incomingCount[node.ID] == 0 {
			roots = append(roots, node.ID)
		}
	}

	// If no roots found (cycle graph), use first node
	if len(roots) == 0 {
		roots = []string{f.Nodes[0].ID}
	}

	// BFS traversal to determine rendering order
	var renderOrder []string
	visited := make(map[string]bool)
	queue := make([]string, len(roots))
	copy(queue, roots)

	for len(queue) > 0 {
		nodeID := queue[0]
		queue = queue[1:]

		if visited[nodeID] {
			continue
		}
		visited[nodeID] = true
		renderOrder = append(renderOrder, nodeID)

		// Add children to queue
		for _, edge := range outgoing[nodeID] {
			if !visited[edge.To] {
				queue = append(queue, edge.To)
			}
		}
	}

	// Render nodes in order
	var output strings.Builder
	rendered := make(map[string]bool)

	for _, nodeID := range renderOrder {
		if rendered[nodeID] {
			continue
		}
		rendered[nodeID] = true

		node := nodeMap[nodeID]

		// Render the node
		if output.Len() > 0 {
			output.WriteString("\n")
		}
		output.WriteString(renderNode(node))
		output.WriteString("\n")

		// Render ALL outgoing edges from this node
		edges := outgoing[nodeID]
		if len(edges) > 0 {
			for i, edge := range edges {
				targetNode := nodeMap[edge.To]
				output.WriteString(renderVerticalEdgeWithTarget(edge, targetNode))
				if i < len(edges)-1 {
					output.WriteString("\n")
				}
			}
			output.WriteString("\n")
		}
	}

	return strings.TrimSpace(output.String())
}

func (f *Flowchart) renderHorizontal() string {
	if len(f.Nodes) == 0 {
		return ""
	}

	// Build graph structure
	nodeMap := make(map[string]Node)
	outgoing := make(map[string][]Edge)
	incomingCount := make(map[string]int)

	for _, node := range f.Nodes {
		nodeMap[node.ID] = node
		incomingCount[node.ID] = 0
	}

	for _, edge := range f.Edges {
		outgoing[edge.From] = append(outgoing[edge.From], edge)
		incomingCount[edge.To]++
	}

	// Find root nodes (nodes with no incoming edges)
	var roots []string
	for _, node := range f.Nodes {
		if incomingCount[node.ID] == 0 {
			roots = append(roots, node.ID)
		}
	}

	// If no roots found (cycle graph), use first node
	if len(roots) == 0 {
		roots = []string{f.Nodes[0].ID}
	}

	// BFS traversal to determine rendering order
	var renderOrder []string
	visited := make(map[string]bool)
	queue := make([]string, len(roots))
	copy(queue, roots)

	for len(queue) > 0 {
		nodeID := queue[0]
		queue = queue[1:]

		if visited[nodeID] {
			continue
		}
		visited[nodeID] = true
		renderOrder = append(renderOrder, nodeID)

		// Add children to queue
		for _, edge := range outgoing[nodeID] {
			if !visited[edge.To] {
				queue = append(queue, edge.To)
			}
		}
	}

	// For horizontal, render nodes inline with edges between them
	var output strings.Builder
	rendered := make(map[string]bool)

	for i, nodeID := range renderOrder {
		if rendered[nodeID] {
			continue
		}
		rendered[nodeID] = true

		node := nodeMap[nodeID]

		// Add spacing if not first node
		if i > 0 {
			output.WriteString("  ")
		}

		// Render the node inline
		output.WriteString(renderNodeInline(node))

		// If there's exactly one outgoing edge, show it inline
		edges := outgoing[nodeID]
		if len(edges) == 1 {
			output.WriteString(" ")
			output.WriteString(renderHorizontalEdge(edges[0]))
		} else if len(edges) > 1 {
			// Multiple edges - show first inline, others on new lines
			output.WriteString(" ")
			output.WriteString(renderHorizontalEdge(edges[0]))
			for j := 1; j < len(edges); j++ {
				output.WriteString("\n       ")
				output.WriteString(BoxVertical)
				output.WriteString(" ")
				output.WriteString(renderHorizontalEdge(edges[j]))
			}
		}
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

func renderVerticalEdgeWithTarget(edge Edge, targetNode Node) string {
	var b strings.Builder

	if edge.Label != "" {
		// Edge with label - just show label, not target (target node renders separately)
		b.WriteString("    ")
		b.WriteString(BoxVertical)
		b.WriteString(" ")
		b.WriteString(edge.Label)
		b.WriteString("\n")
		b.WriteString("    ")
		b.WriteString(ArrowDown)
	} else {
		// No label - just show arrow
		b.WriteString("    ")
		b.WriteString(ArrowDown)
	}

	return b.String()
}

func renderHorizontalEdge(edge Edge) string {
	if edge.Label != "" {
		return fmt.Sprintf("%s[%s]%s", BoxHorizontal, edge.Label, ArrowRight)
	}
	return fmt.Sprintf("%s%s%s", BoxHorizontal, BoxHorizontal, ArrowRight)
}
