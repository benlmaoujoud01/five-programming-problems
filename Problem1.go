package main

import "fmt"


func sumwithfor(arr []int) int {

	sum := 0
	for i :=0; i<len(arr);i++{
		sum +=arr[i]
	}

	return sum
}

func sumwithwhile(arr [] int) int {
	sum:= 0
	i  := 0
	for i<len(arr){
		sum += arr[i]
		i++
		if(i==len(arr)){
			break
		}

	}
	return sum
}

func sumrec(arr []int, rep int) int {

	if(rep >= len(arr)){
		return 0
	}
	
	return arr[rep]+sumrec(arr,rep+1)
}


func main() {
	arr := []int{1, 2, 3, 4, 5}

	fmt.Println("Sum using for loop:", sumwithfor(arr))

	fmt.Println("Sum using while loop:", sumwithwhile(arr))

	fmt.Println("Sum using recursion:", sumrec(arr, 0))
}