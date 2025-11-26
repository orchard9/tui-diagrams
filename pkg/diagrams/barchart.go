package diagrams

import (
	"fmt"
	"math"
	"strings"
)

// BarOrientation defines whether bars are horizontal or vertical
type BarOrientation int

const (
	// Horizontal bars extend left to right
	Horizontal BarOrientation = iota
	// Vertical bars extend bottom to top
	Vertical
)

// Bar represents a single bar in a chart
type Bar struct {
	Label string
	Value float64
	Color string // ANSI color code (optional)
}

// BarChart represents a bar chart
type BarChart struct {
	Title       string
	Bars        []Bar
	Orientation BarOrientation
	Width       int // Chart width (for horizontal) or bar width (for vertical)
	Height      int // Chart height (for vertical) or bar height (for horizontal)
	ShowValues  bool
}

// NewBarChart creates a new bar chart
func NewBarChart(title string, orientation BarOrientation) *BarChart {
	return &BarChart{
		Title:       title,
		Bars:        []Bar{},
		Orientation: orientation,
		Width:       50,
		Height:      10,
		ShowValues:  true,
	}
}

// AddBar adds a bar to the chart
func (b *BarChart) AddBar(label string, value float64) *BarChart {
	b.Bars = append(b.Bars, Bar{
		Label: label,
		Value: value,
	})
	return b
}

// AddBarWithColor adds a bar with ANSI color
func (b *BarChart) AddBarWithColor(label string, value float64, color string) *BarChart {
	b.Bars = append(b.Bars, Bar{
		Label: label,
		Value: value,
		Color: color,
	})
	return b
}

// SetWidth sets the chart width
func (b *BarChart) SetWidth(width int) *BarChart {
	b.Width = width
	return b
}

// SetHeight sets the chart height
func (b *BarChart) SetHeight(height int) *BarChart {
	b.Height = height
	return b
}

// SetShowValues toggles value display
func (b *BarChart) SetShowValues(show bool) *BarChart {
	b.ShowValues = show
	return b
}

// Render converts the bar chart to ASCII art
func (b *BarChart) Render() string {
	if len(b.Bars) == 0 {
		return ""
	}

	if b.Orientation == Horizontal {
		return b.renderHorizontal()
	}
	return b.renderVertical()
}

func (b *BarChart) renderHorizontal() string {
	var output strings.Builder

	// Title
	if b.Title != "" {
		output.WriteString(b.Title)
		output.WriteString("\n")
		output.WriteString(strings.Repeat("=", len(b.Title)))
		output.WriteString("\n\n")
	}

	// Find max value for scaling
	maxValue := 0.0
	for _, bar := range b.Bars {
		if bar.Value > maxValue {
			maxValue = bar.Value
		}
	}

	// Find max label width
	maxLabelWidth := 0
	for _, bar := range b.Bars {
		if len(bar.Label) > maxLabelWidth {
			maxLabelWidth = len(bar.Label)
		}
	}

	// Render each bar
	for _, bar := range b.Bars {
		// Label (right-padded)
		output.WriteString(padRight(bar.Label, maxLabelWidth))
		output.WriteString(" ")
		output.WriteString(BoxVertical)
		output.WriteString(" ")

		// Bar
		barLength := int((bar.Value / maxValue) * float64(b.Width))
		if bar.Color != "" {
			output.WriteString(bar.Color)
		}
		output.WriteString(strings.Repeat("█", barLength))
		if bar.Color != "" {
			output.WriteString("\x1b[0m") // Reset color
		}

		// Value
		if b.ShowValues {
			output.WriteString(" ")
			output.WriteString(formatValue(bar.Value))
		}

		output.WriteString("\n")
	}

	return output.String()
}

func (b *BarChart) renderVertical() string {
	var output strings.Builder

	// Title
	if b.Title != "" {
		output.WriteString(b.Title)
		output.WriteString("\n")
		output.WriteString(strings.Repeat("=", len(b.Title)))
		output.WriteString("\n\n")
	}

	// Find max value for scaling
	maxValue := 0.0
	for _, bar := range b.Bars {
		if bar.Value > maxValue {
			maxValue = bar.Value
		}
	}

	barWidth := b.Width
	if barWidth < 3 {
		barWidth = 3
	}

	// Render from top to bottom
	for row := b.Height; row >= 0; row-- {
		threshold := (float64(row) / float64(b.Height)) * maxValue

		for i, bar := range b.Bars {
			if i > 0 {
				output.WriteString("  ")
			}

			if bar.Value >= threshold {
				// Draw bar segment
				if bar.Color != "" {
					output.WriteString(bar.Color)
				}
				output.WriteString(strings.Repeat("█", barWidth))
				if bar.Color != "" {
					output.WriteString("\x1b[0m")
				}
			} else {
				// Empty space
				output.WriteString(strings.Repeat(" ", barWidth))
			}
		}
		output.WriteString("\n")
	}

	// Baseline
	totalWidth := len(b.Bars)*barWidth + (len(b.Bars)-1)*2
	output.WriteString(strings.Repeat(BoxHorizontal, totalWidth))
	output.WriteString("\n")

	// Labels
	for i, bar := range b.Bars {
		if i > 0 {
			output.WriteString("  ")
		}
		label := bar.Label
		if len(label) > barWidth {
			label = label[:barWidth]
		}
		output.WriteString(padCenter(label, barWidth))
	}
	output.WriteString("\n")

	// Values (if enabled)
	if b.ShowValues {
		for i, bar := range b.Bars {
			if i > 0 {
				output.WriteString("  ")
			}
			value := formatValue(bar.Value)
			output.WriteString(padCenter(value, barWidth))
		}
	}

	return output.String()
}

func padRight(s string, width int) string {
	if len(s) >= width {
		return s
	}
	return s + strings.Repeat(" ", width-len(s))
}

func formatValue(value float64) string {
	if value == math.Floor(value) {
		return fmt.Sprintf("%.0f", value)
	}
	return fmt.Sprintf("%.1f", value)
}
