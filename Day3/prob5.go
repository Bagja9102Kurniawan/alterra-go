/*
	Bagja 9102 Kurniawan
	
	fungsi pair target sum
	menjumlah angka pada indeks tertentu agar hasil sesuai target
	diasumsikan array terurut membesar (dapat dilakukan dengan konsep binary search)
	pendekatran dilakukan dengan depan + belakang dan akan bertemu di tengah
*/

package main
import "fmt"

var size, target int

func pairSum(arr []int, target int) []int{
	kr := 0; kn := len(arr)-1; var index []int; var sum int
	for kr < kn{
		sum = arr[kr] + arr[kn]
		if sum == target{
			index = append(index, kr)
			index = append(index, kn)
		}
		if sum < target{
			kr++
		}else{
			kn--
		}
	}
	return index
}

func input() []int{
	fmt.Print("Ukuran Array = ")
	fmt.Scanf("%d\n", &size)
	
	arr := make([]int, size)
	for i:=0; i<size; i++{
		fmt.Print("Masukkan elemen: ")
		fmt.Scanf("%d\n", &arr[i])
	}
	return arr
}

func main(){
	array := input()
	fmt.Print("Target = ")
	fmt.Scanf("%d\n", &target)
	fmt.Println("pada indeks : ", pairSum(array, target))
}