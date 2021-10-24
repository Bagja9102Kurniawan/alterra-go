/*
	Bagja 9102 Kurniawan
*/

package main
import "fmt"

var a, b int

func swap(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func main() {
	fmt.Print("a = ")
	fmt.Scanf("%d\n", &a)
	fmt.Print("b = ")
	fmt.Scanf("%d\n", &b)
	fmt.Println("BEFORE SWAP ==>\t",a,b)
	swap(&a, &b)
	fmt.Println("AFTER SWAP  ==>\t",a,b)
}