package main

import (
	"fmt"
	"time"
)

func main() {
	for true {
		fmt.Println("Hey its Gofer")
		time.Sleep(1000 * time.Millisecond)
	}
}
