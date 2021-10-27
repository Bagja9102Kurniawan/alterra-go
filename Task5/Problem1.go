/*
	Bagja 9102 Kurniawan
*/
package main
import "fmt"

func primeNumber(number int) bool {
	if number <= 3 {
		return number > 1
	}
	if number % 2 == 0 || number % 3 == 0 {
		return false
	}
	for i := 5; i * i <= number; i += 6 {
		if number % i == 0 || number % (i + 2) == 0 {
			return false
		}
	}
	return true
}

func getPrime(number int) int {
	var n, prime int = 1, 2
	for i := 3; n < number; i++ {
		if primeNumber(i) {
			n++
			prime = i
		}
	}
	return prime
}

func main() {
	fmt.Println("Bilangan prima pertama : ", getPrime(1))
	fmt.Println("Bilangan prima ke-5 : ", getPrime(5))
	fmt.Println("Bilangan prima ke-8 : ", getPrime(8))
	fmt.Println("Bilangan prima ke-9 : ", getPrime(9))
	fmt.Println("Bilangan prima ke-10 : ", getPrime(10))
}