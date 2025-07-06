# Go Programming Cheatsheet

## Table of Contents
1. [Variable Declaration](#variable-declaration)
2. [Basic Program Structure](#basic-program-structure)
3. [Naming Conventions](#naming-conventions)
4. [Keywords in Go](#keywords-in-go)
5. [Best Practices](#best-practices)
6. [Frequently Used Standard Packages in Go](#frequently-used-standard-packages-in-go)
7. [Quick Reference](#quick-reference)

---

## Variable Declaration

### Single Variable Declaration

**Long Hand Declaration:**
```go
var name string
var age int
var isActive bool
```

**Short Hand Declaration:**
```go
name := "John"
age := 25
isActive := true
```

### Multiple Variable Declaration

**Long Hand - Same Line:**
```go
var name, email string
var width, height int
```

**Long Hand - Block Declaration:**
```go
var (
    name     string
    age      int
    isActive bool
)
```

**Short Hand - Multiple Variables:**
```go
x, y := 6, true        // x is int, y is bool
name, age := "John", 25 // name is string, age is int
```

> **Note:** Go is case-sensitive and the compiler infers types based on the assigned values in short-hand declarations. Also long hand variable declaration are known as __static type declaration__ while short hand variable declaration are known as __dynamic type declaration__

---

## Basic Program Structure

```go
package main

import "fmt"

func main() {
    // Write your code logic here
    fmt.Println("Hello, Go!")
}
```

---

## Naming Conventions

### 1. CamelCase for Identifiers
Use camelCase for general identifiers:
- ✅ `userName`, `getUserData`, `calculateTotal`
- ❌ `user_name`, `get_user_data`, `calculate_total`

### 2. PascalCase for Exported Names
Use PascalCase when a name should be exported (visible outside the package):
- ✅ `User`, `FetchData`, `IsActive`
- ❌ `user`, `fetchData`, `isActive` (for exported items)

### 3. camelCase for Unexported Names
Use camelCase for internal/private variables, functions, or types:
- ✅ `userName`, `processRequest`, `internalConfig`
- ❌ `UserName`, `ProcessRequest` (for internal items)

### 4. All Lowercase for Package Names
Package names should be short, descriptive, and lowercase:
- ✅ `http`, `json`, `strings`, `fmt`
- ❌ `Http`, `JsonUtils`, `MyPackage`

> **Tip:** Package names should not stutter with type names (avoid `user.UserModel`, prefer `user.User`)

### 5. Short, Meaningful Variable Names
Keep variable names concise but descriptive:
- ✅ `id`, `msg`, `err`, `user`, `res`, `count`
- ❌ `x1`, `val123`, `tempDataHolder`, `theUserVariable`

### 6. Constants
- **Exported constants:** Use PascalCase
  ```go
  const MaxRetries = 3
  const DefaultTimeout = 30
  ```
- **Unexported constants:** Use camelCase
  ```go
  const maxRetries = 3
  const defaultTimeout = 30
  ```

### 7. Acronyms
Keep acronyms consistently cased:
- ✅ `userID`, `httpRequest`, `JSONParser`, `XMLData`
- ❌ `UserId`, `Httprequest`, `Jsonparser`, `XmlData`

### 8. Interface Naming
Name interfaces with `-er` suffix when describing behavior:
- ✅ `Reader`, `Writer`, `Formatter`, `Validator`
- ❌ `ReadInterface`, `IReader`, `ReaderInterface`

### 9. Avoid Name Stuttering
Don't repeat package names in type names:
- ✅ `user.User`, `http.Client`, `json.Decoder`
- ❌ `user.UserModel`, `http.HttpClient`, `json.JsonDecoder`

### 10. Use Singular Nouns for Types
Type names should be singular:
- ✅ `User`, `Product`, `Order`, `Customer`
- ❌ `Users`, `Products`, `Orders`, `Customers`

---

## Keywords in Go
Keywords are reserved token that can't be used as an identifier 
- break
- case
- *chan 
- const
- continue 
- default
- defer
- else
- failthrough
- for
- func
- go
- goto 
- if
- import
- interface
- map
- package
- range
- return
- select 
- struct
- switch 
- type
- var

---

## Best Practices

### Error Handling
Always use `err` as the variable name for errors:
```go
data, err := fetchData()
if err != nil {
    // handle error
}
```

### Loop Variables
Use short names like `i`, `j` only in very limited scopes:
```go
for i := 0; i < len(items); i++ {
    // process items[i]
}

for _, item := range items {
    // process item
}
```

### Function Naming
Functions should describe what they do:
- ✅ `calculateTax()`, `validateEmail()`, `parseJSON()`
- ❌ `doStuff()`, `handleIt()`, `process()`

### Boolean Variables
Use descriptive names that read naturally:
- ✅ `isActive`, `hasPermission`, `canEdit`
- ❌ `active`, `permission`, `edit`

---

## Frequently Used Standard Packages in Go

Go provides a rich set of built-in packages in its standard library. Below are some of the most commonly used ones:

### 📦 Basic I/O
- `fmt` – Formatting and printing (e.g., `fmt.Println`, `fmt.Scanln`)
- `bufio` – Buffered input/output
- `os` – Platform-independent interface to operating system functionality
- `io` – Basic interfaces to I/O primitives
- `io/ioutil` – Deprecated; replaced by `os` and `io` for reading/writing files

### 📁 File and Directory Handling
- `path/filepath` – Manipulates filename paths in a way that's compatible with the operating system
- `os` – File creation, deletion, reading, writing, etc.

### 🧮 Math and Numbers
- `math` – Basic math functions like `math.Sqrt`, `math.Pow`
- `math/rand` – Random number generation
- `strconv` – String conversions for numbers (`strconv.Atoi`, `strconv.Itoa`)

### ⏰ Time and Date
- `time` – Time-related functions (e.g., `time.Now()`, `time.Sleep()`)

### 🧵 Concurrency
- `sync` – Synchronization primitives like `Mutex`, `WaitGroup`
- `sync/atomic` – Low-level atomic memory primitives
- `context` – Carries deadlines, cancellation signals, and request-scoped values

### 🌐 Networking
- `net` – Core networking interface
- `net/http` – HTTP client and server implementations
- `net/url` – URL parsing and formatting

### 🔐 Security and Encryption
- `crypto` – Cryptographic primitives
- `crypto/md5`, `crypto/sha256`, etc. – Hash algorithms

### 🧪 Testing
- `testing` – Support for automated testing
- `testify` – External but widely used testing framework (not part of standard lib)

### 📋 Data Handling
- `encoding/json` – JSON encoding and decoding
- `encoding/csv` – CSV file support
- `encoding/xml` – XML encoding and decoding

### 📜 Errors and Logging
- `errors` – Error creation and wrapping
- `log` – Simple logging package

### 🗃️ Containers
- `container/list` – Doubly linked list
- `container/heap` – Heap interface
- `container/ring` – Circular list

### 🧰 Utilities
- `reflect` – Reflection and dynamic typing
- `regexp` – Regular expressions
- `flag` – Command-line flag parsing

---

> 🧠 **Note**: All these packages come with Go and do not require separate installation. You can import them directly with `import "package-name"`.

---

## Quick Reference

| Element | Exported | Unexported | Example |
|---------|----------|------------|---------|
| **Variables** | PascalCase | camelCase | `UserName` / `userName` |
| **Functions** | PascalCase | camelCase | `GetUser()` / `getUser()` |
| **Constants** | PascalCase | camelCase | `MaxSize` / `maxSize` |
| **Types** | PascalCase | camelCase | `User` / `user` |
| **Interfaces** | PascalCase + -er | camelCase + -er | `Reader` / `reader` |
| **Packages** | lowercase | lowercase | `json`, `http` |

---

> 📌 **Remember:** Following these conventions helps your Go code stay idiomatic, readable, and consistent with community standards. The Go compiler and tools like `go fmt` and `golint` will help enforce many of these conventions automatically.