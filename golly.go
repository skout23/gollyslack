package main

import (
	"fmt"
	"time"
)

func currentdate() string {
	return time.Now().Format("20060102")
}

func main() {
	datematch := currentdate() + "*"
	fmt.Println(datematch)
}
