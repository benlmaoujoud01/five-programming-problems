package main

import "fmt"

func fibo(n int) int {
    if n == 0 {
        return 0
    }
    if n == 1 {
        return 1
    }

    a, b := 0, 1
    var c int
    for i := 2; i <= n; i++ {
        c = a + b
        a = b 
        b = c  
    }
    return c
}

func main() {
    for i := 0; i < 10; i++ {
        fmt.Println(fibo(i))
    }
}
