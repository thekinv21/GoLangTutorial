package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {

	//! Sayı → sayı dönüşümü: T(değer)

	var intNumber int = 20
	var floatNumber float64 = 2.5

	total := float64(intNumber) + floatNumber

	fmt.Printf("Total: %.2f\n", total)

	//! Float → int: kesme (yuvarlama yok)

	f := 3.99

	fmt.Println("int(f):", int(f), "int(math.Round(f)):", int(math.Round(f)))

	//! Taşma (overflow): sayı başa sarar

	big := 300
	neg := -1

	fmt.Println("int8(300):", int8(big), "uint8(-1):", uint8(neg))

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

	//! Hatalı dönüşüm: number değeri 0 olur

	number, err := strconv.Atoi(myString)

	if err != nil {
		fmt.Println("Error:", err, "number:", number)
	}

	//! Sık yapılan hata: string(sayı)

	x := 65

	fmt.Println("string(rune(65)):", string(rune(x)), "strconv.Itoa(65):", strconv.Itoa(x))

	//! String ↔ []byte ve []rune

	s := "çay"

	fmt.Println("[]byte(\"Go\"):", []byte("Go"), "len:", len(s), "rune sayısı:", len([]rune(s)))

	//! Sayıyı string içine yerleştirme

	fmt.Println(fmt.Sprintf("%d-%.1f", 7, 2.5))
}
