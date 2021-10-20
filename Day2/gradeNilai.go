/*
	Bagja 9102 Kurniawan
*/
package main
import "fmt"

var score int

func main() {
	fmt.Print("Nilai Siswa= ")
	fmt.Scanf("%d",&score)
	
	if(score<=100 && score >=0){
		switch {
		case score >= 80 :
			fmt.Println("Nilai A")
		case score >= 65 :
			fmt.Println("Nilai B")
		case score >= 50 :
			fmt.Println("Nilai C")
		case score >= 35 :
			fmt.Println("Nilai D")
		default:
			fmt.Println("Nilai E")
		}
	} else{
		fmt.Println("Input Salah")
	}
}