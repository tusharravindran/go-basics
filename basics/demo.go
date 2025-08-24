package main

import (
	"fmt"
	"time"
)

func greet(name string, age int) {
	fmt.Printf("Hello, %s! You are %d years old.\n", name, age)
}

func main() {
	name := "Alice"
	age := 30
	greet(name, age)
	time.Sleep(3 * time.Second)
	name2 := "Bob"
	age2 := 25
	greet(name2, age2)
	counter(5)
}

func counter(n int) {
	for i := 1; i <= n; i++ {
		fmt.Printf("Count: %d\n", i)
		time.Sleep(1 * time.Second)
	}
}