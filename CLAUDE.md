# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Commands

**Build and Run:**
```bash
go build ./cmd/api
go run ./cmd/api/main.go
```

**Development:**
```bash
go mod tidy          # Clean up dependencies
go mod download      # Download dependencies
go fmt ./...         # Format code
go vet ./...         # Static analysis
```

**Testing:**
No test framework is currently configured in this project.

## Architecture

This is a Go REST API service with the following structure:

**Entry Point:**
- `cmd/api/main.go` - Main application entry point, starts HTTP server on localhost:8000

**Core Components:**
- `api/` - API type definitions and error handling utilities
- `internal/handlers/` - HTTP request handlers and routing setup
- `internal/middleware/` - Authorization middleware for protected endpoints
- `internal/tools/` - Database interface with mock implementation

**Key Architecture Patterns:**
- **Interface-based database layer**: `DatabaseInterface` in `internal/tools/database.go` with mock implementation in `mockdb.go`
- **Chi router**: Uses go-chi for HTTP routing with middleware support
- **Authorization via query params + headers**: Username in query string, auth token in Authorization header
- **Mock data**: Currently uses in-memory mock data for users (alex, jason, hari) with predefined auth tokens and coin balances

**API Endpoints:**
- `GET /account/coins?username=<user>` - Get user coin balance (requires Authorization header)

**Dependencies:**
- go-chi/chi - HTTP router
- gorilla/schema - Query parameter decoding  
- sirupsen/logrus - Logging
- avukadin/goapi - Additional API utilities

**Known Issues:**
- `internal/middleware/authorization.go:33` - Type case mismatch: `loginDetails` vs `LoginDetails`
- `internal/tools/mockdb.go:31,35` - Username field incorrectly set to "alex" for jason and hari users