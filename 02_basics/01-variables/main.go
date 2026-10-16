package main

import (
	"fmt"
	"strconv"
)

// Paket seviyesinde := kullanılamaz, var kullanılmalı
var Version string = "v1"    // büyük harf: paket dışından erişilebilir (exported)
var newVersion string = "v2" // küçük harf: sadece paket içinden erişilebilir (unexported)

func main() {

	//! Var ile değişken tanımlama örnekleri

	var count int = 20
	var fullName string = "Vadim"
	var isActive bool = false
	var price float64 = 3.434

	fmt.Println(
		"Count:", count,
		"fullName:", fullName,
		"isActive:", isActive,
		"price:", price,
	)

	//! Kısa tanımlama ile değişken tanımlama

	// Not:

	// - Bunu sadece fonksiyon içinde yapınız
	// - Fonksiyon dışı ve paket kullanımlarında var kullanmak daha iyi seçenektir

	message := "Hello World"
	developer := "Developer"

	fmt.Printf("Message: %s\nDeveloper: %s\n", message, developer)

	//! Sıfır değeri konusu, eğer default değeri olmazsa ne atanacağını belirleyecek

	var myMessage string

	var myPrice float64

	var myCount int

	var isActiveValue bool

	fmt.Printf(
		"MyMessage: %s\nMyPrice: %.2f\nMyCount: %d\nIsActiveValue: %t\n", myMessage, myPrice, myCount, isActiveValue)

	//! Çoklu değişken atama konusu

	var a, b, c int = 1, 2, 4

	var (
		user  = "admin"
		port  = 4200
		debug = false
	)

	fmt.Printf("A: %d\nB: %d\nC: %d\n", a, b, c)

	fmt.Printf("User: %s\nPort: %d\nDebug: %t\n", user, port, debug)

	//! Değer takası tek satırda yapılabilir

	a, b = b, a

	fmt.Printf("Takas sonrası A: %d, B: %d\n", a, b)

	//! _ (blank identifier) ile değeri bilerek yok sayma

	_, err := strconv.Atoi("123")

	if err != nil {
		fmt.Println("Error:", err)
	}

	//! Gölgeleme (shadowing)

	x := 1

	if true {
		x := 2 // yeni bir x, dıştakini etkilemez
		fmt.Println("İç blok x:", x)
	}

	fmt.Println("Dış blok x:", x)

	fmt.Println("Version:", Version, "newVersion:", newVersion)

	//NOT:

	//! Eğer projenizde kullanılmayan değerler veya değişkenler varsa go derleme anında hata verir, go bunu bilerek yapıyor, kod kısmını temiz tutmak için

	//! go run ---> codebase kısmı çalıştırıyor
	//! go vet ---> hataları bulmada yardım ediyor
	//! go fmt ---> kod formatlamak için kullanılıyor
}
