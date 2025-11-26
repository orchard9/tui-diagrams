package diagrams

import (
	"fmt"
	"strings"
)

// Actor represents a participant in a sequence diagram
type Actor struct {
	ID   string
	Name string
}

// MessageType defines the style of message arrow
type MessageType int

const (
	// MessageSync is a solid arrow (synchronous call)
	MessageSync MessageType = iota
	// MessageAsync is a dashed arrow (asynchronous message)
	MessageAsync
	// MessageReturn is a dashed arrow pointing back (return value)
	MessageReturn
)

// Message represents a message between actors
type Message struct {
	From   string
	To     string
	Label  string
	Type   MessageType
	IsSelf bool // Self-call (actor sends message to itself)
}

// SequenceDiagram represents a sequence diagram
type SequenceDiagram struct {
	Actors   []Actor
	Messages []Message
}

// NewSequenceDiagram creates a new sequence diagram
func NewSequenceDiagram() *SequenceDiagram {
	return &SequenceDiagram{
		Actors:   []Actor{},
		Messages: []Message{},
	}
}

// AddActor adds a participant to the diagram
func (s *SequenceDiagram) AddActor(id, name string) *SequenceDiagram {
	s.Actors = append(s.Actors, Actor{
		ID:   id,
		Name: name,
	})
	return s
}

// AddMessage adds a message between actors
func (s *SequenceDiagram) AddMessage(from, to, label string, msgType MessageType) *SequenceDiagram {
	isSelf := from == to
	s.Messages = append(s.Messages, Message{
		From:   from,
		To:     to,
		Label:  label,
		Type:   msgType,
		IsSelf: isSelf,
	})
	return s
}

// Render converts the sequence diagram to ASCII art
func (s *SequenceDiagram) Render() string {
	if len(s.Actors) == 0 {
		return ""
	}

	var output strings.Builder
	actorWidth := 12 // Fixed width for actor names
	spacing := 6     // Space between actors

	// Build actor index map
	actorIndex := make(map[string]int)
	for i, actor := range s.Actors {
		actorIndex[actor.ID] = i
	}

	// Render actor headers
	for i, actor := range s.Actors {
		if i > 0 {
			output.WriteString(strings.Repeat(" ", spacing))
		}
		output.WriteString(padCenter(actor.Name, actorWidth))
	}
	output.WriteString("\n")

	// Render actor boxes
	for i := range s.Actors {
		if i > 0 {
			output.WriteString(strings.Repeat(" ", spacing))
		}
		output.WriteString(BoxTopLeft)
		output.WriteString(strings.Repeat(BoxHorizontal, actorWidth-2))
		output.WriteString(BoxTopRight)
	}
	output.WriteString("\n")

	// Render lifelines and messages
	for _, msg := range s.Messages {
		fromIdx := actorIndex[msg.From]
		toIdx := actorIndex[msg.To]

		// Render lifelines
		for i := range s.Actors {
			if i > 0 {
				output.WriteString(strings.Repeat(" ", spacing))
			}
			center := actorWidth / 2
			output.WriteString(strings.Repeat(" ", center))
			output.WriteString(BoxVertical)
			output.WriteString(strings.Repeat(" ", actorWidth-center-1))
		}
		output.WriteString("\n")

		// Render message line
		if msg.IsSelf {
			renderSelfMessage(&output, fromIdx, actorWidth, spacing, len(s.Actors), msg)
		} else {
			renderMessage(&output, fromIdx, toIdx, actorWidth, spacing, len(s.Actors), msg)
		}
		output.WriteString("\n")
	}

	// Render closing lifelines
	for i := range s.Actors {
		if i > 0 {
			output.WriteString(strings.Repeat(" ", spacing))
		}
		center := actorWidth / 2
		output.WriteString(strings.Repeat(" ", center))
		output.WriteString(BoxVertical)
		output.WriteString(strings.Repeat(" ", actorWidth-center-1))
	}
	output.WriteString("\n")

	// Render actor bottom boxes
	for i := range s.Actors {
		if i > 0 {
			output.WriteString(strings.Repeat(" ", spacing))
		}
		output.WriteString(BoxBottomLeft)
		output.WriteString(strings.Repeat(BoxHorizontal, actorWidth-2))
		output.WriteString(BoxBottomRight)
	}

	return output.String()
}

func renderMessage(output *strings.Builder, fromIdx, toIdx, actorWidth, spacing, numActors int, msg Message) {
	minIdx := min(fromIdx, toIdx)
	maxIdx := max(fromIdx, toIdx)

	arrow := ArrowRight
	lineChar := BoxHorizontal
	if msg.Type == MessageAsync || msg.Type == MessageReturn {
		lineChar = "-"
	}
	if fromIdx > toIdx {
		arrow = ArrowLeft
	}

	// Leading spaces before first actor
	for i := 0; i < minIdx; i++ {
		if i > 0 {
			output.WriteString(strings.Repeat(" ", spacing))
		}
		center := actorWidth / 2
		output.WriteString(strings.Repeat(" ", center))
		output.WriteString(BoxVertical)
		output.WriteString(strings.Repeat(" ", actorWidth-center-1))
	}

	// From actor to message start
	if minIdx > 0 {
		output.WriteString(strings.Repeat(" ", spacing))
	}
	center := actorWidth / 2
	output.WriteString(strings.Repeat(" ", center))

	// Message line
	labelLen := len(msg.Label)
	totalWidth := (maxIdx-minIdx)*(actorWidth+spacing) - spacing
	lineLen := totalWidth - labelLen - 2

	// Ensure lineLen is non-negative (handle long labels gracefully)
	if lineLen < 0 {
		lineLen = 0
	}

	if fromIdx < toIdx {
		// Left to right
		output.WriteString(strings.Repeat(lineChar, lineLen/2))
		if lineLen > 0 {
			output.WriteString(" ")
		}
		output.WriteString(msg.Label)
		if lineLen > 0 {
			output.WriteString(" ")
		}
		output.WriteString(strings.Repeat(lineChar, lineLen/2))
		output.WriteString(arrow)
	} else {
		// Right to left
		output.WriteString(arrow)
		output.WriteString(strings.Repeat(lineChar, lineLen/2))
		if lineLen > 0 {
			output.WriteString(" ")
		}
		output.WriteString(msg.Label)
		if lineLen > 0 {
			output.WriteString(" ")
		}
		output.WriteString(strings.Repeat(lineChar, lineLen/2))
	}

	// Trailing spaces after last actor
	for i := maxIdx + 1; i < numActors; i++ {
		output.WriteString(strings.Repeat(" ", spacing))
		center := actorWidth / 2
		output.WriteString(strings.Repeat(" ", center))
		output.WriteString(BoxVertical)
		output.WriteString(strings.Repeat(" ", actorWidth-center-1))
	}
}

func renderSelfMessage(output *strings.Builder, actorIdx, actorWidth, spacing, numActors int, msg Message) {
	// Leading actors
	for i := 0; i < actorIdx; i++ {
		if i > 0 {
			output.WriteString(strings.Repeat(" ", spacing))
		}
		center := actorWidth / 2
		output.WriteString(strings.Repeat(" ", center))
		output.WriteString(BoxVertical)
		output.WriteString(strings.Repeat(" ", actorWidth-center-1))
	}

	// Self-message actor
	if actorIdx > 0 {
		output.WriteString(strings.Repeat(" ", spacing))
	}
	center := actorWidth / 2
	output.WriteString(strings.Repeat(" ", center))
	output.WriteString(BoxVertical)
	output.WriteString(ArrowRight)
	output.WriteString(fmt.Sprintf("[%s]", msg.Label))

	// Trailing actors (if any)
	for i := actorIdx + 1; i < numActors; i++ {
		output.WriteString(strings.Repeat(" ", spacing))
		center := actorWidth / 2
		output.WriteString(strings.Repeat(" ", center))
		output.WriteString(BoxVertical)
		output.WriteString(strings.Repeat(" ", actorWidth-center-1))
	}
}

func padCenter(s string, width int) string {
	if len(s) >= width {
		return s[:width]
	}
	leftPad := (width - len(s)) / 2
	rightPad := width - len(s) - leftPad
	return strings.Repeat(" ", leftPad) + s + strings.Repeat(" ", rightPad)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
