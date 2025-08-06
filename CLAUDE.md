# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Goffeine is a caffeine tracking application built with Go and Templ for the backend/templating, TailwindCSS for styling, and uses OpenAI API for natural language processing of caffeine intake. The application tracks caffeine consumption with a 5-hour half-life model and ignores intake older than 24 hours.

## Architecture

**Core Components:**
- `cmd/main.go` - Application entry point, sets up dependencies and starts HTTP server
- `internal/server/server.go` - HTTP server setup with routing (`/api/status`, `/api/add`, `/`)
- `internal/handler/` - HTTP handlers for status, intake, and page rendering
- `internal/tracker/tracker.go` - Core caffeine tracking logic with exponential decay calculations
- `internal/ask/client.go` - OpenAI client for parsing natural language caffeine intake
- `internal/repl/` - CLI interface with commands (help, status, add)

**Data Flow:**
1. User input → OpenAI API (via `ask.Client`) → structured caffeine data
2. Structured data → `tracker.Tracker` → in-memory repository
3. Caffeine level calculations use exponential decay formula based on 5-hour half-life

**Templates:** Uses [Templ](https://github.com/a-h/templ) for type-safe HTML templating in `internal/handler/*.templ`

## Development Commands

**Setup:**
```bash
make setup  # Install templ tool
```

**Build:**
```bash
make build  # Generate templates, build CSS, compile Go binary
```

**Run:**
```bash
make run    # Build and run the application
```

**Test & Quality:**
```bash
make test   # Run tests, formatting, staticcheck, and gosec
go test -v ./...           # Run tests only
go test -v ./internal/tracker  # Run specific package tests
```

**CSS Development:**
```bash
pnpm run watch  # Watch and rebuild TailwindCSS
pnpm run build  # Build minified CSS
```

## Environment Setup

- Requires Go 1.23+, Node.js LTS, and pnpm
- Uses `mise.toml` for tool version management
- Requires `OPENAI_API_KEY` environment variable
- Optional `PORT` environment variable (defaults to 8080)
- Uses `.env` file for local development (not committed)

## Testing Strategy

- Unit tests for core logic in `internal/tracker/` and `internal/server/`
- Memory repository implementation for testing (`internal/tracker/memrepository.go`)
- CI pipeline runs: build, tests, go fmt, staticcheck, gosec

## Key Dependencies

- `github.com/a-h/templ` - Type-safe templating
- `github.com/sashabaranov/go-openai` - OpenAI API client  
- `github.com/alexedwards/scs/v2` - Session management
- TailwindCSS 4.x for styling

## Deployment

- Docker-based deployment to Google Cloud Platform
- Multi-stage build: Go builder → minimal Debian runtime
- Uses Google Cloud Build for CI/CD pipeline