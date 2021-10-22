/*
	Bagja 9102 Kurniawan
	
	fungsi eksponensial
	menggunakan metode exponentiation by squaring
	konsepnya sama seperti algoritma Divide and Conquer
	kompleksitas O(log2n)
*/
package main
import "fmt"

var n, x int

func pow(x,n int) int{
	var temp int = 1
	for n > 0{
		if n%2==1{
			temp *= x
		}
		n = n/2
		x *= x
	}
	return temp
}

func main(){
	fmt.Print("x= ")
	fmt.Scanf("%d\n", &x)
	fmt.Print("n= ")
	fmt.Scanf("%d\n", &n)
	fmt.Println("hasil eksponensial: ", pow(x,n))
}