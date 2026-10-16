[English](./README_EN.md) | [Turkish](./README_TR.md) | **Russian**

# Преобразование типов в Go

Go не преобразует типы автоматически, это нужно делать явно. Синтаксис всегда одинаковый: `T(значение)`

```bash
 var intNumber int = 20

 var floatNumber float64 = float64(intNumber)

 var uintNumber uint = uint(floatNumber)

 total := intNumber + floatNumber //ОШИБКА: int и float64 нельзя складывать

 total := float64(intNumber) + floatNumber // правильное использование
```

Даже типы из одного семейства не преобразуются автоматически:

```bash
var a int32 = 5
var b int64 = 7

sum := a + b        // ОШИБКА: int32 и int64 нельзя складывать
sum := int64(a) + b // правильное использование --> 12
```

## 1. Float → int: отбрасывание (без округления)

При преобразовании дробного числа в целое дробная часть **отбрасывается**, а не округляется:

```bash
f := 3.99
n := -3.99

int(f) --> 3
int(n) --> -3

int(math.Round(f)) --> 4 // для округления используется пакет math
```

## 2. Переполнение (overflow)

Если значение не помещается в целевой тип, Go не выдаёт ошибку, число **переходит через границу**:

```bash
big := 300
int8(big) --> 44 // диапазон int8: от -128 до 127

neg := -1
uint8(neg) --> 255 // uint8 не может быть отрицательным
```

#### ПРИМЕЧАНИЕ:

С константой вы получите ошибку компиляции: `int8(300)` --> "constant 300 overflows int8"

## 3. Строка ↔ число: пакет strconv

| Функция                       | Что делает                   | Пример                                                  |
| ----------------------------- | ---------------------------- | ------------------------------------------------------- |
| `strconv.Itoa(i)`             | int → string                 | `strconv.Itoa(20)` --> `"20"`                           |
| `strconv.Atoi(s)`             | string → int                 | `strconv.Atoi("123")` --> `123, nil`                    |
| `strconv.ParseInt(s, b, n)`   | string → int64 (основание b) | `strconv.ParseInt("ff", 16, 64)` --> `255`              |
| `strconv.ParseFloat(s, n)`    | string → float64             | `strconv.ParseFloat("3.14", 64)` --> `3.14`             |
| `strconv.ParseBool(s)`        | string → bool                | `strconv.ParseBool("true")` --> `true`                  |
| `strconv.FormatInt(i, b)`     | int64 → string (основание b) | `strconv.FormatInt(255, 2)` --> `"11111111"`            |
| `strconv.FormatFloat(f, ...)` | float64 → string             | `strconv.FormatFloat(3.14159, 'f', 2, 64)` --> `"3.14"` |

`Parse...` и `Atoi` всегда возвращают ещё и ошибку, не забывайте её проверять:

```bash
number, err := strconv.Atoi("abc")

if err != nil {
	fmt.Println(err) // strconv.Atoi: parsing "abc": invalid syntax
}

// при ошибке number равен 0
```

Чтобы вставить число в строку, можно также использовать **fmt.Sprintf**:

```bash
fmt.Sprintf("%d-%.1f", 7, 2.5) --> "7-2.5"
```

## 4. Частая ошибка: string(число)

`string()` не превращает число в текст, а воспринимает его как **код символа**:

```bash
x := 65

string(rune(x))  --> "A"  // 65 — это код буквы "A"
strconv.Itoa(x)  --> "65" // правильное использование
```

## 5. Строка ↔ []byte и []rune

```bash
[]byte("Go") --> [71 111]

s := "çay"
len(s)          --> 4 // количество байтов (ç занимает 2 байта)
len([]rune(s))  --> 3 // количество символов
```

Чтобы посчитать символы в тексте с турецкими или русскими буквами, используйте **[]rune**.

## 6. Нетипизированные константы (untyped constants)

Константы без явного типа получают тип там, где используются, преобразование не нужно:

```bash
const c = 10

var cf float64 = c --> 10 // преобразование не понадобилось
```

## Итог

| Преобразование       | Как сделать                           |
| -------------------- | ------------------------------------- |
| число → число        | `float64(x)`, `int(x)`, `int64(x)`    |
| число → строка       | `strconv.Itoa`, `strconv.FormatInt`   |
| строка → число       | `strconv.Atoi`, `strconv.ParseInt`    |
| строка → bool        | `strconv.ParseBool`                   |
| строка ↔ byte / rune | `[]byte(s)`, `[]rune(s)`, `string(b)` |
