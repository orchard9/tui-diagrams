package diagrams

import (
	"strings"
	"testing"
)

func TestSequenceDiagram_NewSequenceDiagram(t *testing.T) {
	seq := NewSequenceDiagram()

	if seq == nil {
		t.Fatal("NewSequenceDiagram returned nil")
	}

	if len(seq.Actors) != 0 {
		t.Errorf("Expected 0 actors, got %d", len(seq.Actors))
	}

	if len(seq.Messages) != 0 {
		t.Errorf("Expected 0 messages, got %d", len(seq.Messages))
	}
}

func TestSequenceDiagram_AddActor(t *testing.T) {
	seq := NewSequenceDiagram()

	seq.AddActor("user", "User").
		AddActor("server", "Server").
		AddActor("db", "Database")

	if len(seq.Actors) != 3 {
		t.Errorf("Expected 3 actors, got %d", len(seq.Actors))
	}

	// Verify first actor
	if seq.Actors[0].ID != "user" {
		t.Errorf("Expected ID 'user', got '%s'", seq.Actors[0].ID)
	}

	if seq.Actors[0].Name != "User" {
		t.Errorf("Expected Name 'User', got '%s'", seq.Actors[0].Name)
	}
}

func TestSequenceDiagram_AddMessage(t *testing.T) {
	seq := NewSequenceDiagram()

	seq.AddActor("a", "Actor A").
		AddActor("b", "Actor B").
		AddMessage("a", "b", "Hello", MessageSync)

	if len(seq.Messages) != 1 {
		t.Errorf("Expected 1 message, got %d", len(seq.Messages))
	}

	msg := seq.Messages[0]
	if msg.From != "a" {
		t.Errorf("Expected From 'a', got '%s'", msg.From)
	}

	if msg.To != "b" {
		t.Errorf("Expected To 'b', got '%s'", msg.To)
	}

	if msg.Label != "Hello" {
		t.Errorf("Expected Label 'Hello', got '%s'", msg.Label)
	}

	if msg.Type != MessageSync {
		t.Errorf("Expected MessageSync, got %v", msg.Type)
	}

	if msg.IsSelf {
		t.Error("Expected IsSelf to be false")
	}
}

func TestSequenceDiagram_AddMessage_SelfCall(t *testing.T) {
	seq := NewSequenceDiagram()

	seq.AddActor("a", "Actor A").
		AddMessage("a", "a", "Process", MessageSync)

	msg := seq.Messages[0]
	if !msg.IsSelf {
		t.Error("Expected IsSelf to be true for self-call")
	}
}

func TestSequenceDiagram_Render(t *testing.T) {
	seq := NewSequenceDiagram()

	seq.AddActor("user", "User").
		AddActor("server", "Server").
		AddMessage("user", "server", "Request", MessageSync).
		AddMessage("server", "user", "Response", MessageReturn)

	output := seq.Render()

	if output == "" {
		t.Error("Render returned empty string")
	}

	// Should contain actor names
	if !strings.Contains(output, "User") {
		t.Error("Expected 'User' in output")
	}

	if !strings.Contains(output, "Server") {
		t.Error("Expected 'Server' in output")
	}

	// Should contain message labels
	if !strings.Contains(output, "Request") {
		t.Error("Expected 'Request' in output")
	}

	if !strings.Contains(output, "Response") {
		t.Error("Expected 'Response' in output")
	}

	// Should contain box-drawing characters
	if !strings.Contains(output, BoxVertical) {
		t.Error("Expected vertical line for lifeline")
	}

	// Should contain arrows
	if !strings.Contains(output, ArrowRight) && !strings.Contains(output, ArrowLeft) {
		t.Error("Expected arrows for messages")
	}
}

func TestSequenceDiagram_Render_Empty(t *testing.T) {
	seq := NewSequenceDiagram()

	output := seq.Render()

	if output != "" {
		t.Error("Expected empty string for diagram with no actors")
	}
}

func TestPadCenter(t *testing.T) {
	tests := []struct {
		input  string
		width  int
		verify func(string) bool
	}{
		{"Hello", 10, func(s string) bool { return len(s) == 10 && strings.Contains(s, "Hello") }},
		{"Test", 8, func(s string) bool { return len(s) == 8 }},
		{"LongString", 5, func(s string) bool { return len(s) == 5 }}, // Truncated
	}

	for _, tt := range tests {
		result := padCenter(tt.input, tt.width)
		if !tt.verify(result) {
			t.Errorf("padCenter(%q, %d) = %q failed verification", tt.input, tt.width, result)
		}
	}
}

func TestMinMax(t *testing.T) {
	if min(5, 10) != 5 {
		t.Error("min(5, 10) should be 5")
	}

	if min(10, 5) != 5 {
		t.Error("min(10, 5) should be 5")
	}

	if max(5, 10) != 10 {
		t.Error("max(5, 10) should be 10")
	}

	if max(10, 5) != 10 {
		t.Error("max(10, 5) should be 10")
	}
}
