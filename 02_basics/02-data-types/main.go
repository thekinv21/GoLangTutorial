package main

import (
	"fmt"
	"math"
)

func main() {

	//! Tam sayılar (işaretli)

	var small int8 = 127
	var medium int32 = 2147483647
	var large int64 = math.MaxInt64
	var number int = 42 // platforma göre 32 veya 64 bit

	fmt.Printf("int8: %d\nint32: %d\nint64: %d\nint: %d\n", small, medium, large, number)

	//! Tam sayılar (işaretsiz)

	var u8 uint8 = 255
	var u32 uint32 = math.MaxUint32

	fmt.Printf("uint8: %d\nuint32: %d\n", u8, u32)

	//! Ondalıklı sayılar

	var f32 float32 = 3.1415927
	var f64 float64 = 3.141592653589793

	fmt.Printf("float32: %v\nfloat64: %v\n", f32, f64)

	//! Karmaşık sayılar

	var c complex128 = complex(2, 3)

	fmt.Printf("complex: %v, real: %v, imag: %v\n", c, real(c), imag(c))

	//! Diğer tipler

	var isActive bool = true
	var title string = "Go"
	var b byte = 'A' // uint8 takma adı
	var r rune = 'ç' // int32 takma adı (Unicode karakter)

	fmt.Printf("bool: %t\nstring: %s\nbyte: %d\nrune: %d (%c)\n", isActive, title, b, r, r)

	//! %T ile değişkenin tipini öğrenme

	fmt.Printf("%T %T %T %T\n", number, f64, b, r)

	//! Sınırlar için math paketi kullanılabilir

	fmt.Println("MaxInt8:", math.MaxInt8, "MinInt64:", math.MinInt64, "MaxFloat64:", math.MaxFloat64)
}
