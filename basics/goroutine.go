package main

import (
    "fmt"
    "time"
)

func printMessage(msg string) {
    for i := 1; i <= 3; i++ {
        fmt.Println(msg, ":", i)
        time.Sleep(500 * time.Millisecond)
    }
}

func main() {
    // run in a goroutine
    go printMessage("Hello from Goroutine")

    // run in main goroutine
    printMessage("Hello from Main")
}
