/*
	Bagja 9102 Kurniawan
*/
package main
import(
	"fmt"
	"math"
)

var x,n int

func exponent(base, pangkat int) int{
	var temp int = base
	for i:=2; i<=pangkat; i++{
		temp *= base
	}
	return temp
}

func exp(base, pangkat int) int{
	return int(math.Pow(float64(base), float64(pangkat)))
}

func main(){
	fmt.Print("x= ")
	fmt.Scanf("%d\n", &x)
	fmt.Print("n= ")
	fmt.Scanf("%d\n", &n)
	fmt.Println("tanpa library: ", exponent(x,n))
	fmt.Println("dengan library: ", exp(x,n))
}