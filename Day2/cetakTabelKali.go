/*
	Bagja 9102 Kurniawan
*/
package main
import "fmt"

var angka int

func cetakTabelPerkalian(number int){
	for i:=1; i<=number; i++{
		for j:=1; j<=number; j++{
			if i*j<10{
				fmt.Print(" ")
			}
			fmt.Print(i*j, " ")
		}
		fmt.Print("\n")
	}
}

func main(){
	fmt.Print("Banyak Tabel= ")
	fmt.Scanf("%d\n", &angka)
	cetakTabelPerkalian(angka)
}