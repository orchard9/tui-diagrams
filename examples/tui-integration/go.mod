module example.com/tui-integration

go 1.25.4

require (
	github.com/orchard9/tui v0.0.0
	github.com/orchard9/tui-diagrams v0.0.0
)

require (
	github.com/orchard9/get-ansi v0.0.0-20251122230152-9e1539a5130e // indirect
	golang.org/x/sys v0.38.0 // indirect
	golang.org/x/term v0.37.0 // indirect
)

replace github.com/orchard9/tui-diagrams => ../..

replace github.com/orchard9/tui => /Users/jordanwashburn/Workspace/orchard9/tui
