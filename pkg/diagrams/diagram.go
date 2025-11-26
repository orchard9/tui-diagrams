// Package diagrams provides terminal-friendly diagram rendering.
// It supports flowcharts, sequence diagrams, and basic charts using Unicode box-drawing characters.
package diagrams

// Diagram represents any renderable diagram
type Diagram interface {
	// Render converts the diagram to a string suitable for terminal display
	Render() string
}

// Direction represents the flow direction for diagrams
type Direction int

const (
	// TopToBottom flows from top to bottom (vertical)
	TopToBottom Direction = iota
	// LeftToRight flows from left to right (horizontal)
	LeftToRight
)

// Box-drawing characters for clean terminal output
const (
	// Single-line box characters
	BoxVertical    = "│"
	BoxHorizontal  = "─"
	BoxTopLeft     = "┌"
	BoxTopRight    = "┐"
	BoxBottomLeft  = "└"
	BoxBottomRight = "┘"
	BoxTeeDown     = "┬"
	BoxTeeUp       = "┴"
	BoxTeeRight    = "├"
	BoxTeeLeft     = "┤"
	BoxCross       = "┼"

	// Double-line box characters (for emphasis)
	BoxVerticalDouble    = "║"
	BoxHorizontalDouble  = "═"
	BoxTopLeftDouble     = "╔"
	BoxTopRightDouble    = "╗"
	BoxBottomLeftDouble  = "╚"
	BoxBottomRightDouble = "╝"

	// Arrows
	ArrowDown  = "↓"
	ArrowUp    = "↑"
	ArrowRight = "→"
	ArrowLeft  = "←"
)
