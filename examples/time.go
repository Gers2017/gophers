package examples

import (
	"fmt"
	"time"
)

func DisplayDate() {
	t := time.Now()
	fmt.Printf("%02d/%02d/%4d\n", t.Day(), t.Month(), t.Year())
}