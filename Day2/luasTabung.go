/*
	Bagja 9102 Kurniawan
*/
package main
import "fmt"

var T, r float64
const phi = 3.14
func main() {
	fmt.Print("T= ")
	fmt.Scanf("%f\n",&T)
	fmt.Print("r= ")
	fmt.Scanf("%f\n",&r)
	
	fmt.Println("Luas Permukaan Tabung: ", float64(2*phi*r*(r+T)))
}