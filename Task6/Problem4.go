/*
	Bagja 9102 Kurniawan
*/
package main
import (
	"fmt"
	"sort"
)

func BinarySearch(array []int, x int) {
	var mid int
	sort.Ints(array)
	var indeks_kiri, indeks_kanan int = 0, len(array) - 1

	for indeks_kiri < indeks_kanan {
		mid = (indeks_kiri + indeks_kanan) / 2
		if array[mid] < x {
			indeks_kiri = mid + 1
		} else if array[mid] > x {
			indeks_kanan = mid
		} else {
			fmt.Println(mid)
			break
		}
	}

	if array[mid] != x {
		fmt.Println(-1)
	}
}

func main() {
	BinarySearch([]int{1, 1, 3, 5, 5, 6, 7}, 3) // 2
	BinarySearch([]int{1, 2, 3, 5, 6, 8, 10}, 5) // 3
	BinarySearch([]int{12, 15, 15, 19, 24, 31, 53, 59, 60}, 53)  // 6
	BinarySearch([]int{12, 15, 15, 19, 24, 31, 53, 59, 60}, 100) // -1
}