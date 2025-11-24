# GitHub Copilot Guidelines for go-utils

This document outlines the rules, guidelines, and best practices for developing and maintaining the `go-utils` repository with assistance from GitHub Copilot.

---

## Table of Contents

1. [Code Quality Standards](#code-quality-standards)
2. [Project Structure](#project-structure)
3. [Documentation Requirements](#documentation-requirements)
4. [Version Management](#version-management)
5. [Release Process](#release-process)
6. [Commit Conventions](#commit-conventions)
7. [Testing Guidelines](#testing-guidelines)
8. [Linting and Formatting](#linting-and-formatting)

---

## Code Quality Standards

### General Rules

- **Minimal Dependencies**: Use Go standard library whenever possible
- **Clear Separation of Concerns**: Each utility should be isolated and reusable
- **Modular Architecture**: CLI tools should be thin layers over core packages
- **Predictable Behavior**: Follow sensible defaults with clear error messages
- **Error Handling**: Always return meaningful errors with context
- **Code Style**: Follow standard Go conventions (gofmt, goimports)

### Go-Specific Standards

- Use `fmt.Errorf("context: %w", err)` for error wrapping
- Defer cleanup operations (file closes, resource releases)
- Use interfaces for extensibility where appropriate
- Avoid global state and mutable exports
- Keep functions focused and single-purpose

### Naming Conventions

- Package names: lowercase, single word (`envutils`, `formatter`)
- Public functions: PascalCase (`FormatJSON`, `CopyTemplate`)
- Private functions: camelCase (`formatJSON`, `copyTemplate`)
- Constants: ALL_CAPS for exported, camelCase for unexported

---

## Project Structure

### Folder Organization

```
go-utils/
├── cmd/                # CLI applications (one per folder)
│   ├── create-env/
│   ├── setup-env/
│   └── jsonfmt/
│
├── pkg/                # Reusable packages (one per folder)
│   ├── envutils/
│   └── formatter/
│
├── .github/            # GitHub-specific files
│   └── copilot.md      # This file
│
├── dist/               # Built binaries (in .gitignore)
├── .gitignore
├── go.mod
├── LICENSE
└── README.md
```

### Adding New Utilities

When adding a new utility:

1. Create a folder in `cmd/<utility-name>/`
2. Implement `main.go` as a thin CLI wrapper
3. Create `cmd/<utility-name>/README.md` with usage examples
4. Create corresponding package in `pkg/<package-name>/` with core logic
5. Export public functions with proper documentation
6. Update root README with new utility and package documentation

---

## Documentation Requirements

### Markdown Linting Rules

- **MD031**: Fenced code blocks must be surrounded by blank lines
- **MD040**: Fenced code blocks must have a language specified
- **MD001**: Heading hierarchy must be correct (no skipping levels)
- **MD003**: Consistent heading style (use `#` for all headings)
- **MD013**: Line length (prefer < 120 characters for readability)

### Root README Structure

Must include:

- Project description and vision
- Current project structure (updated with each utility addition)
- Included utilities section (one subsection per utility)
- Reusable packages section (one subsection per package)
- Usage modes (as CLI tools and as Go packages)
- Versioning strategy
- Release process
- Design principles
- Future direction

### Per-Utility README Requirements

Must include:

- Tool description
- Usage section with examples
- Features list
- Installation instructions
- Error handling examples
- Implementation details (link to package)

### Code Documentation

- All exported functions must have doc comments
- Doc comments start with the function name
- Include parameter and return value descriptions
- Example: `// FormatJSON formats JSON input with 2-space indentation`

---

## Version Management

### Semantic Versioning

Follow semantic versioning: `MAJOR.MINOR.PATCH`

- **MAJOR**: Breaking changes to API or tool behavior
- **MINOR**: New features or utilities added
- **PATCH**: Bug fixes, documentation updates, or improvements

### Version Tag Operations

When updating the version tag:

1. **Update Root README**:
   - Update project structure diagram if new utilities added
   - Update included utilities section
   - Update reusable packages section
   - Verify all examples work correctly

2. **For Each New Utility**:
   - Create utility-specific README in `cmd/<utility>/README.md`
   - Add documentation to root README
   - Include usage examples
   - Document error handling

3. **Provide Release Notes**:
   - Format release notes in Markdown
   - Include: What's New, Features, Improvements, Distribution info
   - Post in chat/documentation
   - Create annotated git tag with message

### Tag Format

```bash
git tag -a v1.0.0 -m "Release v1.0.0: [Brief description]"
```

---

## Release Process

### Step-by-Step Release Workflow

1. **Commit Changes**:
   - Use conventional commit format
   - Ensure all code is tested and documented

2. **Update Documentation**:
   - Update root README with current structure
   - Add utility README if new tool added
   - Verify markdown linting passes

3. **Create Version Tag**:
   - Delete old tag if updating: `git tag -d v1.0.x`
   - Create new tag: `git tag -a v1.0.x -m "Release message"`

4. **Generate Release Notes**:
   - Document changes in Markdown format
   - Include features, improvements, and distribution info
   - Share in chat/documentation

5. **Push to Remote**:
   - Push main branch: `git push origin main`
   - Push tags: `git push origin v1.0.x`

### Built Artifacts

- Binaries are built locally and added to `dist/` folder
- `dist/` folder is in `.gitignore` and not tracked
- Pre-built binaries can be provided in GitHub Releases separately

---

## Commit Conventions

### Commit Message Format

Follow Conventional Commits:

```
<type>: <subject>

<body>

<footer>
```

### Types

- `feat`: New feature or utility
- `fix`: Bug fix
- `docs`: Documentation changes
- `style`: Code style changes (formatting, missing semicolons, etc.)
- `refactor`: Code refactoring without feature changes
- `perf`: Performance improvements
- `test`: Adding or updating tests
- `build`: Changes to build system or dependencies
- `chore`: Maintenance tasks

### Examples

```
feat: Add jsonfmt CLI tool for JSON formatting

- Add jsonfmt command-line tool
- Add formatter package with JSON processing utilities
- Update README with jsonfmt documentation

fix: Handle empty JSON files in formatter

- Add check for empty input
- Return proper error message

docs: Update root README with current structure

- Add formatter package documentation
- Include jsonfmt in usage examples
```

---

## Testing Guidelines

### Unit Tests

- Create `*_test.go` files alongside implementation
- Use table-driven tests for multiple cases
- Test both success and error paths
- Aim for > 80% code coverage

### Test Naming

- Test functions: `Test<FunctionName>`
- Example: `TestFormatJSON`, `TestCopyTemplate`

### Error Cases

Always test:

- Invalid input
- Missing files
- Permission errors
- Malformed data

### Running Tests

```bash
go test ./...                    # Run all tests
go test -v ./...                # Verbose output
go test -cover ./...            # Coverage report
go test -coverprofile=cov.out ./...  # Detailed coverage
```

---

## Linting and Formatting

### Code Formatting

```bash
gofmt -w .          # Format all Go files
goimports -w .      # Add/remove imports
golint ./...        # Lint check
go vet ./...        # Vet analysis
```

### Markdown Linting

All markdown files must pass linting checks:

- Use consistent heading levels (no skips)
- Surround code blocks with blank lines
- Specify language for all code blocks
- Keep line length reasonable

### Pre-commit Checks

Before committing:

1. Format code: `gofmt -w .`
2. Run tests: `go test ./...`
3. Verify linting: `golint ./...`
4. Check markdown: Use markdown linter
5. Verify go.mod: `go mod tidy`

---

## Workflow Summary

### For Adding a New Utility

1. Create `cmd/<util>/main.go` with CLI implementation
2. Create `cmd/<util>/README.md` with documentation
3. Create `pkg/<package>/` with core logic
4. Update root README with new utility
5. Commit with message: `feat: Add <util> CLI tool`
6. Update version tag

### For Updating Existing Code

1. Make changes to source files
2. Update relevant documentation
3. Commit with appropriate type (`fix:`, `docs:`, `refactor:`)
4. Only update version tag if releasing

### For Documentation-Only Changes

1. Update README files
2. Verify markdown linting
3. Commit with message: `docs: <description>`
4. No version tag update needed

---

## Key Reminders

- **Always follow markdown linting rules** in all `.md` files
- **Update documentation whenever code changes**
- **Use semantic versioning** for all releases
- **Follow commit conventions** for clear history
- **Test thoroughly** before creating version tags
- **Keep utilities modular and isolated**
- **Prefer Go standard library** over external dependencies
- **Document public APIs** with clear comments
- **Handle errors gracefully** with meaningful messages

---

## Questions and Escalation

For complex decisions not covered here:

1. Check existing utilities for patterns
2. Follow Go best practices and conventions
3. Prioritize clarity and maintainability
4. Document non-obvious implementations

---

**Last Updated**: November 24, 2025
**Repository**: github.com/murtazapatel89100/go-utils
