# Contributing to TUI Diagrams

Thank you for considering contributing to TUI Diagrams! This document provides guidelines and instructions for contributing.

## Code of Conduct

- Be respectful and constructive in all interactions
- Focus on technical merit and project goals
- Welcome newcomers and help them get started

## How to Contribute

### Reporting Bugs

1. Check existing issues to avoid duplicates
2. Provide a clear, descriptive title
3. Include:
   - Go version (`go version`)
   - Operating system
   - Minimal code example that reproduces the issue
   - Expected vs actual behavior
   - Terminal type (if rendering issues)

### Suggesting Features

1. Open an issue with `[Feature Request]` prefix
2. Describe the use case and why it's valuable
3. Provide examples of the proposed API
4. Consider backward compatibility

### Pull Requests

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/my-feature`)
3. Make your changes following the coding standards
4. Add tests for new functionality
5. Ensure all tests pass (`go test ./...`)
6. Update documentation (README, godoc comments)
7. Commit with clear messages (see commit conventions below)
8. Push to your fork and submit a pull request

## Development Setup

### Prerequisites

- Go 1.23 or higher
- Git

### Getting Started

```bash
# Clone your fork
git clone https://github.com/YOUR_USERNAME/tui-diagrams.git
cd tui-diagrams

# Run tests
go test ./pkg/diagrams -v

# Run examples
go run cmd/demo/main.go
go run cmd/mermaid-render/main.go example.md

# Build all commands
go build ./...
```

## Coding Standards

### Go Code Style

- Follow standard Go conventions (`gofmt`, `go vet`)
- Use meaningful variable and function names
- Keep functions focused and under 50 lines where possible
- Add godoc comments for all exported types and functions
- Avoid external dependencies (use Go standard library)

### API Design

- Maintain fluent/builder pattern for consistency
- Return `*Type` for chainable methods
- Provide sensible defaults
- Make zero values useful
- Keep the `Diagram` interface simple

### Testing

- Write tests for all new features
- Aim for >90% test coverage
- Use table-driven tests for multiple cases
- Test edge cases and error conditions
- Add visual rendering tests when applicable

Example test structure:
```go
func TestNewFeature(t *testing.T) {
    tests := []struct {
        name     string
        input    string
        expected string
    }{
        {"basic case", "input", "expected"},
        {"edge case", "", "default"},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := NewFeature(tt.input)
            if result != tt.expected {
                t.Errorf("got %v, want %v", result, tt.expected)
            }
        })
    }
}
```

### Documentation

- Update README.md for user-facing changes
- Add godoc comments for exported symbols
- Include code examples in godoc
- Update CHANGELOG.md following Keep a Changelog format
- Add examples in `examples/` for new features

## Commit Conventions

Use conventional commit messages:

- `feat: add Gantt chart support`
- `fix: correct sequence diagram arrow rendering`
- `docs: update README installation steps`
- `test: add flowchart edge label tests`
- `refactor: simplify node rendering logic`
- `perf: optimize regex matching in parser`
- `chore: update dependencies`

## Testing Checklist

Before submitting a PR, ensure:

- [ ] All tests pass (`go test ./...`)
- [ ] Code is formatted (`gofmt -s -w .`)
- [ ] No `go vet` warnings
- [ ] Godoc comments added for exports
- [ ] Examples added/updated if needed
- [ ] CHANGELOG.md updated
- [ ] README.md updated if API changed

Run all checks:
```bash
# Format code
gofmt -s -w .

# Check for issues
go vet ./...

# Run tests
go test ./pkg/diagrams -v -race -cover

# Build all commands
go build ./...
```

## Project Structure

```
tui-diagrams/
├── pkg/
│   └── diagrams/          # Core library code
│       ├── diagram.go     # Diagram interface
│       ├── flowchart.go   # Flowchart implementation
│       ├── sequence.go    # Sequence diagram implementation
│       ├── barchart.go    # Bar chart implementation
│       ├── mermaid.go     # Mermaid parser
│       └── *_test.go      # Test files
├── cmd/
│   ├── demo/              # Demo showcase
│   └── mermaid-render/    # Mermaid file renderer
├── examples/              # Example programs
│   ├── flowchart/
│   ├── sequence/
│   ├── barchart/
│   └── tui-integration/
└── README.md
```

## Adding a New Diagram Type

1. Create `newtype.go` in `pkg/diagrams/`
2. Implement the `Diagram` interface:
   ```go
   type NewType struct { ... }
   func (n *NewType) Render() string { ... }
   ```
3. Add builder methods (fluent API)
4. Create `newtype_test.go` with comprehensive tests
5. Add example in `examples/newtype/main.go`
6. Update README.md with new type documentation
7. Add to `cmd/demo/main.go` showcase

## Questions?

- Open an issue for questions
- Check existing issues and discussions
- Review the README and examples

## License

By contributing, you agree that your contributions will be licensed under the MIT License.
