package examples

import (
	"fmt"
	"strconv"
)

func ErrorCase1() {
	stri := "BBB"
	var intVal int
	intVal, err := strconv.Atoi(stri)

	if err != nil {
		fmt.Printf("error %s\n", stri)
		return
	}
	
	fmt.Printf("parsed %d\n", intVal)
}
