
package main

import (
    "fmt"
    "sync"
)

func worker(id int, wg *sync.WaitGroup) {
    defer wg.Done() // mark done when finished
    fmt.Println("Worker", id, "starting")
    fmt.Println("Worker", id, "done")
}

func main() {
    var wg sync.WaitGroup

    for i := 1; i <= 3; i++ {
        wg.Add(1)
        go worker(i, &wg)
    }

    wg.Wait() // wait for all workers
    fmt.Println("All workers finished!")
}
