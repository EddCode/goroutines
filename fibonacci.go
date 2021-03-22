package main

import (
	"fmt"
	"time"
)

func waitUntilFiboResult(delay time.Duration)  {
    for {
        for _, value := range `\|/-` {
            fmt.Printf("\r%c", value)
            time.Sleep(delay)
        }
    }
}

func fibo(number int) int {
    if number < 2 {
        return number
    }

    return fibo(number-1) + fibo(number-2)
}

func main() {
    var calculate int

    fmt.Println("Insert the number to calculate fibonaci sequence")
    fmt.Scanf("%d", &calculate)

    fmt.Print("====>")
    go waitUntilFiboResult(100 * time.Millisecond)
    time.Sleep(1000 * time.Millisecond)
    fibonacciResult := fibo(calculate)

    fmt.Printf("\rF(%d) => %d", calculate, fibonacciResult)

}
