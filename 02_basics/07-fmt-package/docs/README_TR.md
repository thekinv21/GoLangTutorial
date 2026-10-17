[English](./README_EN.md) | **Turkish** | [Russian](./README_RU.md)

## `fmt` paketi nedir?

fmt (format) **Giriş/Çıkış (I/O)** işlemlerini yapan pakettir. C dilindeki printf ve scanf fonksiyonlarına benzer ama Go'nun tip sistemine göre tasarlanmıştır. Kullanmak için paketi içe aktarmak yeterlidir

```bash
import "fmt"
```

Fonksiyonlar adlarına göre gruplara ayrılır:

| Önek / Sonek                        | Anlamı                                                  |
| ----------------------------------- | ------------------------------------------------------- |
| `Print...`                          | Standart çıktıya (`os.Stdout`) yazar                    |
| `Sprint...`                         | Sonucu `string` olarak döndürür                         |
| `Fprint...`                         | Herhangi bir `io.Writer`'a yazar                        |
| `Append...`                         | Sonucu bir `[]byte` dilimine ekler (Go 1.19+)           |
| `Scan...` / `Sscan...` / `Fscan...` | Sırasıyla stdin'den, string'den ve `io.Reader`'dan okur |
| `...ln`                             | Sona yeni satır ekler (okumada satır sonunda durur)     |
| `...f`                              | Biçim dizesi (format string) kullanır                   |

## Ekrana yazdırma: **`Print`, `Println`, `Printf`**

```bash
fmt.Print("Hello","World") --> Sonuç: HelloWorld

fmt.Print(1,2) --> Sonuç: 1 2

fmt.Println("Hello","World") --> Sonuç: Hello World\n

fmt.Printf("Age: %d\n", 20) --> Sonuç: Age: 20
```

NOT:

- **Print** - değerleri yan yana yazar,araya boşluğu ise sadece 2 değer de **string değilse** ekler ve sona yeni satır yani **\n** eklemez
- **Println** - değerlerin arasına her zaman boşluk ekler ve sona yeni satır ekler yani **\n**
- **Printf** - bir biçim dizesi alır, yani yeni eklemez, onu manuel biz eklemeliyiz

Üçü de **(n int, err error)** döndürür: yazılan bayt sayısı ve olası hata. Ekrana yazarken bu değerler genelde göz ardı edilir.

### Genel biçim (verbs)

| Fiil  | Açıklama                    | Örnek çıktı                   |
| ----- | --------------------------- | ----------------------------- |
| `%v`  | Varsayılan biçim            | `{Ali 30}`                    |
| `%+v` | Struct'larda alan adlarıyla | `{Ad:Ali Yas:30}`             |
| `%#v` | Go sözdizimiyle gösterim    | `main.Kisi{Ad:"Ali", Yas:30}` |
| `%T`  | Değerin tipi                | `main.Kisi`                   |
| `%%`  | Yüzde işaretinin kendisi    | `%`                           |

### Integer değeler için (verbs)

| Fiil        | Açıklama                       | `42` için   |
| ----------- | ------------------------------ | ----------- |
| `%d`        | Onluk taban                    | `42`        |
| `%b`        | İkilik taban                   | `101010`    |
| `%o`        | Sekizlik taban                 | `52`        |
| `%O`        | `0o` önekli sekizlik           | `0o52`      |
| `%x` / `%X` | Onaltılık (küçük / büyük harf) | `2a` / `2A` |
| `%c`        | Unicode karakter               | `*`         |
| `%q`        | Tek tırnaklı karakter          | `'*'`       |
| `%U`        | Unicode biçimi                 | `U+002A`    |

### Float değeler için (verbs)

| Fiil        | Açıklama                                              | `3.14159` için          |
| ----------- | ----------------------------------------------------- | ----------------------- |
| `%f` / `%F` | Ondalık gösterim                                      | `3.141590`              |
| `%e` / `%E` | Bilimsel gösterim                                     | `3.141590e+00`          |
| `%g` / `%G` | Duruma göre `%e` veya `%f`, gereksiz sıfırlar olmadan | `3.14159`               |
| `%x`        | Onaltılık bilimsel gösterim                           | `0x1.921f9f01b866ep+01` |

### String ve byte (verbs)

| Fiil | Açıklama                        | `"Go"` için |
| ---- | ------------------------------- | ----------- |
| `%s` | Ham string                      | `Go`        |
| `%q` | Çift tırnaklı, kaçış karakterli | `"Go"`      |
| `%x` | Bayt başına onaltılık           | `476f`      |

### Boolean ve işaretçi (verbs)

| Fiil | Açıklama      | Örnek çıktı      |
| ---- | ------------- | ---------------- |
| `%t` | Boolean değer | `true` / `false` |
| `%p` | Bellek adresi | `0xc000012345`   |

#### İşaretçiler: %p → 0xc000012345 gibi bir bellek adresi.

```bash
type Person struct {
	Name string
	Age int
}

myPerson := Person {"Vadim",26}

fmt.Printf("%v\n",myPerson) --> Sonuç: {Vadim 26}

fmt.Printf("%+v\n",myPerson) --> Sonuç: {Name:Vadim Age:26}

fmt.Printf("%#v\n",myPerson) --> Sonuç: main.Person{Name:"Vadim", Age:26}

fmt.Printf("%T\n",myPerson) --> Sonuç: main.Person

```

## String üretme: `Sprint`, `Sprintln`, `Sprintf`

Bu fonksiyonlar da Print ailesi gibi çalışır, ama ekrana yazdırmak yerine sonucu **string** olarak döndürür

```bash
stringValue := fmt.Sprintf("%s %d yaşında", "Vadim",26) --> Sonuç: "Vadim 26 yaşında"

totalValue := fmt.Sprint("Total:",10) --> Sonuç: "Total:10"

uValue := fmt.Sprintln("a", "b") --> Sonuç: "a b\n"
```

Log mesajı, dosya adı, SQL dışı metin şablonları gibi string oluşturmanız gereken her yerde kullanılır

**Append** ailesi (Go 1.19+): Mevcut bir []byte dilimine ekleme yapar. Yeni string ayırmadığı için performanslıdır:

```bash
buf := []byte("Total: ")
buf = fmt.Appendf(buf, "%d",42) -->  Sonuç: "Total: 42"
```

## Farklı hedefe yazma: `Fprint`, `Fprintln`, `Fprintf`

iki parametre olarak bir **io.Writer** alırlar. **Write** metodu olan her şeye (dosya,ağ bağlantısı,HTTP response,buffer) yazabiliriz:

```bash
// Standart hata akışına yazma

fmt.Fprintln(os.Stderr, "bir hata oluştu")

// Dosyaya yazma işlemi:

newFile, err := os.Create("newFile.txt")
if err != nil {
	log.Fatal(err)
}
defer newFile.Close()

fmt.Fprintf(newFile, "Row %d\n",1)

// HTTP response'a yazma

func handler(w http.ResponseWriter,r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}


// Buffer'a yazma

var newBuffer strings.Builder
fmt.Fprintf(&newBuffer, "%d-%d",1,2)
```

Aslında **Printf** **Fprintf(os.Stdout, ...)** kısmının kısa yoludur

## Hata oluşturma: `Errorf`

**Errorf** biçimlendirilmiş bir mesajla error değeri üretir:

```bash
err := fmt.Errorf("Kullanıcı %d bulunamadı",id)
```

**%w** ile hata sarmalama (Go 1.13+): Bu fiil, alttaki hatayı saklar. Böylece **errors.Is** ve **errors.As** ile asıl hataya ulaşabilirsiniz:

```bash
_,err := os.Open("yok.text")

if err != nil {
	err = fmt.Errorf("Yapılandırma okunamadı: %w", err)
}

if errors.Is(err,os.ErrNotExist) {
	fmt.Println("Dosya bulunamadı") // çalışır
}
```

- **%v** kullanırsanız mesaj aynı görünür, ama zincir kopar ve errors.Is artık eşleşme bulamaz.
- Go 1.20'den itibaren tek bir Errorf içinde birden fazla **%w** kullanılabilir
- Sabit bir mesajlar için **errors.New("...")** daha basit bir seçenektir

## Kullanıcıdan veri okuma: `Scan`, `Scanln`, `Scanf`

Bu fonksiyonlar değerleri işaretçiler aracılığıyla döndürür, ayrıca okunan değer sayısını ve bir hata döndürür

```bash
var Name string
var Age int

// Scan: boşluk ve yeni satırları ayırıcı sayar
n,err := fmt.Scan(&Name,&Age)

// Scanln: yeni satırda durur; tüm değerler aynı satırda olmalıdır
fmt.Scanln(&Name,&Age)

// Scanf: Belirli bir biçime göre okur
fmt.Scanf("%s %d",&Name,&Age)
```

### Okuma fonksiyonlarında yeni satır davranışı

| Fonksiyon | Yeni satır davranışı                                                      |
| --------- | ------------------------------------------------------------------------- |
| `Scan`    | Yeni satırı boşluk gibi kabul eder, tüm değerler okunana kadar devam eder |
| `Scanln`  | Satır sonunda durur. Eksik değer varsa hata döndürür                      |
| `Scanf`   | Girdideki yeni satırlar biçim dizesindekilerle eşleşmelidir               |

**Önemli:** Scan ailesi kelimeleri boşluklara göre ayırır, yani **"Ali Veli"** tek bir string'e okunamaz. Boşluklu satırları okumak için **bufio** kullanılmalıdır

```bash
reader := bufio.NewReader(os.Stdin)
satir,_ := reader.ReadString('\n')
satir = strings.TrimSpace(satir)

// ya da satır satır okuyalım:

scanner := bufio.NewScanner(os.Stdin)

for scanner.Scan() {
	fmt.Println(scanner.Text())
}
```

Bir **io.Reader** dan okumak için **Fscan**,**Fscanln**,**Fscanf** de vardır

## String'den okuma: `Sscan`, `Sscanf`

Girdi olarak stdin yerine bir string alırlar

```bash
var x,y int

fmt.Sscan("10 20", &x,&y)

var day,month,year int
fmt.Sscanf("17/10/26", "%d/%d/%d", &day,&month,&year)
```

## Özet

**`fmt` paketi**, Go'da ekrana yazdırma, string oluşturma ve kullanıcıdan veri okuma işlerini yapan standart pakettir.

### Fonksiyon adlarından ne yaptıklarını anlayabilirsiniz

- `Print` ekrana yazar, `Sprint` string döndürür, `Fprint` dosya gibi başka bir hedefe yazar.
- Sonu `ln` ile bitenler satır sonu ekler, `f` ile bitenler biçim dizesi (`%d`, `%s` gibi) kullanır.
- `Scan` ailesi kullanıcıdan, `Sscan` ailesi bir string'den veri okur.
- `Errorf` biçimlendirilmiş hata mesajı üretir. `%w` kullanılırsa alttaki hata korunur.

### En çok kullanılan fiiller

- `%v` her türlü değer, `%+v` struct alan adlarıyla, `%T` tip adı
- `%d` tam sayı, `%f` ondalık sayı, `%s` string, `%t` boolean
- `%x` onaltılık, `%q` tırnaklı gösterim

### Biçimlendirme

`%8d` genişlik verir, `%-8d` sola yaslar, `%05d` sıfırla doldurur, `%.2f` iki ondalık basamak gösterir.

### Stringer

Bir tipe `String() string` metodu yazarsanız, `fmt` o tipi yazdırırken bu metodu kullanır.

### Dikkat edilecekler

- `Printf` kendiliğinden yeni satır eklemez, `\n` yazmayı unutmayın.
- Okuma yaparken değişkenin adresini verin: `fmt.Scan(&x)`.
- `Scan` boşlukta durur. Boşluklu bir satırı okumak için `bufio` kullanın.
- `String()` metodunun içinde aynı değeri `%v` ile yazdırmayın, program sonsuz döngüye girer.
- Biçim hatalarını yakalamak için `go vet` çalıştırın.
