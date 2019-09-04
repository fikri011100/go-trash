package main

import "fmt"

func main() {
	stock := make(map[string]int)

	stock["apple"] = 10
	stock["tomato"] = 12
	stock["cherry"] = 0
	stock["blueberry"] = 15

	delete(stock, "blueberry")

	fmt.Println(stock)
	fmt.Println(stock["tomato"])
}
