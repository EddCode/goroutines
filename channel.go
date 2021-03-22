package main

import (
	"fmt"
	"time"
)

func main() {
    fmt.Println("Press ENTER to finish execution")
    ch := make(chan string)

    go sendPing(ch)
    go showPing(ch)

    var input string
    fmt.Scanln(&input)
    fmt.Println("END...")
}

func sendPing(ch chan string) {
    for {
        ch <- "Ping"
    }
}

func showPing(ch chan string)  {
    for {
        fmt.Println(<-ch )
        time.Sleep(time.Second)
    }
}
