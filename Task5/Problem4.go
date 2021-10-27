/*
	Bagja 9102 Kurniawan
*/
package main
import "fmt"

var size, input int
var arr []int

func MaxSequence(arr []int) int {
	var hasil, temp int = -999999999999, 0

	for _, num := range arr {
		temp += num
		if (hasil < temp) {
			hasil = temp
		}
		if temp < 0 {
			temp = 0
		}
	}
	return hasil
}

func main() {
	fmt.Print("Length array = ")
	fmt.Scanf("%d\n", &size)
	for i := 0; i < size; i++ {
		fmt.Print("Input Element = ")
		fmt.Scanf("%d\n", &input)
		arr = append(arr, input)
	}
	fmt.Println("Total maksimum jumlah bilangan : ", MaxSequence(arr))
}