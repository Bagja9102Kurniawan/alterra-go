/*
	Bagja 9102 Kurniawan
*/
package main
import (
	"fmt"
	"sort"
)

var money, size, price int
var productPrice []int
	
func MaximumBuyProduct(money int, productPrice []int) {
	sort.Ints(productPrice)
	var product int = 0
	for _, val := range productPrice {
		if money >= val {
			money -= val
			product++
		}
	}
	fmt.Printf("You can buy %d product\n", product)
}

func main() {
	fmt.Print("Money = ")
	fmt.Scanf("%d\n", &money)
	fmt.Print("Length Product = ")
	fmt.Scanf("%d\n", &size)
	for i := 0; i < size; i++ {
		fmt.Printf("Item %d Price : ", i+1)
		fmt.Scanf("%d\n", &price)
		productPrice = append(productPrice, price)
	}
	MaximumBuyProduct(money, productPrice)
}