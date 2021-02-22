package main

import (
	"fmt"
)

func main() {
	fmt.Println(getData())
}

func getData() (string, int) {
	return "ini data", 1
}
