[English](./README_EN.md) | **Turkish** | [Russian](./README_RU.md)

# Go’da Operatörler

Başka dillerde olduğu gibi Go'da aynı operatörlere sahiptir

## Aritmetik operatörleri

- **Ekleme:** +
- **Çıkarma:** -
- **Çarpma:** \*
- **Bölme:** /
- **Mod Alma:** %

## Dikkat edilmesi gerekenler:

#### Ondalıklı sayıların bölünebilmesi için ikisinin de tipinin aynı olması gerekmektedir

```bash
fmt.Println(7.0 / 2) ---> Sonuç: 3.5

fmt.Println(float64(a) / float64(b)) ---> Sonuç: 3.5
```

#### Eksi değer ile bölme işlemi yaptığımızda sıfır'a yuvarlar.

```bash
fmt.Println(-7 / 2) ---> Sonuç: -3 (sıfıra doğru yuvarlar)
```

#### Mod alırken işaret sol taraftaki sayıdan alınır.

```bash
fmt.Println(-7 % 2) ---> Sonuç: -1 (işaret soldaki sayıdan gelir)
```

#### Farklı tipler karıştırılamaz

```bash
var x int = 5
var y float64 = 2.5

// fmt.Println(x + y) // HATA: mismatched types

fmt.Println(float64(x) + y) // 7.5
```

## Karşılaştırma operatörleri (sonuç bool)

- **Eşit:** ==
- **Eşit Değil:** !=
- **Büyük:** >
- **Küçük:** <
- **Büyük veya Eşit:** >=
- **Küçük veya Eşit:** <=

```bash
a, b := 7, 2

fmt.Println(a == b, a != b, a > b, a < b, a >= b, a <= b)
// false true true false true false
```

## Mantıksal operatörleri

- **VE:** &&
- **VEYA:** ||
- **DEĞİL:** !

```bash
t, f := true, false

fmt.Println(t && f) // false (VE)
fmt.Println(t || f) // true  (VEYA)
fmt.Println(!t)     // false (DEĞİL)

```

### Önemli Notlar:

Kısa devre: && solu false ise sağa bakmaz

```bash
var p *int

if p != nil && *p > 0 { // p nil olduğu için *p hiç çalışmaz, panik yok
  fmt.Println("pozitif")
}
```

## Atama operatörleri

```bash
n := 10
n += 5  // 15
n -= 3  // 12
n *= 2  // 24
n /= 4  // 6
n %= 4  // 2

fmt.Println(n) // 2

// ++ ve -- ifade değil, deyimdir
n++
n--


// m := n++   // HATA
// ++n        // HATA, ön ek yok
```

### NOT:

Go'da üçlü operatör (?:) YOKTUR yani if/else yazmak zorundayız

```bash

a, b := 7, 2

max := a

if b > a {
	max = b
}

fmt.Println(max) // 7
```
