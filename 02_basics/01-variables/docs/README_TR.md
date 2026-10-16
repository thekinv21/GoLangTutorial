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

## 5. Dikkat edilmesi gerekenler

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

## 6. Görünürlük (Exported / Unexported)

Paket seviyesindeki değişkenlerde büyük harf önemlidir:

```bash
var Version string = "v1" // paket içinde ve dışında da kullanılabilir

var newVersion string = "v2" // sadece paket içinde kullanılabilir, dışarıdaki erişemez
```

## Kısaca Özet

- Fonksiyon içinde çoğunlukla **:=**, fonksiyon dışında **var** kullanın.
- Tipi belirtmek istediğinizde veya sıfır değerle başlatacağınızda var tercih edin.
- Tip dönüşümleri her zaman açıktır.
- Kullanılmayan değişkenler derlenmez; gerekirse \_ kullanın.
