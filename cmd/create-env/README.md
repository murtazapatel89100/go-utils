# create-env CLI Tool (Release Version)

A lightweight command-line utility written in Go that automatically generates a `.env` file from an `env.template` file.

This tool is intended for developers who want a quick and consistent way to bootstrap environment configuration for local development.

---

## What does create-env do?

When you run the command:

```bash
create-env
```

It will:

1. Look for a file named `env.template` in the current directory
2. Extract valid environment variables in `KEY=VALUE` format
3. Create a new file named `.env`
4. Ignore:

   - Empty lines
   - Comments starting with `#`
   - Invalid formats

---

## Required File Conventions

For this tool to work correctly, your project directory MUST contain:

### Input file (you create this)

```
env.template
```

### Output file (auto-generated)

```
.env
```

Both files must exist / be created in the SAME directory where you run the command.

---

## Installation (From GitHub Releases)

1. Go to the **Releases** section of this repository
2. Download the correct binary for your OS:

| OS      | File to download       |
| ------- | ---------------------- |
| Linux   | create-env-linux       |
| macOS   | create-env-mac         |
| Windows | create-env-windows.exe |

3. Move the binary to a directory in your PATH (recommended):

### Linux / macOS

```bash
chmod +x create-env-linux
mv create-env-linux /usr/local/bin/create-env
```

### Windows

Move `create-env-windows.exe` to any folder in your PATH,
or rename it to `create-env.exe` for convenience.

Now you can run it globally:

```bash
create-env
```

---

## Alternative Usage: Run as Local Project Binary (./create-env)

You do NOT have to install this tool globally. You can keep the binary inside your own project and run it directly using `./`.

### When should you use this?

- You want the tool version locked per project
- You don’t want to modify PATH
- You want to commit the binary to your repo
- You are running it in CI pipelines

### Example structure

```
my-project/
├── create-env
├── env.template
└── src/
```

### Steps

1. Download the correct binary from GitHub Releases
2. Place it in your project root
3. Give it execute permission (Linux / macOS):

```bash
chmod +x create-env
```

4. Run it using:

```bash
./create-env
```

The `.env` file will be generated in the same directory.

This method does NOT require global installation or PATH configuration.

---

## How To Use

### 1. Create the template file

Inside your project:

```bash
touch env.template
```

Add your variables:

```env
# Database
DB_HOST=localhost
DB_PORT=5432

# API
API_KEY=your-api-key
```

### 2. Run the tool

From the same directory:

```bash
create-env
```

### 3. Result

A new `.env` file will be created:

```env
DB_HOST=localhost
DB_PORT=5432
API_KEY=your-api-key
```

---

## Example Project Structure

```
my-project/
├── env.template
├── .env   ✅ auto generated
└── src/
```

---

## Important Rules

| Rule              | Requirement                    |
| ----------------- | ------------------------------ |
| Template filename | Must be exactly `env.template` |
| Output filename   | Always `.env`                  |
| Run location      | Same directory as env.template |
| Format            | KEY=VALUE only                 |

---

## Common Errors

### "template not found"

Cause:

- `env.template` does not exist
- You are running the command in the wrong directory

Fix:

```bash
ls env.template
```

---

## Technical Overview

Internally the tool performs:

```go
envutils.CopyTemplate("env.template", ".env")
```

This function:

- Parses and filters valid environment variables
- Writes them cleanly to `.env`

---

## Safe to Run Multiple Times

Running `create-env` again will regenerate the `.env` file using the template content.

---

## Summary

1. Create `env.template`
2. Add KEY=VALUE pairs
3. Run:

```bash
create-env
```

4. `.env` is generated automatically

---

## Maintainer

Murtaza Patel
Go Utilities Toolkit

GitHub: [https://github.com/murtazapatel89100/go-utils](https://github.com/murtazapatel89100/go-utils)

---

If you need CI/pipeline support or template variable replacement, check the companion tool: **setup-env**.
