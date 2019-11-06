package main

import (
	"fmt"
	"time"
)

func main() {
	t1 := time.Unix(0, 0)
	fmt.Println(t1.Unix(), t1.UnixNano())

	t2 := time.Unix(0, 0).UTC()
	fmt.Println(t2.Unix(), t2.UnixNano())
}
