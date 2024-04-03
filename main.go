package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("welcome to golang time handling")

	presenteTime := time.Now()

	fmt.Println(presenteTime)

	fmt.Println(presenteTime.Format("01-02-2006"))
}