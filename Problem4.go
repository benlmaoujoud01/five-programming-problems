package main

import (
	"fmt"
	"strconv"
	"strings"
)

// largestNumber creates the largest possible number from a slice of integers
func largestNumber(nums []int) string {
	// Convert numbers to strings
	strNums := make([]string, len(nums))
	for i, num := range nums {
		strNums[i] = strconv.Itoa(num)
	}
	
	for i := 0; i < len(strNums); i++ {
		for j := 0; j < len(strNums)-i-1; j++ {
			if strNums[j]+strNums[j+1] < strNums[j+1]+strNums[j] {
				strNums[j], strNums[j+1] = strNums[j+1], strNums[j]
			}
		}
	}
	
	if strNums[0] == "0" {
		return "0"
	} 
	return strings.Join(strNums, "")
}

func main() {
	// Test cases
	testCases := [][]int{
		{50, 2, 1, 9},       
		{3, 30, 34, 5, 9},    
		{0, 0},               
		{10, 2},             
	}
	
	// Run and print results
	for _, nums := range testCases {
		result := largestNumber(nums)
		fmt.Printf("Input: %v, Largest Number: %s\n", nums, result)
	}
}