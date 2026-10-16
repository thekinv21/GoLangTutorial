package main

import "fmt"

func main() {

	//! Aritmetik operatörlere örnek:

	a, b := 7, 2

	fmt.Println("a / b: ", a/b) // Tam sayı bölme sonucu hep tam sayı çıkacaktır: --> Sonuç: 3

	fmt.Println("7.0 / 2 : ", 7.0/2) // Ondalıklı sayı ve Int sayı bölme işleminde ilk önce sol tarafa bakarak baz alınır --> Sonuç: 3.5

	fmt.Println("float64(a) / float64(b) : ", float64(a)/float64(b)) // Eğer iki değerin tipi aynı ise sonuçlar hep doğru olacaktır --> Sonuç: 3.5

	fmt.Println("-7 / 2 : ", -7/2) // Burada eğer bölen değer 0'dan küçük ise sonuç 0'a doğru yuvarlar --> Sonuç: -3

	fmt.Println("-12 % 5 : ", -12%5) // Mod alırken işaret hep sol taraftaki değerden gelir --> Sonuç: -2

	var xValue int = 12
	var yValue float64 = 2.0

	fmt.Println("xValue + yValue : ", xValue+int(yValue)) // Farklı tiplerdeki değerler bir biri ile toplanamaz,bölünemez veya çıkarılamaz, Go direkt hata verir --> Sonuç: mismatched types hatası verir, aynı tipe çevirmezsek


	//! Karşılaştırma operatörlere örnek:

	m,n := 7,2

	fmt.Println(
		m == n, 
		m != n, 
		m>n , 
		m<n, 
		m>=n, 
		m<=n,
	);

	// Sonuç: false,true,true,false,true,false


	//! Kısa devre: && solu false ise sağa bakmaz

	var p *int

	if p != nil && *p > 0 {
		fmt.Println("Pozitif")
	}

}
