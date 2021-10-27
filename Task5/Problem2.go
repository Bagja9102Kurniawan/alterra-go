/*
	Bagja 9102 Kurniawan
*/
package main
import "fmt"

var n int

func fibonacci(number int) int {
	if number == 0 || number == 1 {
		return number
	}
	return fibonacci(number - 2) + fibonacci(number - 1)
}

func main() {
	fmt.Print("Fibonacci ke : ")
	fmt.Scanf("%d\n", &n)
	fmt.Println("Fibonacci ke-", n, " : ", fibonacci(n))
}