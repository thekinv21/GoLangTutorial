**English** | [Turkish](./README_TR.md) | [Russian](./README_RU.md)

# Operators in Go

Go has the same operators as most other languages

## Arithmetic operators

- **Addition:** +
- **Subtraction:** -
- **Multiplication:** \*
- **Division:** /
- **Modulo:** %

## Things to keep in mind:

#### To get a decimal result from division, both operands must be of the same (floating-point) type

```bash
fmt.Println(7.0 / 2) ---> Result: 3.5

fmt.Println(float64(a) / float64(b)) ---> Result: 3.5
```

#### Integer division with a negative value truncates toward zero.

```bash
fmt.Println(-7 / 2) ---> Result: -3 (truncates toward zero)
```

#### With modulo, the sign of the result comes from the left operand.

```bash
fmt.Println(-7 % 2) ---> Result: -1 (sign comes from the left operand)
```

#### Different types cannot be mixed

```bash
var x int = 5
var y float64 = 2.5

// fmt.Println(x + y) // ERROR: mismatched types

fmt.Println(float64(x) + y) // 7.5
```

## Comparison operators (result is bool)

- **Equal:** ==
- **Not Equal:** !=
- **Greater:** >
- **Less:** <
- **Greater or Equal:** >=
- **Less or Equal:** <=

```bash
a, b := 7, 2

fmt.Println(a == b, a != b, a > b, a < b, a >= b, a <= b)
// false true true false true false
```

## Logical operators

- **AND:** &&
- **OR:** ||
- **NOT:** !

```bash
t, f := true, false

fmt.Println(t && f) // false (AND)
fmt.Println(t || f) // true  (OR)
fmt.Println(!t)     // false (NOT)

```

### Important Notes:

Short-circuit: if the left side of && is false, the right side is not evaluated

```bash
var p *int

if p != nil && *p > 0 { // p is nil, so *p never runs, no panic
  fmt.Println("positive")
}
```

## Assignment operators

```bash
n := 10
n += 5  // 15
n -= 3  // 12
n *= 2  // 24
n /= 4  // 6
n %= 4  // 2

fmt.Println(n) // 2

// ++ and -- are statements, not expressions
n++
n--


// m := n++   // ERROR
// ++n        // ERROR, no prefix form
```

### NOTE:

Go has NO ternary operator (?:), so we have to write if/else

```bash

a, b := 7, 2

max := a

if b > a {
	max = b
}

fmt.Println(max) // 7
```
