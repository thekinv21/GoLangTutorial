**English** | [Turkish](./README_TR.md) | [Russian](./README_RU.md)

# Variables in Go

Go is a statically typed language: the type of every variable is known at compile time and never changes afterwards. There are several ways to declare a variable.

## 1. Declaring with var:

```bash
var fullName string = "Vadim Kiniabaev"
var age int = 26
```

- If you give an initial value, Go can infer the type by itself (type inference):

```bash
var fullName = "Vadim Kiniabaev" (inferred as string)
var age = 21 (inferred as integer)
```

## 2. Short declaration: :=

- The most common way inside functions

```bash
 message := "Hello World"
 title := "Full-Stack Developer"
```

#### NOTE:

:= can only be used inside a function. At package level (outside functions) you have to use var.

## 3. Zero values

Variables declared without a value automatically get the zero value of their type. There are no "undefined" variables in Go:

| Type                                          | Zero value |
| --------------------------------------------- | ---------- |
| `int`, `float64`                              | `0`        |
| `string`                                      | `""`       |
| `bool`                                        | `false`    |
| pointer, slice, map, channel, func, interface | `nil`      |

### Examples

```bash
 var message string -> ""

 var price float64 -> 0

 var count int -> 0

 var isActive bool -> false
```

## 4. Multiple declarations

```bash
 var a,b,c int = 1,2,3

 var (
	user = "admin"
	port = 4200
	debug = false
 )
```

Values can also be swapped in a single line:

```bash
 a,b = b,a
```

## 5. Things to watch out for

**An unused variable is a compile error.** Go does this on purpose, to keep the code clean:

```bash
x := 5   // If x is not used anywhere: "declared and not used" error
```

### NOTE:

\_ (blank identifier) is used to ignore a value on purpose:

```bash
_, err := strconv.Atoi("123") // the converted number is not used, only the error is checked
```

- **:=** requires at least one new variable. If all variables on the left already exist, it is an error; if at least one is new, the others are just assigned.

- **Shadowing:** Using := in an inner block does not change the outer variable, it creates a new one:

```bash
x := 1

if true {
	x := 2 // a new x, does not affect the outer one
	fmt.Println(x) // the value here is 2
}

// the value here is 1, because using **:=** in the inner block does not change the outer variable, it only creates a new one
fmt.Println(x)
```

## 6. Visibility (Exported / Unexported)

For package-level variables, the capital letter matters:

```bash
var Version string = "v1" // can be used both inside and outside the package

var newVersion string = "v2" // can only be used inside the package, not accessible from outside
```

## Quick Summary

- Use **:=** mostly inside functions and **var** outside them.
- Prefer var when you want to specify the type or start with the zero value.
- Type conversions are always explicit.
- Unused variables don't compile; use \_ if needed.
