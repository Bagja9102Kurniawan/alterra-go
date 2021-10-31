/*
	Bagja 9102 Kurniawan
*/
package main
import (
	"fmt"
	"strings"
)

type Freq map[rune]int

func Frequency(s string) Freq {
	x := Freq{}
	for _, val := range s {
		x[val]++
		fmt.Print(string(val), ":", x[val], " ")
	}
	return x
}

func ConcuFreq(str []string) {
	res := make(chan Freq)
	for _, s := range str {
		go func(s string) {
			res <- Frequency(s)
		}(s)
	}
	fmt.Println("\nCounter =")
	totals := make(map[rune]int)
	for range str {
		freqmap := <- res
		for char, count := range freqmap {
			totals[char] += count
		}
	}

	fmt.Println("\n")
	fmt.Println("Letter Frequency =")
	for val := range totals {
		fmt.Print(string(val), ":", totals[rune(val)], " ")
	}
	fmt.Println("")
}

func main() {
	arr := strings.Fields("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua")
	ConcuFreq((arr))
}