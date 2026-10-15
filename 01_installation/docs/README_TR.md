[English](./README_EN.md) | **Turkish** | [Russian](./README_RU.md)

## Go (Golang) nedir?

Go, Google'ın 2009'da duyurduğu açık kaynaklı, statik tipli ve derlenen bir programlama dilidir. Hızlı, sade ve eşzamanlı işlemlerde güçlü yazılımlar geliştirmek için tasarlanmıştır.

<table>
  <tr>
    <td rowspan="2"><img src="../images/go_image_v3.jpeg" alt="go_image_v3" width="400"></td>
    <td><img src="../images/go_image_v1.jpeg" alt="go_image_v1" width="330"></td>
  </tr>
  <tr>
    <td><img src="../images/go_image_v2.jpeg" alt="go_image_v2" width="330"></td>
  </tr>
</table>

## Go Kurulumu

[go.dev/dl](https://go.dev/dl) adresinden işletim sistemine uygun dosyayı indir ve kur.

- **Windows:** `.msi` dosyasını çalıştır.
- **macOS:** `.pkg` dosyasını çalıştır veya `brew install go`
- **Linux:** `.tar.gz` dosyasını `/usr/local` içine çıkar, `/usr/local/go/bin` yolunu PATH'e ekle.

Kontrol et:

```bash
go version
```

## İlk Proje

```bash
mkdir merhaba && cd merhaba
go mod init merhaba
```

`main.go`:

```go
package main

import "fmt"

func main() {
    fmt.Println("Merhaba, Go!")
}
```

```bash
go run main.go
```

**Editör:** VS Code + resmi Go eklentisi

## Go ne işe yarar?

- **Backend ve REST API:** Kimlik doğrulama, CRUD ve veritabanı işlemleri.
- **Mikroservisler:** Büyük uygulamaları küçük, bağımsız servislere bölmek.
- **Cloud ve DevOps:** Docker, Kubernetes ve Terraform Go ile yazılmıştır; CLI araçları için idealdir.
- **Eşzamanlı sistemler:** Çok sayıda isteği ve bağlantıyı aynı anda yönetmek.

## Avantajları

- Yüksek performans
- Goroutine ile kolay eşzamanlılık
- Basit sözdizimi
- Statik tip sistemi
- Tek dosyalık kolay dağıtım
- Güçlü standart kütüphane ve araçlar

## Dezavantajları

- **Tekrarlayan hata yönetimi:** Exception yoktur; sürekli `if err != nil` yazmak gerekir.
- **Sınırlı generics:** 1.18 ile geldi ama Java ve C#'a göre daha basittir.
- **Class ve kalıtım yok:** OOP alışkanlığı olanlar için farklı bir düşünme şekli gerekir.
- **Eksik dil özellikleri:** Enum, ternary operatörü (`? :`) ve varsayılan parametre yoktur.
- **Daha küçük ekosistem:** Java ve Node.js kadar hazır kütüphane ve framework bulunmaz.
- **Garbage collector:** Gerçek zamanlı ve düşük seviyeli sistemlerde (oyun motoru, gömülü sistem) Rust veya C++ daha uygundur.
- **Frontend ve GUI için zayıf:** Masaüstü ve mobil arayüz geliştirmede tercih edilmez.
- **Nil pointer hataları:** Dikkat edilmezse çalışma anında `panic` oluşabilir.

## Go vs Java vs C# vs Node.js

| Özellik                   | Go                         | Java                                   | C#                                        | Node.js                                     |
| ------------------------- | -------------------------- | -------------------------------------- | ----------------------------------------- | ------------------------------------------- |
| **Tür**                   | Programlama dili           | Programlama dili                       | Programlama dili                          | JavaScript çalışma ortamı                   |
| **Çalışma**               | Derlenmiş makine kodu      | JVM üzerinde bytecode (JIT)            | .NET CLR üzerinde (JIT, AOT desteği var)  | V8 üzerinde JavaScript                      |
| **Performans**            | Yüksek                     | Yüksek, ısınma süresi sonrası          | Yüksek                                    | I/O işlemlerinde güçlü, CPU işlerinde zayıf |
| **Başlangıç süresi**      | Çok hızlı                  | Yavaş                                  | Orta                                      | Hızlı                                       |
| **Eşzamanlılık**          | Goroutines, channels       | Thread'ler, virtual threads (Java 21+) | async/await, Task                         | Event loop, async/await                     |
| **Bellek kullanımı**      | Düşük                      | Yüksek                                 | Orta                                      | Orta                                        |
| **Tip sistemi**           | Statik                     | Statik                                 | Statik                                    | Dinamik (TypeScript ile statik)             |
| **Sözdizimi**             | Sade, az anahtar kelime    | Uzun ve detaylı                        | Zengin, çok özellikli                     | Esnek                                       |
| **Öğrenme**               | Kolay                      | Orta–zor                               | Orta                                      | JavaScript biliyorsan kolay                 |
| **Popüler framework'ler** | net/http, Gin, Echo, Fiber | Spring Boot, Quarkus                   | ASP.NET Core                              | Express, Fastify, NestJS                    |
| **Dağıtım**               | Tek binary dosya           | JAR + JVM gerekir                      | .NET runtime veya self-contained          | Node.js + node_modules gerekir              |
| **Güçlü olduğu alan**     | Mikroservis, cloud, CLI    | Kurumsal sistemler, bankacılık         | Kurumsal sistemler, Windows, oyun (Unity) | Gerçek zamanlı uygulamalar, full-stack JS   |

**Özet:** Go sadelik ve hız sunar. Java ve C# büyük kurumsal projelerde güçlü ekosistemlere sahiptir. Node.js ise frontend ile aynı dili kullanma avantajı sağlar.
