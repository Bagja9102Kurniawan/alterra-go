/*
	Bagja 9102 Kurniawan
*/
package main
import "fmt"

var a, b, c int

func SimpleEquations(a, b, c int) string {
	for x := 1; x < c; x++ {
		for y := 1; y < c; y ++ {
			for z := 1; z < c; z++ {
				if (x + y + z) == a && (x * y * z) == b && ((x * x) + (y * y) + (z * z)) == c {
					return fmt.Sprintf("x=%d; y=%d; z=%d", x, y, z)
				}
			}
		}
	}
	return fmt.Sprintf("no solution\n")
}

func main() {
	fmt.Print("Elemen a = ")
	fmt.Scanf("%d\n", &a)
	fmt.Print("Elemen b = ")
	fmt.Scanf("%d\n", &b)
	fmt.Print("Elemen c = ")
	fmt.Scanf("%d\n", &c)
	fmt.Println(SimpleEquations(a, b, c))
}