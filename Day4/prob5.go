/*
	Bagja 9102 Kurniawan
*/

package main
import "fmt"

var std = Student{}
var name string
var score int

type Student struct{
    name []string
    score []int
}
 
func (s Student) Average() float64{
    sum := 0.0
    for _, value := range s.score{
        sum += float64(value)
    }
    return sum / float64(len(s.score))
}
 
func (s Student) Min() (min int, name string){
    min = s.score[0]
    indeks := 0
    for i, value := range s.score{
        if min > value{
            min = value
            indeks = i
        }
    }
    return min, s.name[indeks]
}
 
func (s Student) Max() (max int, name string){
    max = s.score[0]
    indeks := 0
    for i, value := range s.score{
        if max < value{
            max = value
            indeks = i
        }
    }
    return max, s.name[indeks]
}
 
func main(){
    for i := 1; i < 6; i++ {
        fmt.Print("Input ", i ," Student’s Name : ")
        fmt.Scanf("%s\n",&name)
		fmt.Print("Input " + name + "'s Score : ")
        fmt.Scanf("%d\n",&score)
		for j := 0; j < len(std.name); j++ {
			if std.name[j] == name && std.score[j] == score {
				fmt.Println("Data already exist! Try again")
				fmt.Print("Input ", i ," Student’s Name : ")
        		fmt.Scanf("%s\n",&name)
				fmt.Print("Input " + name + " Score : ")
				fmt.Scanf("%d\n",&score)
			}
		}
        std.name = append(std.name, name)
        std.score = append(std.score, score)
    }
    fmt.Println("\nAvarage Score Students is", std.Average())
	scoreMin, nameMin := std.Min()
    fmt.Println("Min Score Students is : "+nameMin+" (", scoreMin, ")")
    scoreMax, nameMax := std.Max()
    fmt.Println("Max Score Students is : "+nameMax+" (", scoreMax, ")")
}