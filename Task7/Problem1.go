/*
	Bagja 9102 Kurniawan
*/
package main
import "fmt"

var n int

func fibo(n int, F []int) int {
	if n == 0 || n == 1 {
		F[n] = n
	} else if F[n] == 0 {
		F[n] = fibo(n-1, F) + fibo(n-2, F)
	}
	return F[n]
}

func main() {
	fmt.Print("Fibonacci ke = ")
	fmt.Scanf("%d ", &n)

	F := make([]int, n+1)
	fmt.Print(fibo(n, F))
}