/*
	Bagja 9102 Kurniawan
*/
package main
import "fmt"

var a int

func primeNumber(n int) bool{
	var b int = 0
	for i:=1; i<=n; i++{
		if n%i ==0 {
			b++
		}
	}
	if b == 2 {
		return true
	}
	return false
}

func main () {
	fmt.Print("Masukkan Angka : ")
	fmt.Scan(&a)
	if primeNumber(a) {
		fmt.Println(a,"Adalah Bilangan Prima")
	} else {
		fmt.Println(a,"Bukan Bilangan Prima")
		}
}