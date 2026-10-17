**English** | [Turkish](./README_TR.md) | [Russian](./README_RU.md)

# The `fmt` Package in Go

## What is the `fmt` package?

`fmt` (format) is the package that handles **Input/Output (I/O)** operations. It is similar to C's `printf` and `scanf` functions, but it is designed around Go's type system. To use it, just import the package:

```bash
import "fmt"
```

Functions are grouped by their names:

| Prefix / Suffix                     | Meaning                                                      |
| ----------------------------------- | ------------------------------------------------------------ |
| `Print...`                          | Writes to standard output (`os.Stdout`)                      |
| `Sprint...`                         | Returns the result as a `string`                             |
| `Fprint...`                         | Writes to any `io.Writer`                                    |
| `Append...`                         | Appends the result to a `[]byte` slice (Go 1.19+)            |
| `Scan...` / `Sscan...` / `Fscan...` | Reads from stdin, from a string and from an `io.Reader`      |
| `...ln`                             | Adds a newline at the end (when reading, stops at line end)  |
| `...f`                              | Uses a format string                                         |

## Printing to the screen: `Print`, `Println`, `Printf`

```bash
fmt.Print("Hello","World") --> Result: HelloWorld

fmt.Print(1,2) --> Result: 1 2

fmt.Println("Hello","World") --> Result: Hello World\n

fmt.Printf("Age: %d\n", 20) --> Result: Age: 20
```

NOTE:

- **Print** - writes the values side by side, adds a space between two values only when **neither of them is a string**, and does not add a newline (**\n**) at the end
- **Println** - always adds a space between values and adds a newline (**\n**) at the end
- **Printf** - takes a format string and does not add a newline, so we have to add it manually

All three return **(n int, err error)**: the number of bytes written and a possible error. These values are usually ignored when printing to the screen.

### General verbs

| Verb  | Description                  | Example output                   |
| ----- | ---------------------------- | -------------------------------- |
| `%v`  | Default format               | `{Ali 30}`                       |
| `%+v` | Struct with field names      | `{Name:Ali Age:30}`              |
| `%#v` | Go-syntax representation     | `main.Person{Name:"Ali", Age:30}` |
| `%T`  | Type of the value            | `main.Person`                    |
| `%%`  | A literal percent sign       | `%`                              |

### Verbs for integers

| Verb        | Description                          | For `42`    |
| ----------- | ------------------------------------ | ----------- |
| `%d`        | Base 10                              | `42`        |
| `%b`        | Base 2                               | `101010`    |
| `%o`        | Base 8                               | `52`        |
| `%O`        | Base 8 with `0o` prefix              | `0o52`      |
| `%x` / `%X` | Base 16 (lowercase / uppercase)      | `2a` / `2A` |
| `%c`        | Unicode character                    | `*`         |
| `%q`        | Single-quoted character              | `'*'`       |
| `%U`        | Unicode format                       | `U+002A`    |

### Verbs for floats

| Verb        | Description                                              | For `3.14159`           |
| ----------- | -------------------------------------------------------- | ----------------------- |
| `%f` / `%F` | Decimal notation                                         | `3.141590`              |
| `%e` / `%E` | Scientific notation                                      | `3.141590e+00`          |
| `%g` / `%G` | `%e` or `%f` depending on the value, no trailing zeros   | `3.14159`               |
| `%x`        | Hexadecimal scientific notation                          | `0x1.921f9f01b866ep+01` |

### Verbs for strings and bytes

| Verb | Description                   | For `"Go"` |
| ---- | ----------------------------- | ---------- |
| `%s` | Plain string                  | `Go`       |
| `%q` | Double-quoted, escaped        | `"Go"`     |
| `%x` | Hexadecimal, per byte         | `476f`     |

### Verbs for booleans and pointers

| Verb | Description     | Example output   |
| ---- | --------------- | ---------------- |
| `%t` | Boolean value   | `true` / `false` |
| `%p` | Memory address  | `0xc000012345`   |

#### Pointers: `%p` → a memory address such as 0xc000012345.

```bash
type Person struct {
	Name string
	Age int
}

myPerson := Person{"Vadim", 26}

fmt.Printf("%v\n", myPerson) --> Result: {Vadim 26}

fmt.Printf("%+v\n", myPerson) --> Result: {Name:Vadim Age:26}

fmt.Printf("%#v\n", myPerson) --> Result: main.Person{Name:"Vadim", Age:26}

fmt.Printf("%T\n", myPerson) --> Result: main.Person
```

## Building strings: `Sprint`, `Sprintln`, `Sprintf`

These functions work like the Print family, but instead of printing to the screen they return the result as a **string**

```bash
stringValue := fmt.Sprintf("%s is %d years old", "Vadim", 26) --> Result: "Vadim is 26 years old"

totalValue := fmt.Sprint("Total:", 10) --> Result: "Total:10"

uValue := fmt.Sprintln("a", "b") --> Result: "a b\n"
```

They are used anywhere you need to build a string, such as log messages, file names or text templates

The **Append** family (Go 1.19+): appends to an existing []byte slice. It is efficient because it does not allocate a new string:

```bash
buf := []byte("Total: ")
buf = fmt.Appendf(buf, "%d", 42) --> Result: "Total: 42"
```

## Writing to other destinations: `Fprint`, `Fprintln`, `Fprintf`

They take an **io.Writer** as the first parameter. We can write to anything that has a **Write** method (a file, a network connection, an HTTP response, a buffer):

```bash
// Writing to the standard error stream

fmt.Fprintln(os.Stderr, "an error occurred")

// Writing to a file:

newFile, err := os.Create("newFile.txt")
if err != nil {
	log.Fatal(err)
}
defer newFile.Close()

fmt.Fprintf(newFile, "Row %d\n", 1)

// Writing to an HTTP response

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}

// Writing to a buffer

var newBuffer strings.Builder
fmt.Fprintf(&newBuffer, "%d-%d", 1, 2)
```

In fact, **Printf** is just a shortcut for **Fprintf(os.Stdout, ...)**

## Creating errors: `Errorf`

**Errorf** creates an error value with a formatted message:

```bash
err := fmt.Errorf("user %d not found", id)
```

Wrapping errors with **%w** (Go 1.13+): this verb keeps the underlying error, so you can reach the original error with **errors.Is** and **errors.As**:

```bash
_, err := os.Open("missing.txt")

if err != nil {
	err = fmt.Errorf("could not read config: %w", err)
}

if errors.Is(err, os.ErrNotExist) {
	fmt.Println("File not found") // this runs
}
```

- If you use **%v**, the message looks the same, but the chain is broken and errors.Is no longer finds a match.
- Since Go 1.20, a single Errorf can contain more than one **%w**
- For fixed messages, **errors.New("...")** is the simpler option

## Reading user input: `Scan`, `Scanln`, `Scanf`

These functions return the values through pointers, and they also return the number of values read and an error

```bash
var Name string
var Age int

// Scan: treats spaces and newlines as separators
n, err := fmt.Scan(&Name, &Age)

// Scanln: stops at a newline; all values must be on the same line
fmt.Scanln(&Name, &Age)

// Scanf: reads according to a specific format
fmt.Scanf("%s %d", &Name, &Age)
```

### Newline behavior of the reading functions

| Function | Newline behavior                                                       |
| -------- | ---------------------------------------------------------------------- |
| `Scan`   | Treats a newline like a space and continues until all values are read  |
| `Scanln` | Stops at the end of the line. Returns an error if a value is missing   |
| `Scanf`  | Newlines in the input must match those in the format string            |

**Important:** the Scan family splits words by spaces, so **"Ali Veli"** cannot be read into a single string. To read lines that contain spaces, use **bufio**

```bash
reader := bufio.NewReader(os.Stdin)
line, _ := reader.ReadString('\n')
line = strings.TrimSpace(line)

// or read line by line:

scanner := bufio.NewScanner(os.Stdin)

for scanner.Scan() {
	fmt.Println(scanner.Text())
}
```

To read from an **io.Reader** there are also **Fscan**, **Fscanln** and **Fscanf**

## Reading from a string: `Sscan`, `Sscanf`

They take a string as input instead of stdin

```bash
var x, y int

fmt.Sscan("10 20", &x, &y)

var day, month, year int
fmt.Sscanf("17/10/26", "%d/%d/%d", &day, &month, &year)
```

## Summary

The **`fmt` package** is the standard Go package for printing to the screen, building strings and reading user input.

### You can tell what a function does from its name

- `Print` writes to the screen, `Sprint` returns a string, `Fprint` writes to another destination such as a file.
- Functions ending in `ln` add a newline, functions ending in `f` use a format string (like `%d`, `%s`).
- The `Scan` family reads from the user, the `Sscan` family reads from a string.
- `Errorf` creates a formatted error message. With `%w` the underlying error is preserved.

### Most used verbs

- `%v` any value, `%+v` struct with field names, `%T` type name
- `%d` integer, `%f` float, `%s` string, `%t` boolean
- `%x` hexadecimal, `%q` quoted representation

### Formatting

`%8d` sets the width, `%-8d` aligns to the left, `%05d` pads with zeros, `%.2f` shows two decimal places.

### Stringer

If you write a `String() string` method for a type, `fmt` uses that method when printing the type.

### Things to keep in mind

- `Printf` does not add a newline by itself, don't forget to write `\n`.
- When reading, pass the address of the variable: `fmt.Scan(&x)`.
- `Scan` stops at a space. Use `bufio` to read a line that contains spaces.
- Don't print the same value with `%v` inside its own `String()` method, the program will end up in infinite recursion.
- Run `go vet` to catch format errors.
