# jsonfmt

A command-line tool for formatting and validating JSON files.

## Usage

### Format a JSON file

```bash
./jsonfmt input.json
```

### Output formatted JSON to a file

```bash
./jsonfmt input.json > output.json
```

## Features

- Pretty-print JSON with 2-space indentation
- Validate JSON syntax
- Simple file-based input
- Fast and minimal dependencies
- Cross-platform (Linux, macOS, Windows)

## Examples

**Input file (input.json):**

```json
{"name":"John","age":30,"city":"New York"}
```

**Command:**

```bash
./jsonfmt input.json
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

If the file is not found or contains invalid JSON, jsonfmt will exit with status code 1 and display an error message:

```bash
$ ./jsonfmt invalid.json
Error reading file: open invalid.json: no such file or directory

$ ./jsonfmt bad.json
invalid JSON: invalid character 'i' looking for beginning of object key string
```

## Implementation

jsonfmt uses the `formatter` package which provides the `FormatJSON()` function for JSON processing. The CLI accepts a file path as an argument and outputs the formatted JSON to stdout.

---

Maintained as part of the go-utils toolkit.
