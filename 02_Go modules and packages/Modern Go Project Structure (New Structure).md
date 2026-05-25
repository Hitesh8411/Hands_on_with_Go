# Modern Go Project Structure (New Structure) ✅

Nowadays Go uses **Go Modules (`go.mod`)** instead of old `GOPATH/src/pkg/bin` structure.

---

# Modern Go Structure

```text id="sz9m2u"
myapp/
│
├── cmd/
│    └── app/
│         └── main.go
│
├── internal/
│    ├── handlers/
│    ├── services/
│    └── database/
│
├── pkg/
│    └── utils/
│
├── api/
│    └── routes.go
│
├── configs/
│    └── config.go
│
├── middleware/
│    └── auth.go
│
├── models/
│    └── user.go
│
├── repositories/
│    └── user_repo.go
│
├── go.mod
├── go.sum
└── README.md
```

---

# Explanation of Each Folder

---

# 1. `cmd/`

```text id="4n4xsp"
cmd/
```

Contains the main application entry point.

Example:

```text id="d2q2ix"
cmd/app/main.go
```

This is where your app starts.

---

# Example

```go id="2nwy9v"
package main

func main() {

}
```

---

# Simple Meaning

`cmd/` = Starting point of application

---

# 2. `internal/`

```text id="h8y89r"
internal/
```

Contains private application code.

Other projects cannot import this.

Usually contains:

* business logic
* services
* handlers
* database logic

---

# Example Structure

```text id="r14fb0"
internal/
 ├── handlers/
 ├── services/
 └── database/
```

---

# Simple Meaning

`internal/` = Main backend logic

---

# 3. `pkg/`

```text id="vymc8f"
pkg/
```

Contains reusable packages/utilities.

Example:

* validators
* helpers
* utility functions

---

# Example

```text id="k0u06r"
pkg/
 └── utils/
```

---

# Simple Meaning

`pkg/` = Shared reusable code

---

# 4. `api/`

```text id="v4mwmj"
api/
```

Contains:

* API routes
* API definitions
* Swagger files

---

# Example

```text id="22k2dl"
api/routes.go
```

---

# Simple Meaning

`api/` = API routing layer

---

# 5. `configs/`

```text id="jlwmq4"
configs/
```

Stores application configuration.

Example:

* database config
* environment variables
* server settings

---

# Example

```go id="g9s05d"
DB_URL=localhost
PORT=8080
```

---

# Simple Meaning

`configs/` = Application settings

---

# 6. `middleware/`

```text id="j71bh5"
middleware/
```

Contains middleware logic.

Example:

* JWT authentication
* logging
* request validation

---

# Example

```text id="r4bkhf"
middleware/auth.go
```

---

# Simple Meaning

`middleware/` = Runs before request reaches API

---

# 7. `models/`

```text id="u5gnf9"
models/
```

Contains database models/structures.

---

# Example

```go id="szn1tx"
type User struct {
    ID   int
    Name string
}
```

---

# Simple Meaning

`models/` = Database structure definitions

---

# 8. `repositories/`

```text id="8pw3wb"
repositories/
```

Handles database queries.

---

# Example

```go id="b68lg9"
func GetUserByID(id int) {

}
```

---

# Simple Meaning

`repositories/` = Database operations

---

# 9. `go.mod` ⭐ VERY IMPORTANT

```text id="e1pj06"
go.mod
```

This is the heart of modern Go projects.

It manages:

* project name
* dependencies
* package versions

---

# Example

```go id="l4nyb4"
module myapp

go 1.24
```

---

# Simple Meaning

`go.mod` = Project manager

---

# 10. `go.sum`

```text id="m7wpnd"
go.sum
```

Stores dependency checksum/security info.

Automatically generated.

---

# Simple Meaning

`go.sum` = Dependency security tracking

---

# How Modern Go Works

```text id="ry1hvs"
Write code
     ↓
go.mod manages dependencies
     ↓
Build application
     ↓
Single executable generated
```

---

# Create Modern Go Project

## Step 1

Create folder:

```bash id="b7kkm2"
mkdir myapp
```

---

## Step 2

Initialize module:

```bash id="jlwmr8"
go mod init myapp
```

This creates:

```text id="jlwmrq"
go.mod
```

---

## Step 3

Create main.go

```go id="s60y8j"
package main

import "fmt"

func main() {
    fmt.Println("Hello Go")
}
```

---

## Step 4

Run app

```bash id="8oh6gz"
go run main.go
```

---

# Real Backend Structure Example

```text id="vbz0qh"
ecommerce-api/
│
├── cmd/
├── internal/
├── middleware/
├── models/
├── repositories/
├── routes/
├── services/
├── utils/
├── configs/
├── go.mod
└── main.go
```

---

# Why Modern Structure is Better?

| Advantage             | Why Important       |
| --------------------- | ------------------- |
| Modular               | Easy large projects |
| Clean Architecture    | Better organization |
| Dependency Management | Easy packages       |
| Scalable              | Production ready    |
| Team Friendly         | Organized code      |

---

# Which Structure Should YOU Use?

✅ Use Modern Go Modules structure.

Do NOT start with old:

```text id="d4y5x4"
src/pkg/bin
```

Use:

```text id="n9iwt9"
go.mod
```

based projects.

---

# Most Important Folders for Beginners

Focus first on:

```text id="q10vlu"
main.go
go.mod
handlers/
models/
routes/
services/
```

These are enough to build APIs and backend projects.

---

# Final Easy Analogy

```text id="oj47yx"
cmd/          → Engine Start
internal/     → Car Engine
models/       → Car Design
repositories/ → Fuel System
middleware/   → Security Check
configs/      → Settings
go.mod        → Car Registration
```
