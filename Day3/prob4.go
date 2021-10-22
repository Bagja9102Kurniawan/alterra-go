/*
	Bagja 9102 Kurniawan
	
	fungsi pengecekan kemunculan angka
*/
package main
import (
	"fmt"
	"strings"
	"strconv"
)

var input string

func munculSekali(angka string)[]int{
	strArr := strings.Split(angka, "")
	
	checker := make(map[string]int)
	for _, val := range strArr{
		checker[val]++
	}
	var sekali []int
	for number := range checker{
		if checker[number] == 1{
			num, err := strconv.Atoi(number)
			if err != nil{
				fmt.Println(err)
			}
			sekali = append(sekali, num)
		}
	}
	return sekali
}

func main(){
	fmt.Print("Masukkan angka = ")
	fmt.Scanf("%s\n", &input)
	
	fmt.Println("Angka yang muncul sekali: ", munculSekali(input))
}