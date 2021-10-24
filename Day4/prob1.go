/*
	Bagja 9102 Kurniawan
*/

package main
import (
	"fmt"
	"strings"
	"errors"
)

var strA, strB string

func Compare(a, b string) (string, error) {
	if len(a) == 0 && len(b) == 0 {
		return "", errors.New("Strings are empty")
	} else if len(a) > len(b) {
		if strings.Contains(a, b) {
			return b, nil
		}
	} else {
		if strings.Contains(b, a) {
			return a, nil
		}
	}
	return "", errors.New("Neither " + a + " nor " + b + " are substrings for each other" )
}

func main() {
	fmt.Print("String A = ")
	fmt.Scanf("%s\n", &strA)
	fmt.Print("String B = ")
	fmt.Scanf("%s\n", &strB)
	
	if value, err := Compare(strA, strB); value != "" {
		fmt.Println(value)
	} else {
		fmt.Println(err.Error())
	}
}