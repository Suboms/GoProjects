# Go Programming Cheatsheet

## Table of Contents
1. [Variable Declaration](#variable-declaration)
2. [Basic Program Structure](#basic-program-structure)
3. [Naming Conventions](#naming-conventions)
4. [Best Practices](#best-practices)

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

> **Note:** Go is case-sensitive and the compiler infers types based on the assigned values in short-hand declarations.

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