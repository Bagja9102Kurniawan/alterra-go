/*
	Bagja 9102 Kurniawan
*/
package main
import "fmt"

var a int

func main () {
	fmt.Print("Masukkan Angka : ")
	fmt.Scan(&a)
	fmt.Print("Faktor Bilangan : ")
	for i:=1; i<=a; i++ {
		if a%i == 0 {
			fmt.Print(i, " ")
		}
	}
	fmt.Print("\n")
}