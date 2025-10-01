# Repository Guidelines

## Project Structure & Module Organization
The CLI entry point lives in `main.go`, delegating to packages under `cmd/passh` for command wiring. Reusable logic sits in `pkg`, with `pkg/password` for credential handling, `pkg/login` managing auth state, and `pkg/database` persisting records. Environment helpers (`pkg/env`) load config from `config.ini`, while `pkg/prompt` defines interactive flows. Assets like onboarding copy reside in `assets/`. Tests currently live beside the code in `pkg/password/password_test.go`; place new tests next to the code they cover.

## Build, Test, and Development Commands
`go install github.com/mclacore/passh@latest` installs the CLI locally. Use `go run ./cmd/passh` to iterate without installing. `make test` runs `go test ./... -cover` across modules. Run a focused test with `go test -run TestName ./pkg/<package>`. Format changes with `go fmt ./...` or `gofmt -w <path>`. `./test.sh` mirrors the CI smoke checks.

## Coding Style & Naming Conventions
Follow standard Go idioms: tabs for indentation, PascalCase for exported identifiers, camelCase for locals. Group imports into stdlib, third-party, and internal sections. Keep functions small and push shared helpers into `pkg`. Document non-obvious decisions with short comments. Handle errors explicitly and wrap them with context when returning.

## Testing Guidelines
Add table-driven tests for new behavior using Go's `testing` package. Maintain coverage parity with `make test`; target the existing coverage threshold by exercising success and failure paths. Name tests `Test<Feature>` and keep fixtures under the same package. Update or add mocks when touching database or prompt flows.

## Commit & Pull Request Guidelines
Match the existing history: one-line, present-tense commit messages (`fix timeout with config.ini`). Break work into logical commits and include configuration updates. PRs should summarize impact, list user-facing changes, link issues, and note any config or schema steps. Attach before/after output where behavior changes.

## Security & Configuration Tips
Never commit real secrets. Use `.env` or `config.ini` templates under version control for examples only. Validate that sensitive data stays in the database layer and strip debug logging before merge.
