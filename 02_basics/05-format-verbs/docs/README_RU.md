[English](./README_EN.md) | [Turkish](./README_TR.md) | **Russian**

# Спецификаторы формата в Go

| Спецификатор | Тип                           | Пример кода                 | Вывод                       |
| ------------ | ----------------------------- | --------------------------- | --------------------------- |
| `%s`         | string                        | `fmt.Printf("%s", "Hello")` | `Hello`                     |
| `%q`         | string (в кавычках)           | `fmt.Printf("%q", "Hello")` | `"Hello"`                   |
| `%d`         | int                           | `fmt.Printf("%d", 42)`      | `42`                        |
| `%f`         | float                         | `fmt.Printf("%f", 19.99)`   | `19.990000`                 |
| `%.2f`       | float (2 знака после запятой) | `fmt.Printf("%.2f", 19.99)` | `19.99`                     |
| `%t`         | bool                          | `fmt.Printf("%t", true)`    | `true`                      |
| `%v`         | любой тип                     | `fmt.Printf("%v", 42)`      | `42`                        |
| `%+v`        | struct (с именами полей)      | `fmt.Printf("%+v", u)`      | `{Name:Ali Age:30}`         |
| `%T`         | тип значения                  | `fmt.Printf("%T", 3.14)`    | `float64`                   |
| `%%`         | знак процента                 | `fmt.Printf("%%")`          | `%`                         |
| `\n`         | новая строка                  | `fmt.Printf("a\nb")`        | `a` и `b` на разных строках |
