package main

import "fmt"

const Pi = 3.14159

//! iota: her satırda bir artan sayaç, enum benzeri yapılar için

const (
	Pazartesi = iota // 0
	Sali             // 1
	Carsamba         // 2
)

func main() {

	fmt.Println("Pi:", Pi)

	fmt.Println("Pazartesi:", Pazartesi, "Sali:", Sali, "Carsamba:", Carsamba)

	//! Tipsiz sabitler kullanıldıkları yere göre tip alır

	const c = 10

	var cf float64 = c

	fmt.Printf("cf: %v (%T)\n", cf, cf)
}
