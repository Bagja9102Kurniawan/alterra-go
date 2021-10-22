/*
	Bagja 9102 Kurniawan
	
	mencek bilangan prima
	menggunakan algo primality test dengan optimisasi 6k+i
	merujuk algoritma sieve of eratosthenes
	kompleksitas O(√n)
*/
package main
import "fmt"

var n int

func primeNumber(number int) bool{
	if number <= 3{
		return number > 1
	}
	if number%2==0 || number %3 ==0{
		return false
	}
	for i:=5; i*i<=number; i += 6{
		if number % i == 0 || number%(i+2)==0{
			return false
		}
	}
	return true
}

func main() {
	fmt.Print("Bilangan= ")
	fmt.Scanf("%d",&n)
	
	if primeNumber(n){
		fmt.Println("Bilangan Prima")
	} else {
		fmt.Println("Bukan Bilangan Prima")
	}
}