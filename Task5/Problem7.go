/*
	Bagja 9102 Kurniawan
*/
package main
import "fmt"

func playingDomino(cards [][]int, deck []int) interface{} {
	var biggest int = 0
	out := make([]int, 2)
	for _, val := range cards {
		for _, v := range deck {
			if val[0] == v || val[1] == v {
				if val[0] + val[1] > biggest {
					biggest = val[0] + val[1]
					out[0], out[1] = val[0], val[1]
				}
			}
		}
	}
	if biggest != 0 {
		return out
	}
	return "Tutup Kartu"
}

func main() {
	fmt.Println(playingDomino([][]int{{6, 5}, {3, 4}, {2, 1}, {3, 3}}, []int{4, 3}))
	fmt.Println(playingDomino([][]int{{6, 5}, {3, 3}, {3, 4}, {2, 1}}, []int{3, 6}))
	fmt.Println(playingDomino([][]int{{6, 6}, {2, 4}, {3, 6}}, []int{5, 1}))
}