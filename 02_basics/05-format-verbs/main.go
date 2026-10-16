package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {

	u := User{Name: "Ali", Age: 30}

	//! Ekrana yazdırma konuları:

	fmt.Printf("%s\n", "Hello") // %s ---> string için kullanılıyor
	fmt.Printf("%q\n", "Hello") // %q ---> string'i tırnaklı yazdırır
	fmt.Printf("%d\n", 42)      // %d ---> integer için kullanılıyor
	fmt.Printf("%f\n", 19.99)   // %f ---> float için kullanılır
	fmt.Printf("%.2f\n", 19.99) // %.2f ---> nokta sonrası sadece 2 basamak
	fmt.Printf("%t\n", true)    // %t ---> bool değerleri için kullanılıyor
	fmt.Printf("%v\n", 42)      // %v ---> herhangi bir tip için kullanılıyor
	fmt.Printf("%+v\n", u)      // %+v ---> struct alanlar için kullanılıyor yani objeler için
	fmt.Printf("%T\n", 3.14)    // %T ---> verinin tipini öğrenmek için kullanılabilir
	fmt.Printf("%%\n")          // %% ---> % işareti için
	fmt.Printf("a\nb\n")        // \n ---> yeni satır
}
