**English** | [Turkish](./README_TR.md) | [Russian](./README_RU.md)

# Go Format Verbs

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
