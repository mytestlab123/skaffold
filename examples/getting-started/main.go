package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Println("Hello k8s world !!!")

		time.Sleep(time.Second * 1)
	}
}
