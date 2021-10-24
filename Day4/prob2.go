/*
	Bagja 9102 Kurniawan
*/

package main
import "fmt"

var off int
var inp string

func caesar(offset int, input string) string {
	runes := []rune(input)
	var hasil string
	var tmp rune
	for _,converted := range runes {
		tmp = (converted + rune(offset) -97) % 26 + 97
		hasil += string(tmp)
	}
	return hasil
}

func main() {
	fmt.Print("Offset = ")
	fmt.Scanf("%d\n", &off)
	fmt.Print("Strings = ")
	fmt.Scanf("%s\n", &inp)
	fmt.Println("encoded string =",caesar(off, inp))
}