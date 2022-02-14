package examples

import (
	"fmt"
	"strconv"
)

type predicate func(int) bool
type mapperStr func(int) string

func Filter(ls []int, cb predicate) []int {
	var res []int

	for _, val := range ls {
		if cb(val) {
			res = append(res, val)
		}
	}

	return res
}

func MapString(ls []int, cb mapperStr) []string {
	var res []string

	for _, val := range ls {
		res = append(res, cb(val))
	}

	return res
}

func toString(value int) string {
	return "0x" + strconv.Itoa(value) + ";"
}

func TryCallback() {
	slice := []int{1, 2, 3, 4, 5, 7}
	isEven := func(n int) bool{ return n % 2 == 0 }
	fmt.Println(Filter(slice, isEven))
	fmt.Println(MapString(slice, toString))
}