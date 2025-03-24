package main

import "fmt"

func combine(l1 []any ,l2 []any)any { 
	s1 :=len(l1)
	s2 :=len(l2)
	rang := max(s1,s2)
	l3  := []any{}
	for i:=0 ;i<rang; i++{
		if i<s1{
			l3 = append(l3, l1[i])
		}
		if i<s2{
			l3 = append(l3, l2[i])
		}
	}
	return l3
}

func main() { 
	
	l1 := []any{1,2,3}      // Mixed type slice
	l2 := []any{"a", "b", "c"}      // Another mixed type slice

	result := combine(l1, l2)
	fmt.Println(result) // Output: [1 a 3 b 5 c]
}