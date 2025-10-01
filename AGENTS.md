# Agents Guide for passh Repository

## Build, Lint, and Test Commands

- Build and install: `go install github.com/mclacore/passh@latest`
- Run all tests with coverage: `make test` (runs `go test ./... -cover`)
- Run a single test: `go test -run <TestName> <package_path>` (e.g., `go test -run TestSimplePassword ./pkg/password`)
- Format code: `go fmt ./...` or `gofmt -w .`

## Code Style Guidelines

- Follow standard Go formatting and style conventions.
- Use `go fmt` or `gofmt` to format code before committing.
- Imports should be grouped in three sections: standard library, third-party, and local packages.
- Use camelCase for variable and function names.
- Use PascalCase for exported types, functions, and constants.
- Handle errors explicitly; return errors up the call stack with descriptive messages.
- Keep functions small, focused, and testable.
- Use Go's built-in `testing` package for unit tests.
- Write clear and concise comments where necessary.

## Additional Notes

- No Cursor or Copilot rules detected in the repository.

This guide is intended for agentic coding agents operating in this repository to maintain consistency and quality.
