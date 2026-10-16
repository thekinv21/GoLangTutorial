package main

import (
	"fmt"
	"strconv"
)

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

	//! Tip dönüştürmeler ile ilgili konular

	// NOT:
	// - Parse... ve strconv.Atoi hata fırlatıyor

	var myNum string = "123"
	var myString string = "abc"
	var myNumber int = 435

	var integerNum int = 343

	convertedMyNum, err := strconv.Atoi(myNum)
	convertedMyNumber := strconv.Itoa(myNumber)

	convertedToFloatNumber := float64(integerNum)

	if err != nil {
		fmt.Println("Error when converting value:", err)
		return
	}

	fmt.Printf(
		"Converted Value: %d\nMyString: %q\nConverted to String Value: %s\nconvertedToFloatNumber: %.2f\n", convertedMyNum, myString, convertedMyNumber, convertedToFloatNumber)

	//! Ekrana yazdırma konuları:

	// %s ---> string için kullanılıyor

	// %d ---> integer için kullanılıyor

	// %f ---> float için kullanılır

	// %.2f ---> float için kullanılıyor yani sadece 2 basamak gösterilsin istersek nokta sonrası

	// %t ---> bool değerleri için kullanılıyor

	// %v ---> herhangi bir tip için kullanılıyor

	// %+v ---> struct alanlar için kullanılıyor yani objeler için

	// %T ---> verinin tipini öğrenmek için kullanılabilir

	// %% ---> % işareti için

	//NOT:

	//! Eğer projenizde kullanılmayan değerler veya değişkenler varsa go derleme anında hata verir, go bunu bilerek yapıyor, kod kısmını temiz tutmak için

	//! go run ---> codebase kısmı çalıştırıyor
	//! go vet ---> hataları bulmada yardım ediyor
	//! go fmt ---> kod formatlamak için kullanılıyor
}
