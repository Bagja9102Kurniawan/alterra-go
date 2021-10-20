/*
	Bagja 9102 Kurniawan
*/
package main
import "fmt"

var t int

func playWithAsterik(height int){
	for i:=1; i<=height; i++{
		for sp:=1; sp<=height-i; sp++{
			fmt.Print(" ")
		}
		for st:=1; st<=i; st++{
			fmt.Print("* ")
		}
		fmt.Print("\n")
	}
}

func main(){
	fmt.Print("Masukkan Tinggi Segitiga= ")
	fmt.Scanf("%d\n", &t)
	playWithAsterik(t)
}