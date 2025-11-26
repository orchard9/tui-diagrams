package diagrams

import (
	"strings"
	"testing"
)

func TestBarChart_NewBarChart(t *testing.T) {
	chart := NewBarChart("Test Chart", Horizontal)

	if chart == nil {
		t.Fatal("NewBarChart returned nil")
	}

	if chart.Title != "Test Chart" {
		t.Errorf("Expected title 'Test Chart', got '%s'", chart.Title)
	}

	if chart.Orientation != Horizontal {
		t.Errorf("Expected Horizontal, got %v", chart.Orientation)
	}

	if len(chart.Bars) != 0 {
		t.Errorf("Expected 0 bars, got %d", len(chart.Bars))
	}

	if chart.Width != 50 {
		t.Errorf("Expected default width 50, got %d", chart.Width)
	}

	if chart.Height != 10 {
		t.Errorf("Expected default height 10, got %d", chart.Height)
	}

	if !chart.ShowValues {
		t.Error("Expected ShowValues to be true by default")
	}
}

func TestBarChart_AddBar(t *testing.T) {
	chart := NewBarChart("", Horizontal)

	chart.AddBar("Item 1", 10.5).
		AddBar("Item 2", 20.0).
		AddBar("Item 3", 15.75)

	if len(chart.Bars) != 3 {
		t.Errorf("Expected 3 bars, got %d", len(chart.Bars))
	}

	// Verify first bar
	if chart.Bars[0].Label != "Item 1" {
		t.Errorf("Expected label 'Item 1', got '%s'", chart.Bars[0].Label)
	}

	if chart.Bars[0].Value != 10.5 {
		t.Errorf("Expected value 10.5, got %f", chart.Bars[0].Value)
	}

	if chart.Bars[0].Color != "" {
		t.Errorf("Expected empty color, got '%s'", chart.Bars[0].Color)
	}
}

func TestBarChart_AddBarWithColor(t *testing.T) {
	chart := NewBarChart("", Horizontal)

	chart.AddBarWithColor("Red", 50, "\x1b[31m")

	if len(chart.Bars) != 1 {
		t.Errorf("Expected 1 bar, got %d", len(chart.Bars))
	}

	if chart.Bars[0].Color != "\x1b[31m" {
		t.Errorf("Expected ANSI red color, got '%s'", chart.Bars[0].Color)
	}
}

func TestBarChart_SetWidth(t *testing.T) {
	chart := NewBarChart("", Horizontal)
	chart.SetWidth(100)

	if chart.Width != 100 {
		t.Errorf("Expected width 100, got %d", chart.Width)
	}
}

func TestBarChart_SetHeight(t *testing.T) {
	chart := NewBarChart("", Vertical)
	chart.SetHeight(20)

	if chart.Height != 20 {
		t.Errorf("Expected height 20, got %d", chart.Height)
	}
}

func TestBarChart_SetShowValues(t *testing.T) {
	chart := NewBarChart("", Horizontal)
	chart.SetShowValues(false)

	if chart.ShowValues {
		t.Error("Expected ShowValues to be false")
	}
}

func TestBarChart_RenderHorizontal(t *testing.T) {
	chart := NewBarChart("Sales", Horizontal)

	chart.AddBar("Jan", 100).
		AddBar("Feb", 150).
		AddBar("Mar", 120)

	chart.SetWidth(30)

	output := chart.Render()

	if output == "" {
		t.Error("Render returned empty string")
	}

	// Should contain title
	if !strings.Contains(output, "Sales") {
		t.Error("Expected title 'Sales' in output")
	}

	// Should contain labels
	if !strings.Contains(output, "Jan") {
		t.Error("Expected 'Jan' label")
	}

	if !strings.Contains(output, "Feb") {
		t.Error("Expected 'Feb' label")
	}

	if !strings.Contains(output, "Mar") {
		t.Error("Expected 'Mar' label")
	}

	// Should contain bar blocks
	if !strings.Contains(output, "█") {
		t.Error("Expected bar blocks (█)")
	}

	// Should contain vertical separator
	if !strings.Contains(output, BoxVertical) {
		t.Error("Expected vertical separator")
	}

	// Should contain values
	if !strings.Contains(output, "100") {
		t.Error("Expected value '100'")
	}
}

func TestBarChart_RenderVertical(t *testing.T) {
	chart := NewBarChart("", Vertical)

	chart.AddBar("A", 50).
		AddBar("B", 75).
		AddBar("C", 25)

	chart.SetHeight(10).SetWidth(5)

	output := chart.Render()

	if output == "" {
		t.Error("Render returned empty string")
	}

	// Should contain labels
	if !strings.Contains(output, "A") {
		t.Error("Expected 'A' label")
	}

	if !strings.Contains(output, "B") {
		t.Error("Expected 'B' label")
	}

	if !strings.Contains(output, "C") {
		t.Error("Expected 'C' label")
	}

	// Should contain bar blocks
	if !strings.Contains(output, "█") {
		t.Error("Expected bar blocks (█)")
	}

	// Should contain baseline
	if !strings.Contains(output, BoxHorizontal) {
		t.Error("Expected horizontal baseline")
	}
}

func TestBarChart_RenderEmpty(t *testing.T) {
	chart := NewBarChart("Empty", Horizontal)

	output := chart.Render()

	if output != "" {
		t.Error("Expected empty string for chart with no bars")
	}
}

func TestBarChart_ColorReset(t *testing.T) {
	chart := NewBarChart("", Horizontal)

	chart.AddBarWithColor("Test", 50, "\x1b[32m")

	output := chart.Render()

	// Should contain color code
	if !strings.Contains(output, "\x1b[32m") {
		t.Error("Expected ANSI color code in output")
	}

	// Should contain reset code
	if !strings.Contains(output, "\x1b[0m") {
		t.Error("Expected ANSI reset code after colored bar")
	}
}

func TestFormatValue(t *testing.T) {
	tests := []struct {
		input    float64
		expected string
	}{
		{100.0, "100"},
		{100.5, "100.5"},
		{99.9, "99.9"},
		{0.0, "0"},
		{12.34, "12.3"}, // Should round to 1 decimal
	}

	for _, tt := range tests {
		result := formatValue(tt.input)
		if result != tt.expected {
			t.Errorf("formatValue(%f) = %s, expected %s", tt.input, result, tt.expected)
		}
	}
}

func TestPadRight(t *testing.T) {
	tests := []struct {
		input    string
		width    int
		expected int // Expected length
	}{
		{"Hello", 10, 10},
		{"Test", 8, 8},
		{"LongString", 5, 10}, // Should not truncate
	}

	for _, tt := range tests {
		result := padRight(tt.input, tt.width)
		if len(result) != tt.expected {
			t.Errorf("padRight(%q, %d) length = %d, expected %d", tt.input, tt.width, len(result), tt.expected)
		}
	}
}
