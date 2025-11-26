package diagrams

import (
	"strings"
	"testing"
)

func TestFlowchart_NewFlowchart(t *testing.T) {
	flow := NewFlowchart(TopToBottom)

	if flow == nil {
		t.Fatal("NewFlowchart returned nil")
	}

	if flow.Direction != TopToBottom {
		t.Errorf("Expected TopToBottom, got %v", flow.Direction)
	}

	if len(flow.Nodes) != 0 {
		t.Errorf("Expected 0 nodes, got %d", len(flow.Nodes))
	}

	if len(flow.Edges) != 0 {
		t.Errorf("Expected 0 edges, got %d", len(flow.Edges))
	}
}

func TestFlowchart_AddNode(t *testing.T) {
	flow := NewFlowchart(TopToBottom)

	flow.AddNode("start", "Start", ShapeRounded).
		AddNode("process", "Process Data", ShapeBox).
		AddNode("decision", "Valid?", ShapeDiamond)

	if len(flow.Nodes) != 3 {
		t.Errorf("Expected 3 nodes, got %d", len(flow.Nodes))
	}

	// Verify first node
	if flow.Nodes[0].ID != "start" {
		t.Errorf("Expected ID 'start', got '%s'", flow.Nodes[0].ID)
	}

	if flow.Nodes[0].Label != "Start" {
		t.Errorf("Expected Label 'Start', got '%s'", flow.Nodes[0].Label)
	}

	if flow.Nodes[0].Shape != ShapeRounded {
		t.Errorf("Expected ShapeRounded, got %v", flow.Nodes[0].Shape)
	}
}

func TestFlowchart_AddEdge(t *testing.T) {
	flow := NewFlowchart(TopToBottom)

	flow.AddNode("a", "Node A", ShapeBox).
		AddNode("b", "Node B", ShapeBox).
		AddEdge("a", "b", "connects")

	if len(flow.Edges) != 1 {
		t.Errorf("Expected 1 edge, got %d", len(flow.Edges))
	}

	edge := flow.Edges[0]
	if edge.From != "a" {
		t.Errorf("Expected From 'a', got '%s'", edge.From)
	}

	if edge.To != "b" {
		t.Errorf("Expected To 'b', got '%s'", edge.To)
	}

	if edge.Label != "connects" {
		t.Errorf("Expected Label 'connects', got '%s'", edge.Label)
	}
}

func TestFlowchart_RenderVertical(t *testing.T) {
	flow := NewFlowchart(TopToBottom)

	flow.AddNode("start", "Start", ShapeRounded).
		AddNode("end", "End", ShapeRounded).
		AddEdge("start", "end", "")

	output := flow.Render()

	if output == "" {
		t.Error("Render returned empty string")
	}

	// Should contain box-drawing characters
	if !strings.Contains(output, "─") {
		t.Error("Expected horizontal line characters")
	}

	// Should contain node labels
	if !strings.Contains(output, "Start") {
		t.Error("Expected 'Start' label in output")
	}

	if !strings.Contains(output, "End") {
		t.Error("Expected 'End' label in output")
	}

	// Should contain arrow
	if !strings.Contains(output, ArrowDown) {
		t.Error("Expected down arrow in vertical flowchart")
	}
}

func TestFlowchart_RenderHorizontal(t *testing.T) {
	flow := NewFlowchart(LeftToRight)

	flow.AddNode("a", "A", ShapeBox).
		AddNode("b", "B", ShapeBox).
		AddEdge("a", "b", "")

	output := flow.Render()

	if output == "" {
		t.Error("Render returned empty string")
	}

	// Should contain node labels in inline format
	if !strings.Contains(output, "[A]") {
		t.Error("Expected '[A]' in horizontal flowchart")
	}

	if !strings.Contains(output, "[B]") {
		t.Error("Expected '[B]' in horizontal flowchart")
	}

	// Should contain arrow
	if !strings.Contains(output, ArrowRight) {
		t.Error("Expected right arrow in horizontal flowchart")
	}
}

func TestRenderNode_AllShapes(t *testing.T) {
	tests := []struct {
		shape    NodeShape
		label    string
		expected string
	}{
		{ShapeBox, "Test", "Test"},
		{ShapeRounded, "Test", "Test"},
		{ShapeDiamond, "Test", "Test"},
		{ShapeCircle, "Test", "Test"},
	}

	for _, tt := range tests {
		node := Node{ID: "test", Label: tt.label, Shape: tt.shape}
		output := renderNode(node)

		if output == "" {
			t.Errorf("renderNode returned empty string for shape %v", tt.shape)
		}

		if !strings.Contains(output, tt.expected) {
			t.Errorf("Expected output to contain '%s', got: %s", tt.expected, output)
		}
	}
}

func TestRenderBox(t *testing.T) {
	output := renderBox("Hello")

	// Should have 3 lines (top, content, bottom)
	lines := strings.Split(output, "\n")
	if len(lines) != 3 {
		t.Errorf("Expected 3 lines, got %d", len(lines))
	}

	// Should contain label
	if !strings.Contains(output, "Hello") {
		t.Error("Expected 'Hello' in box")
	}

	// Should contain box corners
	if !strings.Contains(output, BoxTopLeft) {
		t.Error("Expected top-left corner")
	}

	if !strings.Contains(output, BoxTopRight) {
		t.Error("Expected top-right corner")
	}

	if !strings.Contains(output, BoxBottomLeft) {
		t.Error("Expected bottom-left corner")
	}

	if !strings.Contains(output, BoxBottomRight) {
		t.Error("Expected bottom-right corner")
	}
}
