package main

import (
	"fmt"
	"time"
)

func main() {
	//timer("push up")
	go timer("push up")
	go timer("pull up")
	//go timer("pull up")

	//fmt.Scanln()
}

func timer(todo string) {
	for i := 1; i < 5; i++ {
		fmt.Println(todo, i)
		time.Sleep(time.Second)
	}
}
