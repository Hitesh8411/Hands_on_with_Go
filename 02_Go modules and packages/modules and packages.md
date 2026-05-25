
This image shows the **older traditional Go workspace structure** using `GOPATH`.

# Structure in Image

```text id="nvigcz"
$HOME/
 └── go/
      ├── bin/
      ├── pkg/
      └── src/
           └── myproject/
                ├── main.go
                └── other files
```

Now let’s understand each folder one by one.

---

# 1. `bin/` → Executable Files

```text id="7clggo"
go/bin/
```

This folder stores the **compiled executable programs**.

When you build a Go project:

```bash
go build
```

or

```bash
go install
```

Go creates executable files here.

---

## Example

If your project name is:

```text id="4ajv22"
calculator
```

After build:

```text id="v1mvus"
go/bin/calculator
```

This is the runnable application.

---

# Simple Meaning

`bin/` = Final runnable programs

---

# 2. `pkg/` → Compiled Package Files

```text id="a3n8bd"
go/pkg/
```

This folder stores **compiled package objects**.

When Go compiles reusable packages/libraries, it keeps optimized compiled files here.

---

## Example

Suppose you create:

```go id="dc41o5"
package mathutils
```

Go compiles it and stores internal compiled files in:

```text id="kg44f4"
pkg/
```

---

# Simple Meaning

`pkg/` = Cached compiled libraries/packages

---

# 3. `src/` → Source Code Folder

```text id="gy8t5n"
go/src/
```

This is the MOST IMPORTANT folder.

All Go source code was traditionally written here.

---

# Example

```text id="yj1s5g"
src/
 └── myproject/
      ├── main.go
      └── helper.go
```

Your actual Go project lives inside `src/`.

---

# Simple Meaning

`src/` = Your real Go code

---

# What is `main.go`?

```go id="ml6x7j"
package main

func main() {

}
```

This is the entry point of Go application.

Execution starts from:

```go id="mj6lzn"
func main()
```

---

# Full Workflow

```text id="4z88cf"
Write code in src/
        ↓
Compile project
        ↓
Go stores package files in pkg/
        ↓
Executable created in bin/
```

---

# Example Flow

Suppose you create:

```text id="q7zsp4"
src/myapp/main.go
```

Then run:

```bash
go install
```

Go will:

## Step 1

Compile source code

## Step 2

Store package cache in:

```text id="v79jzv"
pkg/
```

## Step 3

Create executable in:

```text id="g6qdbt"
bin/myapp
```

---

# IMPORTANT ⚠️

This structure is from the **old Go workspace system (GOPATH)**.

Modern Go uses:

# Go Modules (`go.mod`)

Nowadays you usually create project like this:

```text id="wz6bgj"
myproject/
 ├── go.mod
 ├── main.go
 └── handlers/
```

No need to place everything inside:

```text id="7df93d"
go/src/
```

---

# Modern Go Project Structure

```text id="e6lmd0"
myapp/
 ├── cmd/
 ├── internal/
 ├── pkg/
 ├── api/
 ├── configs/
 ├── go.mod
 └── main.go
```

---

# Difference Between Old and Modern Go

| Old GOPATH                 | Modern Go Modules          |
| -------------------------- | -------------------------- |
| Uses `src/`                | Any folder                 |
| Uses `GOPATH`              | Uses `go.mod`              |
| Hard dependency management | Easy dependency management |
| Older system               | Modern standard            |

---

# What Should YOU Learn?

✅ Learn modern Go Modules system first.

You should know this old structure only because:

* Some tutorials still use it
* Some legacy projects use it
* Interviewers may ask about it

But for real projects:

✅ Use `go.mod`

---

# Modern Go Commands

## Initialize Project

```bash
go mod init myapp
```

## Run Project

```bash
go run main.go
```

## Build Project

```bash
go build
```

## Install Dependencies

```bash
go get package-name
```

---

# Final Summary

| Folder | Purpose                |
| ------ | ---------------------- |
| `src/` | Source code            |
| `pkg/` | Compiled package files |
| `bin/` | Executable binaries    |

---

# Easy Analogy

```text id="3jywtg"
src/  → Kitchen (where food is cooked)
pkg/  → Storage area
bin/  → Final packed food ready to eat
```
