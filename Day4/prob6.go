/*
	Bagja 9102 Kurniawan
*/

package main
import (
	"fmt"
	"errors"
)

var menu int
var a = student{}
var c Chiper = &a

type student struct {
    name       string
    nameEncode string
    score      int
}

type Chiper interface {
    Encode() (string, error)
    Decode() (string, error)
}

func (s *student) Encode() (string, error) {
    var nameEncode = ""
    runes := []rune(s.name)
	var ascii rune
	var save int
    for i := 0; i < len(s.name); i++ {
		if s.name[i] >= 97 && s.name[i] <= 122 {
			save = 97 - int(runes[i]) + 122
			ascii = rune(save)
		} else if s.name[i] >= 65 && s.name[i] <= 90 {
			save = 65 - int(runes[i]) + 90
			ascii = rune(save)
		} else {
			return "", errors.New("Input must be character 'a' to 'z' or 'A' to 'Z'")
		}
        nameEncode = nameEncode + string(ascii)
    }
    return nameEncode, nil
}

func (s *student) Decode() (string, error) {
    var nameDecode = ""
    runes := []rune(s.nameEncode)
	var ascii rune
	var save int
    for i := 0; i < len(s.nameEncode); i++ {
		if s.nameEncode[i] >= 97 && s.nameEncode[i] <= 122 {
			save = 97 - int(runes[i]) + 122
			ascii = rune(save)
		} else if s.nameEncode[i] >= 65 && s.nameEncode[i] <= 90 {
			save = 65 - int(runes[i]) + 90
			ascii = rune(save)
		} else {
			return "", errors.New("Input must be character 'a' to 'z' or 'A' to 'Z'")
		}
        nameDecode = nameDecode + string(ascii)
    }
    return nameDecode, nil
}
func main() {
    fmt.Print("Menu\n[1] Encrypt \n[2] Decrypt \nChoose your menu ? ")
    fmt.Scanf("%d\n",&menu)
    if menu == 1 {
        fmt.Print("\nInput Student's Name : ")
        fmt.Scanf("%s\n",&a.name)
		if value, err := c.Encode(); value != "" {
			fmt.Print("\nEncode of Student's Name " + a.name + " is : " + value)
		} else {
			fmt.Println(err.Error())
		}
    } else if menu == 2 {
        fmt.Print("\nInput Student's Encode Name : ")
        fmt.Scanf("%s\n", &a.nameEncode)
		if value, err := c.Decode(); value != "" {
			fmt.Print("\nEncode of Student's Name " + a.name + " is : " + value)
		} else {
			fmt.Println(err.Error())
		}
    } else {
        fmt.Println("Wrong Input")
    }
}