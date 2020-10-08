package main

import (
	"fmt"
)

func main() {
	if v := 5; v < 10 {
		fmt.Println(v, "is a small number.")
	} else {
		fmt.Println(v, "is a big number.")
	}
}