[English](./README_EN.md) | **Turkish** | [Russian](./README_RU.md)

# Go Format Belirteçleri

| Belirteç | Tip                     | Örnek Kod                   | Çıktı                   |
| -------- | ----------------------- | --------------------------- | ----------------------- |
| `%s`     | string                  | `fmt.Printf("%s", "Hello")` | `Hello`                 |
| `%q`     | string (tırnaklı)       | `fmt.Printf("%q", "Hello")` | `"Hello"`               |
| `%d`     | int                     | `fmt.Printf("%d", 42)`      | `42`                    |
| `%f`     | float                   | `fmt.Printf("%f", 19.99)`   | `19.990000`             |
| `%.2f`   | float (2 basamak)       | `fmt.Printf("%.2f", 19.99)` | `19.99`                 |
| `%t`     | bool                    | `fmt.Printf("%t", true)`    | `true`                  |
| `%v`     | herhangi bir tip        | `fmt.Printf("%v", 42)`      | `42`                    |
| `%+v`    | struct (alan adlarıyla) | `fmt.Printf("%+v", u)`      | `{Name:Ali Age:30}`     |
| `%T`     | değişkenin tipi         | `fmt.Printf("%T", 3.14)`    | `float64`               |
| `%%`     | yüzde işareti           | `fmt.Printf("%%")`          | `%`                     |
| `\n`     | yeni satır              | `fmt.Printf("a\nb")`        | `a` ve `b` ayrı satırda |
