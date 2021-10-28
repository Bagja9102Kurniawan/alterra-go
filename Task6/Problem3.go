/*
	Bagja 9102 Kurniawan
*/
package main
import (
	"fmt"
	"sort"
)

func DragonOfLoowater(dragonHead, knightHeight []int) {
	sort.Ints(dragonHead)
	sort.Ints(knightHeight)
	
	var choppedHead,totalHeight int = 0, 0
	
	if len(dragonHead) > len(knightHeight) {
		fmt.Println("knight fall")
	} else {
		for i := range dragonHead {
			for j := range knightHeight {
				if knightHeight[j] >= dragonHead[i] {
					totalHeight += knightHeight[j]
					choppedHead += 1
					break
				}
			}
		}
		if choppedHead == len(dragonHead) {
			fmt.Println(totalHeight)
		} else {
			fmt.Println("knight fall")
		}
	}
}

func main() {
	DragonOfLoowater([]int{5, 4}, []int{7, 8, 4})    // 11
	DragonOfLoowater([]int{5, 10}, []int{5})         // knight fall
	DragonOfLoowater([]int{7, 2}, []int{4, 3, 1, 2}) // knight fall
	DragonOfLoowater([]int{7, 2}, []int{2, 1, 8, 5}) // 10
}