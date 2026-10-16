**English** | [Turkish](./README_TR.md) | [Russian](./README_RU.md)

# Type Conversion in Go

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

## 1. Float → int: truncation (no rounding)

When a floating-point number is converted to an integer, the decimal part is **dropped**, not rounded:

```bash
f := 3.99
n := -3.99

int(f) --> 3
int(n) --> -3

int(math.Round(f)) --> 4 // use the math package to round
```

## 2. Overflow

If the value does not fit in the target type, Go does not raise an error, the number **wraps around**:

```bash
big := 300
int8(big) --> 44 // int8 range: -128 to 127

neg := -1
uint8(neg) --> 255 // uint8 cannot be negative
```

#### NOTE:

With a constant value you get a compile error instead: `int8(300)` --> "constant 300 overflows int8"

## 3. String ↔ number: the strconv package

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

## 4. Common mistake: string(number)

`string()` does not turn a number into text, it treats it as a **character code**:

```bash
x := 65

string(rune(x))  --> "A"  // 65 is the code of the letter "A"
strconv.Itoa(x)  --> "65" // correct usage
```

## 5. String ↔ []byte and []rune

```bash
[]byte("Go") --> [71 111]

s := "çay"
len(s)          --> 4 // number of bytes (ç takes 2 bytes)
len([]rune(s))  --> 3 // number of characters
```

Use **[]rune** to count characters in text with Turkish or Russian letters.

## 6. Untyped constants

Constants without an explicit type take their type from where they are used, no conversion needed:

```bash
const c = 10

var cf float64 = c --> 10 // no conversion needed
```

## Summary

| Conversion           | How                                   |
| -------------------- | ------------------------------------- |
| number → number      | `float64(x)`, `int(x)`, `int64(x)`    |
| number → string      | `strconv.Itoa`, `strconv.FormatInt`   |
| string → number      | `strconv.Atoi`, `strconv.ParseInt`    |
| string → bool        | `strconv.ParseBool`                   |
| string ↔ byte / rune | `[]byte(s)`, `[]rune(s)`, `string(b)` |
