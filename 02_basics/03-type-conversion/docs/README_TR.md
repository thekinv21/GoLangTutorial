[English](./README_EN.md) | **Turkish** | [Russian](./README_RU.md)

# Go’da Tip Dönüşümü

Go otomatik tip dönüşümü yapmaz, açıkça belirtmeniz gerekir. Yazımı her zaman aynıdır: `T(değer)`

```bash
 var intNumber int = 20

 var floatNumber float64 = float64(intNumber)

 var uintNumber uint = uint(floatNumber)

 total := intNumber + floatNumber //HATA: int ile float64 toplanamaz

 total := float64(intNumber) + floatNumber // doğru kullanım
```

Aynı ailedeki tipler bile birbirine otomatik dönüşmez:

```bash
var a int32 = 5
var b int64 = 7

sum := a + b        // HATA: int32 ile int64 toplanamaz
sum := int64(a) + b // doğru kullanım --> 12
```

## 1. Float → int: kesme (yuvarlama yok)

Ondalıklı sayı tam sayıya çevrilirken virgülden sonrası **atılır**, yuvarlanmaz:

```bash
f := 3.99
n := -3.99

int(f) --> 3
int(n) --> -3

int(math.Round(f)) --> 4 // yuvarlamak için math paketi kullanılır
```

## 2. Taşma (overflow)

Değer, hedef tipe sığmıyorsa Go hata vermez, sayı **başa sarar**:

```bash
big := 300
int8(big) --> 44 // int8 aralığı: -128 ile 127

neg := -1
uint8(neg) --> 255 // uint8 negatif olamaz
```

#### NOT:

Sabit (const) bir değerle yapılırsa derleme hatası alırsınız: `int8(300)` --> "constant 300 overflows int8"

## 3. String ↔ sayı: strconv paketi

| Fonksiyon                     | Ne yapar                 | Örnek                                                   |
| ----------------------------- | ------------------------ | ------------------------------------------------------- |
| `strconv.Itoa(i)`             | int → string             | `strconv.Itoa(20)` --> `"20"`                           |
| `strconv.Atoi(s)`             | string → int             | `strconv.Atoi("123")` --> `123, nil`                    |
| `strconv.ParseInt(s, b, n)`   | string → int64 (taban b) | `strconv.ParseInt("ff", 16, 64)` --> `255`              |
| `strconv.ParseFloat(s, n)`    | string → float64         | `strconv.ParseFloat("3.14", 64)` --> `3.14`             |
| `strconv.ParseBool(s)`        | string → bool            | `strconv.ParseBool("true")` --> `true`                  |
| `strconv.FormatInt(i, b)`     | int64 → string (taban b) | `strconv.FormatInt(255, 2)` --> `"11111111"`            |
| `strconv.FormatFloat(f, ...)` | float64 → string         | `strconv.FormatFloat(3.14159, 'f', 2, 64)` --> `"3.14"` |

`Parse...` ve `Atoi` fonksiyonları her zaman hata da döndürür, kontrol etmeyi unutmayın:

```bash
number, err := strconv.Atoi("abc")

if err != nil {
	fmt.Println(err) // strconv.Atoi: parsing "abc": invalid syntax
}

// hata durumunda number değeri 0 olur
```

Sayıyı string içine yerleştirmek için **fmt.Sprintf** de kullanılabilir:

```bash
fmt.Sprintf("%d-%.1f", 7, 2.5) --> "7-2.5"
```

## 4. Sık yapılan hata: string(sayı)

`string()` bir sayıyı metne çevirmez, onu **karakter kodu** olarak görür:

```bash
x := 65

string(rune(x))  --> "A"  // 65, "A" harfinin kodu
strconv.Itoa(x)  --> "65" // doğru kullanım
```

## 5. String ↔ []byte ve []rune

```bash
[]byte("Go") --> [71 111]

s := "çay"
len(s)          --> 4 // byte sayısı (ç 2 byte yer kaplar)
len([]rune(s))  --> 3 // karakter sayısı
```

Türkçe veya Rusça karakter içeren metinlerde karakter saymak için **[]rune** kullanın.

## 6. Tipsiz sabitler (untyped constants)

Tipi belirtilmemiş sabitler, kullanıldıkları yere göre tip alır, dönüşüm gerekmez:

```bash
const c = 10

var cf float64 = c --> 10 // dönüşüm gerekmedi
```

## Özet

| Dönüşüm              | Nasıl yapılır                         |
| -------------------- | ------------------------------------- |
| sayı → sayı          | `float64(x)`, `int(x)`, `int64(x)`    |
| sayı → string        | `strconv.Itoa`, `strconv.FormatInt`   |
| string → sayı        | `strconv.Atoi`, `strconv.ParseInt`    |
| string → bool        | `strconv.ParseBool`                   |
| string ↔ byte / rune | `[]byte(s)`, `[]rune(s)`, `string(b)` |
