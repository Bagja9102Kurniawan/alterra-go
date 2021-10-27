/*
	Bagja 9102 Kurniawan
*/
package main
import "fmt"

var h, w, s int

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

func primaSegiEmpat (high, wide, start int) {
	var prime []int
	var j, sum int = 0, 0
	for i := start + 1; len(prime) < high * wide; i++ {
		if primeNumber(i) == true {
			prime = append(prime, i)
		}
	}
	for i := 0; i < len(prime); i++ {
		j = len(prime) / wide
		fmt.Print(prime[i], " ")
		if (i+1) % j == 0 {
			fmt.Println("")
		}
		sum += prime[i]
	}
	fmt.Println(sum)
}

func main() {
	fmt.Print("Height : ")
	fmt.Scanf("%d\n", &h)
	fmt.Print("Width : ")
	fmt.Scanf("%d\n", &w)
	fmt.Print("Start From : ")
	fmt.Scanf("%d\n", &s)

	primaSegiEmpat(h, w, s)
}