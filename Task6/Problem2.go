/*
	Bagja 9102 Kurniawan
*/
package main
import "fmt"

var money int

func moneyCoins(money int) []int {
	arrCoin := []int{10000, 5000, 2000, 1000, 500, 200, 100, 50, 20, 10, 1}
	var res []int
	var temp int
	for _, val := range arrCoin {
		if money % val != money {
			temp =  money / val
			for i := 0; i < temp; i++ {
				res = append(res, val)
				money -= val
			}
		}
	}
	return res
}

func main() {
	fmt.Print("Uang yang ingin ditukar = ")
	fmt.Scanf("%d\n", &money)
	fmt.Println("Uang ditukar dengan pecahan =", moneyCoins(money))
}