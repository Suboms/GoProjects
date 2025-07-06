# Go Programming
First of all note, Go is case sensitive and below is how variables are declared in Go

`var name string`  // Long hand variable declaration 
`name := ""` // Short hand variable declaration

Also multiple variables can be declared in one line

 ### For Long Hand 
`var name, age string`
```go
var (
    x int
    y string
)
```
### For Short Hand

`x, y := 6, true` // Go compiler infer the type based on position so in this case, `x` is a **integer** and `y` is a **boolean**

```go
package main

func main(){
  / **
  write you code logic here
  **/
}
```

# Go (Golang) Naming Conventions

Follow these naming guidelines to write clean, idiomatic Go code.

---

## 1. CamelCase for Identifiers

- ✅ `userName`, `getUserData`
- ❌ `user_name`, `get_user_data`

---

## 2. PascalCase for Exported Names

Use PascalCase when a name should be exported (i.e., visible outside the package).

- ✅ `User`, `FetchData`, `IsActive`

---

## 3. camelCase for Unexported Names

Use camelCase for internal/private variables, functions, or types.

- ✅ `userName`, `processRequest`

---

## 4. All Lowercase for Package Names

- ✅ `http`, `json`, `strings`
- ❌ `Http`, `JsonUtils`, `MyPackage`

> 🔸 Tip: Package names should be short, descriptive, and not stutter with type names.

---

## 5. Short, Meaningful Variable Names

- ✅ `id`, `msg`, `err`, `user`, `res`
- ❌ `x1`, `val123`, `tempDataHolder`

---

## 6. Constants

- Use PascalCase for exported constants: `MaxRetries`
- Use camelCase for unexported constants: `defaultTimeout`

---

## 7. Acronyms

Use consistent casing with acronyms:
- ✅ `userID`, `httpRequest`, `JSONParser`
- ❌ `UserId`, `Httprequest`, `Jsonparser`

---

## 8. Interface Naming

Name interfaces with `-er` suffix when it fits the behavior:
- ✅ `Reader`, `Writer`, `Formatter`
- ❌ `ReadInterface`, `IReader`

---

## 9. Avoid Name Stuttering

Avoid repeating package names in type names.

- ✅ `user.User`
- ❌ `user.UserModel`

---

## 10. Use Singular Nouns for Types

- ✅ `User`, `Product`, `Order`
- ❌ `Users`, `Products`

---

## Bonus Tips

- Functions that return an error should use `err` as the variable name.
- Temporary loop variables can use short names like `i`, `j` only in very limited scopes.

---

> 📌 Following these conventions helps your Go code stay idiomatic, readable, and consistent with community standards.
