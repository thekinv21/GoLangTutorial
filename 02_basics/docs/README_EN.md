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

## 5. Basic types

- **Integers:** int, int8, int16, int32, int64, and for unsigned uint, uint8 (byte), etc.
- **Floating point:** float32, float64
- **Text:** string, and rune (alias for int32) for a single character
- **Boolean:** bool
- **Complex numbers:** complex64, complex128

### Integers (Signed)

| Type             | Size        | Minimum                    | Maximum                    |
| ---------------- | ----------- | -------------------------- | -------------------------- |
| `int8`           | 8 bit       | -128                       | +127                       |
| `int16`          | 16 bit      | -32,768                    | +32,767                    |
| `int32` (`rune`) | 32 bit      | -2,147,483,648             | +2,147,483,647             |
| `int64`          | 64 bit      | -9,223,372,036,854,775,808 | +9,223,372,036,854,775,807 |
| `int`            | 32 / 64 bit | Platform dependent         | Platform dependent         |

### Integers (Unsigned)

| Type             | Size        | Minimum | Maximum                        |
| ---------------- | ----------- | ------- | ------------------------------ |
| `uint8` (`byte`) | 8 bit       | 0       | 255                            |
| `uint16`         | 16 bit      | 0       | 65,535                         |
| `uint32`         | 32 bit      | 0       | 4,294,967,295                  |
| `uint64`         | 64 bit      | 0       | 18,446,744,073,709,551,615     |
| `uint`           | 32 / 64 bit | 0       | Platform dependent             |
| `uintptr`        | 32 / 64 bit | 0       | Large enough to hold a pointer |

### Floating-Point Numbers

| Type      | Size   | Range (approx.)       | Precision     |
| --------- | ------ | --------------------- | ------------- |
| `float32` | 32 bit | ±1.18e-38 … ±3.4e38   | ~6-7 digits   |
| `float64` | 64 bit | ±2.23e-308 … ±1.8e308 | ~15-16 digits |

### Complex Numbers

| Type         | Size    | Description                            |
| ------------ | ------- | -------------------------------------- |
| `complex64`  | 64 bit  | Real and imaginary parts are `float32` |
| `complex128` | 128 bit | Real and imaginary parts are `float64` |

### Other Types

| Type     | Values                                 | Zero Value |
| -------- | -------------------------------------- | ---------- |
| `bool`   | `true` / `false`                       | `false`    |
| `string` | UTF-8 text (immutable)                 | `""`       |
| `byte`   | Alias for `uint8`                      | `0`        |
| `rune`   | Alias for `int32` (Unicode code point) | `0`        |

> Note: The zero value of all numeric types is `0`. Limits are available in the `math` package: `math.MaxInt8`, `math.MinInt64`, `math.MaxUint32`, `math.MaxFloat64`.

## 6. Type conversion

Go does not convert types automatically, you have to do it explicitly. The syntax is always the same: `T(value)`

```bash
 var intNumber int = 20

 var floatNumber float64 = float64(intNumber)

 var uintNumber uint = uint(floatNumber)

 total := intNumber + floatNumber //ERROR: int and float64 cannot be added

 total := float64(intNumber) + floatNumber // correct usage
```

Even types from the same family are not converted automatically:

```bash
var a int32 = 5
var b int64 = 7

sum := a + b        // ERROR: int32 and int64 cannot be added
sum := int64(a) + b // correct usage --> 12
```

### 6.1 Float → int: truncation (no rounding)

When a floating-point number is converted to an integer, the decimal part is **dropped**, not rounded:

```bash
f := 3.99
n := -3.99

int(f) --> 3
int(n) --> -3

int(math.Round(f)) --> 4 // use the math package to round
```

### 6.2 Overflow

If the value does not fit in the target type, Go does not raise an error, the number **wraps around**:

```bash
big := 300
int8(big) --> 44 // int8 range: -128 to 127

neg := -1
uint8(neg) --> 255 // uint8 cannot be negative
```

#### NOTE:

With a constant value you get a compile error instead: `int8(300)` --> "constant 300 overflows int8"

### 6.3 String ↔ number: the strconv package

| Function                      | What it does            | Example                                                 |
| ----------------------------- | ----------------------- | ------------------------------------------------------- |
| `strconv.Itoa(i)`             | int → string            | `strconv.Itoa(20)` --> `"20"`                           |
| `strconv.Atoi(s)`             | string → int            | `strconv.Atoi("123")` --> `123, nil`                    |
| `strconv.ParseInt(s, b, n)`   | string → int64 (base b) | `strconv.ParseInt("ff", 16, 64)` --> `255`              |
| `strconv.ParseFloat(s, n)`    | string → float64        | `strconv.ParseFloat("3.14", 64)` --> `3.14`             |
| `strconv.ParseBool(s)`        | string → bool           | `strconv.ParseBool("true")` --> `true`                  |
| `strconv.FormatInt(i, b)`     | int64 → string (base b) | `strconv.FormatInt(255, 2)` --> `"11111111"`            |
| `strconv.FormatFloat(f, ...)` | float64 → string        | `strconv.FormatFloat(3.14159, 'f', 2, 64)` --> `"3.14"` |

`Parse...` and `Atoi` always return an error as well, don't forget to check it:

```bash
number, err := strconv.Atoi("abc")

if err != nil {
	fmt.Println(err) // strconv.Atoi: parsing "abc": invalid syntax
}

// on error, number is 0
```

**fmt.Sprintf** can also be used to put numbers into a string:

```bash
fmt.Sprintf("%d-%.1f", 7, 2.5) --> "7-2.5"
```

### 6.4 Common mistake: string(number)

`string()` does not turn a number into text, it treats it as a **character code**:

```bash
x := 65

string(rune(x))  --> "A"  // 65 is the code of the letter "A"
strconv.Itoa(x)  --> "65" // correct usage
```

### 6.5 String ↔ []byte and []rune

```bash
[]byte("Go") --> [71 111]

s := "çay"
len(s)          --> 4 // number of bytes (ç takes 2 bytes)
len([]rune(s))  --> 3 // number of characters
```

Use **[]rune** to count characters in text with Turkish or Russian letters.

### 6.6 Untyped constants

Constants without an explicit type take their type from where they are used, no conversion needed:

```bash
const c = 10

var cf float64 = c --> 10 // no conversion needed
```

### Summary

| Conversion           | How                                   |
| -------------------- | ------------------------------------- |
| number → number      | `float64(x)`, `int(x)`, `int64(x)`    |
| number → string      | `strconv.Itoa`, `strconv.FormatInt`   |
| string → number      | `strconv.Atoi`, `strconv.ParseInt`    |
| string → bool        | `strconv.ParseBool`                   |
| string ↔ byte / rune | `[]byte(s)`, `[]rune(s)`, `string(b)` |

## 7. Constants: const

```bash
const Pi = 3.14159

const (
    Monday = iota     // 0
    Tuesday           // 1
    Wednesday         // 2
)
```

**iota** is a counter that increases by one on each line; it is ideal for enum-like constructs.

## 8. Things to watch out for

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

## 9. Visibility (Exported / Unexported)

For package-level variables, the capital letter matters:

```bash
var Version string = "v1" // can be used both inside and outside the package

var newVersion string = "v2" // can only be used inside the package, not accessible from outside
```

## Go Format Verbs

| Verb   | Type                      | Example Code                | Output                        |
| ------ | ------------------------- | --------------------------- | ----------------------------- |
| `%s`   | string                    | `fmt.Printf("%s", "Hello")` | `Hello`                       |
| `%q`   | string (quoted)           | `fmt.Printf("%q", "Hello")` | `"Hello"`                     |
| `%d`   | int                       | `fmt.Printf("%d", 42)`      | `42`                          |
| `%f`   | float                     | `fmt.Printf("%f", 19.99)`   | `19.990000`                   |
| `%.2f` | float (2 decimals)        | `fmt.Printf("%.2f", 19.99)` | `19.99`                       |
| `%t`   | bool                      | `fmt.Printf("%t", true)`    | `true`                        |
| `%v`   | any type                  | `fmt.Printf("%v", 42)`      | `42`                          |
| `%+v`  | struct (with field names) | `fmt.Printf("%+v", u)`      | `{Name:Ali Age:30}`           |
| `%T`   | type of the value         | `fmt.Printf("%T", 3.14)`    | `float64`                     |
| `%%`   | percent sign              | `fmt.Printf("%%")`          | `%`                           |
| `\n`   | new line                  | `fmt.Printf("a\nb")`        | `a` and `b` on separate lines |

## Quick Summary

- Use **:=** mostly inside functions and **var** outside them.
- Prefer var when you want to specify the type or start with the zero value.
- Type conversions are always explicit.
- Unused variables don't compile; use \_ if needed.
