/*
	Bagja 9102 Kurniawan
*/
package main
import "fmt"

var size, input int
var arr []int

func FindMinAndMax(arr []int) string {
	var min, max, iMax, iMin int = arr[0], arr[0], 0, 0
	for i, num := range arr {
		if num > max {
			max, iMax = num, i
		} else if num < min {
			min, iMin = num, i
		}
	}
	return fmt.Sprintf("min: %d index: %d max: %d index: %d", min, iMin, max, iMax)
}

func main() {
	fmt.Print("Length array = ")
	fmt.Scanf("%d\n", &size)
	for i := 0; i < size; i++ {
		fmt.Print("Input Element = ")
		fmt.Scanf("%d\n", &input)
		arr = append(arr, input)
	}
	fmt.Println(FindMinAndMax(arr))
}