/*
	Bagja 9102 Kurniawan
*/
package main
import "fmt"

var n int

func fibo(n int) int{
	var F []int
	F = append(F, 0)
	F = append(F, 1)

	for i := 2; i <= n; i++ {
		F = append(F, F[i-1] + F[i-2])
	}
	return F[n]
}

func main() {
	fmt.Print("Fibonacci ke = ")
	fmt.Scanf("%d ", &n)
	fmt.Println(fibo(n))
}