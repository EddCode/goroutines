package main

import (
	"fmt"
    "math/rand"
)

func generate(ch chan int)  {
    count := 10
    for i := 0; i < count; i++ {
        x := rand.Intn(100)
        fmt.Println("Generate Number ", x)
        ch <- x
    }

    close(ch)
}

func double(chDouble, ch chan int)  {
   for i := range ch {
     chDouble <- i * 2
   }

   close(chDouble)
}

func main() {
    ch := make(chan int)
    chDouble := make(chan int)

    go generate(ch)
    go double(chDouble, ch)

    for i := range chDouble {
        fmt.Println("Doubled number: ", i)
    }
}
