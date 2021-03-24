package main

import "fmt"

func firstMesage(count int, ch chan string)  {
    for i := 0; i < count; i++ {
       ch <- `channel one`
    }

   close(ch)
}


func secondMessage(count int, ch chan string)  {
   for i := 0; i < count; i++ {
       ch <- "channel two"
   }
   close(ch)
}

func main() {
    ch1 := make(chan string)
    ch2 := make(chan string)

    go firstMesage(5, ch1)
    go secondMessage(2, ch2)

    for {
        select {
        case msg, ok := <-ch1:
            fmt.Println(msg)
            if !ok {
                ch1 = nil
            }
        case msg, ok := <-ch2:
            fmt.Println(msg)
            if !ok {
                ch2 = nil
            }
        }

        if ch1 == nil && ch2 == nil {
            break
        }
    }
}
