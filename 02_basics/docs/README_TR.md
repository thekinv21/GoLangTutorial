[English](./README_EN.md) | **Turkish** | [Russian](./README_RU.md)

# Go’da Değişkenler

Go statik tipli bir dildir ve her değişkenin tipi derleme anında bellidir ve sonradan değişmez. Değişken tanımlamanın birkaç yolu vardır.

## 1. var ile tanımlama:

```bash
var fullName string = "Vadim Kiniabaev"
var age int = 26
```

- Default değerini verirseniz Go tipi kendisi çıkarabilir (type inference):

```bash
var fullName = "Vadim Kiniabaev" (string olarak algılanır)
var age = 21 (integer olarak algılanır)
```

## 2. Kısa tanımlama: :=

- Fonksiyon içinde en sık kullanılan yöntemdir

```bash
 message := "Hello World"
 title := "Full-Stack Developer"
```

#### NOT:

:= yalnızca fonksiyon içinde kullanılabilir. Paket seviyesinde (fonksiyon dışında) var kullanmanız gerekir.

## 3. Sıfır değerler (zero values)

Değer atanmadan tanımlanan değişkenler otomatik olarak tipinin sıfır değerini alır. Go’da “tanımsız” değişken yoktur:

| Tip                                           | Sıfır değeri |
| --------------------------------------------- | ------------ |
| `int`, `float64`                              | `0`          |
| `string`                                      | `""`         |
| `bool`                                        | `false`      |
| pointer, slice, map, channel, func, interface | `nil`        |

### Örnekler

```bash
 var message string -> ""

 var price float64 -> 0

 var count int -> 0

 var isActive bool -> false
```

## 4. Çoklu tanımlama

```bash
 var a,b,c int = 1,2,3

 var (
	user = "admin"
	port = 4200
	debug = false
 )
```

Değer takası da tek satırda yapılabilir:

```bash
 a,b = b,a
```

## 5. Temel tipler

- **Tam sayılar:** int, int8, int16, int32, int64, işaretsizler için uint, uint8 (byte) vb.
- **Ondalıklı:** float32, float64
- **Metin:** string, tek karakter için rune (int32 takma adı)
- **Mantıksal:** bool
- **Karmaşık sayı:** complex64, complex128

## 6. Tip dönüşümü

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

### 6.1 Float → int: kesme (yuvarlama yok)

Ondalıklı sayı tam sayıya çevrilirken virgülden sonrası **atılır**, yuvarlanmaz:

```bash
f := 3.99
n := -3.99

int(f) --> 3
int(n) --> -3

int(math.Round(f)) --> 4 // yuvarlamak için math paketi kullanılır
```

### 6.2 Taşma (overflow)

Değer, hedef tipe sığmıyorsa Go hata vermez, sayı **başa sarar**:

```bash
big := 300
int8(big) --> 44 // int8 aralığı: -128 ile 127

neg := -1
uint8(neg) --> 255 // uint8 negatif olamaz
```

#### NOT:

Sabit (const) bir değerle yapılırsa derleme hatası alırsınız: `int8(300)` --> "constant 300 overflows int8"

### 6.3 String ↔ sayı: strconv paketi

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

### 6.4 Sık yapılan hata: string(sayı)

`string()` bir sayıyı metne çevirmez, onu **karakter kodu** olarak görür:

```bash
x := 65

string(rune(x))  --> "A"  // 65, "A" harfinin kodu
strconv.Itoa(x)  --> "65" // doğru kullanım
```

### 6.5 String ↔ []byte ve []rune

```bash
[]byte("Go") --> [71 111]

s := "çay"
len(s)          --> 4 // byte sayısı (ç 2 byte yer kaplar)
len([]rune(s))  --> 3 // karakter sayısı
```

Türkçe veya Rusça karakter içeren metinlerde karakter saymak için **[]rune** kullanın.

### 6.6 Tipsiz sabitler (untyped constants)

Tipi belirtilmemiş sabitler, kullanıldıkları yere göre tip alır, dönüşüm gerekmez:

```bash
const c = 10

var cf float64 = c --> 10 // dönüşüm gerekmedi
```

### Özet

| Dönüşüm              | Nasıl yapılır                         |
| -------------------- | ------------------------------------- |
| sayı → sayı          | `float64(x)`, `int(x)`, `int64(x)`    |
| sayı → string        | `strconv.Itoa`, `strconv.FormatInt`   |
| string → sayı        | `strconv.Atoi`, `strconv.ParseInt`    |
| string → bool        | `strconv.ParseBool`                   |
| string ↔ byte / rune | `[]byte(s)`, `[]rune(s)`, `string(b)` |

## 7. Sabitler: const

```bash
const Pi = 3.14159

const (
    Pazartesi = iota  // 0
    Sali              // 1
    Carsamba          // 2
)
```

**iota**, her satırda bir artan sayaçtır; enum benzeri yapılar için idealdir.

## 8. Dikkat edilmesi gerekenler

**Kullanılmayan değişken derleme hatasıdır.** Go bunu bilerek yapar, kodu temiz tutmak için:

```bash
x := 5   // Eğer x değeri hiçbir yerde kullanılmıyorsa: "declared and not used" hatası verir
```

### NOT:

Bir değeri bilerek yok saymak için \_ (blank identifier) kullanılır:

```bash
_, err := strconv.Atoi("123") // dönüştürülen sayı kullanılmıyor, sadece hata kontrol ediliyor
```

- **:=** en az bir yeni değişken ister. Soldaki değişkenlerin hepsi zaten varsa hata verir; en az biri yeniyse diğerlerine sadece atama yapar.

- **Gölgeleme (shadowing):** İç blokta := kullanmak dıştaki değişkeni değiştirmez, yeni bir değişken oluşturur:

```bash
x := 1

if true {
	x := 2 // yeni bir x, dıştakini etkilemez
	fmt.Println(x) // buradaki değeri 2
}

// buradaki değeri ise 1, sebebi de iç blokta **:=** kullanmak üst kısımdaki değişkeni değiştirmez sadece yeni değişken oluşturur
fmt.Println(x)
```

## 9. Görünürlük (Exported / Unexported)

Paket seviyesindeki değişkenlerde büyük harf önemlidir:

```bash
var Version string = "v1" // paket içinde ve dışında da kullanılabilir

var newVersion string = "v2" // sadece paket içinde kullanılabilir, dışarıdaki erişemez
```

## Go Format Belirteçleri

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

## Kısaca Özet

- Fonksiyon içinde çoğunlukla **:=**, fonksiyon dışında **var** kullanın.
- Tipi belirtmek istediğinizde veya sıfır değerle başlatacağınızda var tercih edin.
- Tip dönüşümleri her zaman açıktır.
- Kullanılmayan değişkenler derlenmez; gerekirse \_ kullanın.
