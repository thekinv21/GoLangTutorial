🌐 [English](./ROADMAP_EN.md) | **Türkçe** | [Русский](./ROADMAP_RU.md)

# Go (Golang) Öğrenme Yol Haritası

## Aşama 0 — Kurulum ve Giriş

- [x] Go nedir? ([README_TR.md](../01_installation/docs/README_TR.md))
- [ ] Go kurulumu
- [ ] VS Code + Go eklentisi
- [ ] İlk program: `Hello, World!`
- [ ] `go run`, `go build`, `go fmt`, `go vet`
- [ ] `go mod init` ve `go.mod`
- [ ] `package main` ve `func main()`

**Alıştırmalar:**

- Ekrana adını ve yaşını yazdıran bir program yapabilirsin.
- Aynı programı `go build` ile derleyip oluşan dosyayı çalıştırabilirsin.

## Aşama 1 — Temeller

- [ ] Değişkenler
- [ ] Sabitler
- [ ] Temel tipler
- [ ] Tip dönüşümü
- [ ] Operatörler
- [ ] `fmt` paketi
- [ ] Kullanıcıdan girdi alma
- [ ] Koşullar: `if / else`, `switch`
- [ ] Döngüler: `for`

**Alıştırmalar:**

- Komut satırında çalışan bir hesap makinesi yapabilirsin.
- Sayı tahmin oyunu yapabilirsin.
- 1'den 100'e kadar FizzBuzz yazdırabilirsin.

## Aşama 2 — Fonksiyonlar ve Veri Yapıları

- [ ] Fonksiyonlar
- [ ] Birden fazla dönüş değeri
- [ ] Variadic fonksiyonlar
- [ ] Anonim fonksiyonlar ve closure
- [ ] `defer`
- [ ] Array
- [ ] Slice
- [ ] Map
- [ ] String işlemleri
- [ ] `strconv` paketi

**Alıştırmalar:**

- Bir metindeki kelimelerin kaç kez geçtiğini sayan bir program yapabilirsin.
- Not ortalaması hesaplayıcı yapabilirsin.

## Aşama 3 — Pointer, Struct ve Method

- [ ] Pointer
- [ ] Değer ile geçirme vs pointer ile geçirme
- [ ] Struct
- [ ] Struct embedding
- [ ] Method
- [ ] Constructor deseni
- [ ] Struct tag'leri

**Alıştırmalar:**

- Bellekte çalışan bir kütüphane veya öğrenci yönetim sistemi yapabilirsin.

## Aşama 4 — Interface ve Hata Yönetimi

- [ ] Interface
- [ ] `any` / `interface{}`
- [ ] Type assertion ve type switch
- [ ] Yaygın interface'ler
- [ ] Hata yönetimi
- [ ] Özel hata tipleri
- [ ] `errors.Is`, `errors.As`
- [ ] `panic` ve `recover`

**Alıştırmalar:**

- `Shape` interface'i ile şekillerin alanını hesaplayan bir program yapabilirsin.
- Hatalı işlemlerde error dönen bir banka hesabı yapabilirsin.

## Aşama 5 — Paketler, Modüller ve Proje Yapısı

- [ ] Kendi paketini yazma
- [ ] Dış paket ekleme
- [ ] Proje klasör yapısı
- [ ] `init()` fonksiyonu

**Alıştırmalar:**

- Önceki projelerinden birini paketlere bölebilirsin.

## Aşama 6 — Generics

- [ ] Generic fonksiyonlar
- [ ] Type constraint'ler
- [ ] Generic struct'lar
- [ ] `slices` ve `maps` paketleri

**Alıştırmalar:**

- Generic bir Stack veya Queue yazabilirsin.

## Aşama 7 — Concurrency

- [ ] Goroutine
- [ ] `sync.WaitGroup`
- [ ] Channel
- [ ] `select`
- [ ] Mutex
- [ ] Race condition
- [ ] `context` paketi
- [ ] Concurrency desenleri

**Alıştırmalar:**

- Bir URL listesini aynı anda kontrol edip durum kodlarını yazdıran bir araç yapabilirsin.
- Worker pool ile dosya işleyen bir program yapabilirsin.

## Aşama 8 — Test ve Kod Kalitesi

- [ ] `testing` paketi
- [ ] Table-driven testler
- [ ] Benchmark
- [ ] Test coverage
- [ ] Mock'lama
- [ ] `golangci-lint`

**Alıştırmalar:**

- Önceki projelerine test yazabilirsin.

## Aşama 9 — Standart Kütüphane

- [ ] Dosya okuma / yazma
- [ ] JSON
- [ ] `time` paketi
- [ ] Komut satırı argümanları
- [ ] Loglama
- [ ] Ortam değişkenleri

**Alıştırmalar:**

- Görevleri JSON dosyasında saklayan bir CLI Todo uygulaması yapabilirsin.

## Aşama 10 — Web ve Backend

- [ ] HTTP sunucusu
- [ ] Routing
- [ ] JSON API
- [ ] Middleware
- [ ] HTTP client
- [ ] Framework'ler
- [ ] Veritabanı
- [ ] ORM / query builder
- [ ] Migration
- [ ] Config yönetimi
- [ ] Authentication
- [ ] Validation

**Alıştırmalar:**

- PostgreSQL ve JWT kullanan bir Blog veya Todo REST API'si yapabilirsin.

## Aşama 11 — İleri Seviye ve Yayına Alma

- [ ] Clean / Hexagonal Architecture
- [ ] Graceful shutdown
- [ ] Docker
- [ ] Docker Compose
- [ ] gRPC
- [ ] WebSocket
- [ ] Redis
- [ ] Message queue
- [ ] Profiling
- [ ] CI/CD

**Alıştırmalar:**

- Mikroservis tabanlı bir URL kısaltıcı, chat uygulaması veya e-ticaret backend'i yapabilirsin.
