/*
	Bagja 9102 Kurniawan
	
	fungsi array merge
	dilakukan pengecekan kesamaan anatara array a dan b
	merge tanpa terjadi duplikat
*/
package main
import "fmt"

var size, count int

func duplikat(arrayA, arrayB []string) ([]string, []string){
	for i:=0; i<len(arrayA); i++{
		for j:=0; j<len(arrayB); j++{
			if arrayA[i] == arrayB[j]{
				arrayB[j] = arrayB[j+1]
				arrayB = arrayB[:len(arrayB)-1]
			}
		}
	}
	return arrayA, arrayB
}

func ArrayMerge(arrayA, arrayB []string) []string {
	arrayA, arrayB = duplikat(arrayA, arrayB)
	merged := append(arrayA, arrayB...)
	return merged
}

func input() []string{
	fmt.Print("Ukuran Array ",count,"= ")
	fmt.Scanf("%d\n", &size)
	
	arr := make([]string, size)
	for i:=0; i<size; i++{
		fmt.Print("Masukkan elemen: ")
		fmt.Scanf("%s\n", &arr[i])
	}
	count++
	return arr
}

func main() {
	count = 1
	arrA := input()
	arrB := input()
	fmt.Println("Array hasil: ", ArrayMerge(arrA, arrB))
}