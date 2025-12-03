## Build, Lint, and Test Commands

- **Build:** `go build ./...`
- **Test:** `go test ./...`
- **Run a single test:** `go test -run ^TestName$`
- **Lint:** `golangci-lint run`

## Database Operations

- **Create migration:** `make migrate-create name=migration_name`
- **Run migrations:** `make migrate-up`
- **Rollback migration:** `make migrate-down`

## Code Style Guidelines

- **Imports:** Group imports into standard library and third-party library blocks.
- **Formatting:** Use tabs for indentation.
- **Types:** Use structs for handlers.
- **Naming Conventions:**
  - Packages: lowercase (e.g., `memberships`)
  - Structs: PascalCase (e.g., `Handler`)
  - Functions: PascalCase (e.g., `NewHandler`)
  - Variables: camelCase (e.g., `membershipHandler`)
- **Error Handling:** Check errors immediately after a function call. Use `log.Fatal` for fatal errors and `fmt.Println` for non-fatal errors.
