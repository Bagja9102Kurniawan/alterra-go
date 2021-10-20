/*
	Bagja 9102 Kurniawan
*/
package main
import "fmt"

var kata string

func palindrome(input string) bool{
	for i:=0; i<int(len(input)/2); i++ {
		j := (len(input)-1-i)
		if(input[i] != input[j]){
			return false
		}
	}
	return true
}

func main(){
	fmt.Print("Masukkan kata= ")
	fmt.Scanf("%s",&kata)
	
	if palindrome(kata){
		fmt.Println("Palindrome")
	} else{
		fmt.Println("Bukan Palindrome")
	}
}