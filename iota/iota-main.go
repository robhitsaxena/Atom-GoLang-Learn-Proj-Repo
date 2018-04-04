package main

import "fmt"

const (
	A = iota // 0
	B = iota // 1
	C = iota // 3

)
const (
	_ = iota      // 0 - Throwing away the 0 by using the blank  identifier
	D = iota * 10 // 1 * 10
	E = iota * 10 // 2 * 10
)

const (
	_  = iota             //0
	KB = 1 << (iota * 10) // 1 << (1*10)
	MB = 1 << (iota * 10) // 1 << (2*10)
	GB = 1 << (iota * 10) // 1 << (3*10)
)

//So n << x is "n times 2, x times". And y >> z is "y divided by 2, z times".
var KB_chk = 1 << 10 // This evaluates to (1*10)-- 2 times
var MB_chk = 1 << 20 // This evaluates to (1*10)-- 2 times
var GB_chk = 1 << 30 // This evaluates to (1*10)-- 2 times
var chk1 = 32 >> 5   // This evaluates to (32/2) -- 5 times

func main() {
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
	fmt.Println(D)
	fmt.Println(E)
	fmt.Println(KB)
	fmt.Println(KB_chk)
	fmt.Println(MB)
	fmt.Println(MB_chk)
	fmt.Println(GB)
	fmt.Println(GB_chk)
	fmt.Println(chk1)

}
