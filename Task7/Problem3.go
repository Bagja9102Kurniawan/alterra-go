/*
	Bagja 9102 Kurniawan
*/
package main
import (
	"fmt"
	"math"
)

var input, size int
var jumps []int

func Frog(jumps []int, cost []int) int {
	cost[1] = int(math.Abs(float64(jumps[0] - jumps[1])))
	for i := 2; i < len(jumps); i++ {
		cost[i] = int(math.Min(
			math.Abs(float64(jumps[i] - jumps[i - 1])) + float64(cost[i - 1]), 
			math.Abs(float64(jumps[i] - jumps[i - 2])) + float64(cost[i - 2])))
	}
	return cost[len(jumps)-1]
}

func main() {
	fmt.Print("How many stones are there? ")
	fmt.Scanf("%d\n", &size)
	for i := 0; i < size; i++ {
		fmt.Printf("Stone %d height = ", i+1)
		fmt.Scanf("%d\n", &input)
		jumps = append(jumps, input)
	}
	cost := make([]int, len(jumps)+1)
	fmt.Println("Total Jump Costs = ", Frog(jumps, cost))
}