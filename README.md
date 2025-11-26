# go-utils

A growing collection of reusable Go-based utilities, tools, and helper packages designed to simplify common development workflows across projects, environments, and teams.

This repository serves as a central place for production-ready Go utilities that can be used as:

- Standalone command-line tools (CLI)
- Importable Go packages
- Internal dev tooling for automation and scripting

The goal of this project is to provide clean, minimal, and extensible utilities that follow good engineering practices and can evolve over time without being limited to a single category.

---

## Project Vision

go-utils is intentionally structured to be future-proof. While it currently contains CLI-based utilities, the repository is designed to expand into multiple categories, including:

- Environment management tools
- File system helpers
- Automation utilities
- DevOps support modules
- Configuration loaders
- Validation and parsing libraries
- API helpers
- Build and deployment tooling

Each utility is modular, isolated, and versioned, allowing selective usage without coupling the entire repository.

---

## Current Structure

```
go-utils/
├── cmd/                # CLI applications
│   ├── create-env/     # CLI to generate .env from env.template
│   ├── setup-env/      # CLI to generate pipeline .env from env.test
│   ├── jsonfmt/        # CLI to format and validate JSON files
│   └── prettier/       # CLI to initialize and configure Prettier
│
├── pkg/                # Reusable Go packages
│   ├── envutils/       # Environment file helpers
│   │   ├── copy.go
│   │   └── replace.go
│   ├── formatter/      # JSON formatting utilities
│   │   └── formatter.go
│   └── prettier/       # Prettier configuration generation
│       └── generator.go
│
├── dist/               # Built binaries for releases (in .gitignore)
├── .github/            # GitHub-specific files
│   └── copilot.md      # Development guidelines
├── .gitignore
├── go.mod
├── LICENSE
└── README.md
```

---

## Included Utilities

### 1. create-env

A CLI tool that generates a `.env` file from an `env.template` file.

Primary use case:

- Local development environment setup

Key features:

- Copies only valid key-value pairs
- Filters out comments and empty lines
- Prevents overwriting existing `.env` files

---

### 2. setup-env

A CLI tool designed for automated environments such as CI/CD pipelines and test runners.

Primary use case:

- Pipeline and deployment configuration

Key features:

- Reads from `env.test`
- Replaces `{{VARIABLE_NAME}}` placeholders
- Pulls values from system environment variables
- Outputs a ready-to-use `.env` file

---

### 3. jsonfmt

A CLI tool for formatting and validating JSON files.

Primary use case:

- Format and validate JSON files with consistent indentation
- Quick JSON validation in development and CI/CD pipelines

Key features:

- Pretty-print JSON with 2-space indentation
- Validate JSON syntax
- Fast and minimal dependencies
- Cross-platform (Linux, macOS, Windows)

---

### 4. prettier

A CLI tool to initialize and configure Prettier for code formatting projects.

Primary use case:

- Quick setup of Prettier configuration in new projects
- Automated Prettier installation and `.prettierrc` generation

Key features:

- Support for multiple package managers (npm, pnpm)
- Generates standard `.prettierrc` configuration
- Configurable indentation and formatting rules
- Cross-platform (Linux, macOS, Windows)

---

## Reusable Package: envutils

The `envutils` package provides reusable logic used by the CLI tools and other Go projects.

Available functions:

- `CopyTemplate(templatePath, outputPath string) error`
  Copies a template file to a destination while filtering invalid or commented lines.

- `ReplaceTemplateVars(content string) string`
  Replaces `{{VAR_NAME}}` placeholders with values from environment variables.

This package can be imported into any Go project:

```go
import "github.com/murtazapatel89100/go-utils/pkg/envutils"
```

---

## Reusable Package: formatter

The `formatter` package provides JSON formatting utilities used by jsonfmt and other Go projects.

Available functions:

- `FormatJSON(input []byte) ([]byte, error)`
  Formats JSON input with consistent 2-space indentation and returns formatted output or error.

This package can be imported into any Go project:

```go
import "github.com/murtazapatel89100/go-utils/pkg/formatter"
```

---

## Reusable Package: prettier

The `prettier` package provides utilities for generating Prettier configuration files.

Available functions:

- `GeneratePrettierConfig() error`
  Generates a `.prettierrc` configuration file with standard formatting rules.

This package can be imported into any Go project:

```go
import "github.com/murtazapatel89100/go-utils/pkg/prettier"
```

---

## Usage Modes

### As CLI Tools

Download binaries from GitHub Releases and run directly:

```bash
./create-env
./setup-env <project-directory>
./jsonfmt input.json
./prettier -pkg npm
```

### As Go Packages

Add to your project:

```bash
go get github.com/murtazapatel89100/go-utils
```

Then import the required utilities as needed.

---

## Versioning Strategy

This repository follows semantic versioning:

- MAJOR: Breaking changes
- MINOR: New features or utilities
- PATCH: Bug fixes and improvements

Each tagged version corresponds to a stable snapshot of all tools and packages at that time.

---

## Release Process

1. Code is committed and reviewed
2. A version tag is created (example: v1.0.0)
3. Binaries are built and added to GitHub Releases
4. Documentation is updated accordingly

Releases are intended to be production-ready and stable.

---

## Design Principles

- Minimal dependencies
- Clear separation of concerns
- Modular and reusable architecture
- CLI tools as thin layers over core packages
- Sensible defaults with predictable behavior
- Expandable without breaking existing tools

---

## Contribution Guidelines

Contributions and internal expansions should follow:

- Each utility in its own folder
- Clear naming conventions
- Well-documented public functions
- No hard-coded environment assumptions
- Controlled and tested updates

---

## Future Direction

The repository will continue expanding as new development needs arise. Planned areas of growth include:

- Configuration loaders for multiple formats
- Deployment automation tools
- Code generation utilities
- Cross-platform helpers
- Secure secrets handling

This README will evolve alongside the repository as new utilities are introduced.

---

## License

This project is licensed under the terms specified in the LICENSE file.

---

Maintained by Murtaza Hakem Patel.
