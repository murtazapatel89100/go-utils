# jsonfmt

A command-line tool for formatting and validating JSON files.

## Usage

### Format from file

```bash
./jsonfmt input.json
```

### Format from stdin

```bash
echo '{"name":"John","age":30}' | ./jsonfmt
```

### Output formatted JSON

```bash
./jsonfmt data.json > output.json
```

## Features

- Pretty-print JSON with 2-space indentation
- Validate JSON syntax
- Support both file and stdin input
- Fast and minimal dependencies
- Cross-platform (Linux, macOS, Windows)

## Examples

**Input:**

```json
{"name":"John","age":30,"city":"New York"}
```

**Output:**

```json
{
  "name": "John",
  "age": 30,
  "city": "New York"
}
```

## Installation

Build from source:

```bash
go build -o jsonfmt ./cmd/jsonfmt
```

Or download pre-built binaries from GitHub Releases.

## Error Handling

If invalid JSON is provided, jsonfmt will exit with status code 1 and display an error message:

```bash
$ echo '{invalid}' | jsonfmt
Error formatting JSON: invalid JSON: invalid character 'i' looking for beginning of object key string
```

## Implementation

jsonfmt uses the `formatter` package which provides the `FormatJSON()` function for JSON processing. The CLI is a thin wrapper around this core utility, allowing it to be used both as a standalone tool and imported into other Go projects.
