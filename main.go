package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Println("Hello felix!")
		time.Sleep(time.Second * 1)
	}
}
