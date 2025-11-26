# prettier-init — Prettier Configuration Generator (Go CLI)

## Overview

`prettier-init` is a Go-based command-line utility that quickly sets up Prettier in any JavaScript or TypeScript project.  
It performs two tasks:

1. Generates a `.prettierrc` file with a standardised formatting configuration  
2. Installs Prettier using your preferred package manager (`npm` or `pnpm`)

This tool is designed to help developers initialise consistent formatting rules across new or existing projects with a single command.

---

## Features

- Generates a `.prettierrc` file with opinionated defaults
- Supports both **npm** and **pnpm**
- Optional `--pkg` flag to choose a package manager
- Default package manager: `npm`
- Executes the install command automatically
- Lightweight and cross-platform (Linux, macOS, Windows)

---

## Prettier Config Generated

The CLI writes the following configuration into `.prettierrc`:

```json
{
  "semi": false,
  "singleQuote": true,
  "tabWidth": 2,
  "trailingComma": "es5",
  "printWidth": 100,
  "bracketSpacing": true,
  "jsxSingleQuote": false,
  "arrowParens": "always",
  "endOfLine": "lf"
}
