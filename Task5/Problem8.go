/*
	Bagja 9102 Kurniawan
*/
package main
import (
	"fmt"
	"sort"
)

var size int
var items string
var arrItem []string

type pair struct {
	name string
	muncul int
}

type paired []pair
func (a paired) Len() int 			{ return len(a) }
func (a paired) Less(i, j int) bool	{ return a[i].muncul < a[j].muncul }
func (a paired) Swap(i, j int)		{ a[i], a[j] = a[j], a[i] }

func MostAppearItem(items []string) []pair {
	sort.Strings(items)
	var arr []pair
	arr = append(arr, pair{items[0], 1})
	indeks := 0
	for _, val := range items[1:] {
		if val == arr[indeks].name {
			arr[indeks].muncul ++
		} else {
			indeks ++
			arr = append(arr, pair{val, 1})
		}
	}
	sort.Sort(paired(arr))
	return arr
}

func main() {
	fmt.Print("Length Items = ")
	fmt.Scanf("%d\n", &size)

	for i := 0; i < size; i++ {
		fmt.Print("Item Name : ")
		fmt.Scanf("%s\n", &items)
		arrItem = append(arrItem, items)
	}

	item := MostAppearItem(arrItem)
	for _, val := range item {
		fmt.Print(val.name, "->", val.muncul, " ")
	}
}