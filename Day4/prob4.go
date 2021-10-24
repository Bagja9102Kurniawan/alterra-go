/*
	Bagja 9102 Kurniawan
*/

package main
import "fmt"

var in1, in2, in3, in4, in5, in6, min, max int

func getMinMax(numbers ...*int) (min int, max int) {
	min, max = *numbers[0], *numbers[0]
	for _, number := range numbers {
		if *number > max {
			max = *number
		}
		if *number < min {
			min = *number
		}
	}
	return min, max
}

func main() {
	fmt.Print("Number Input = ")
	fmt.Scanf("%d\n", &in1)
	fmt.Print("Number Input = ")
	fmt.Scanf("%d\n", &in2)
	fmt.Print("Number Input = ")
	fmt.Scanf("%d\n", &in3)
	fmt.Print("Number Input = ")
	fmt.Scanf("%d\n", &in4)
	fmt.Print("Number Input = ")
	fmt.Scanf("%d\n", &in5)
	fmt.Print("Number Input = ")
	fmt.Scanf("%d\n", &in6)

	min, max = (getMinMax(&in1, &in2, &in3, &in4, &in5, &in6))
	fmt.Printf("%d is the minimum number\n", min)
	fmt.Printf("%d is the maximum number\n", max)
}