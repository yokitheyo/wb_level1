package main

import (
	"fmt"
	"reflect"
)

func detectType(v interface{}) {
	switch val := v.(type) {
	case int:
		fmt.Println("type: int, value:", val)
	case string:
		fmt.Println("type: string, value:", val)
	case bool:
		fmt.Println("type: bool, value:", val)
	default:
		t := reflect.TypeOf(v)
		if t.Kind() == reflect.Chan {
			fmt.Println("type: chan, value:", t.String())
		} else {
			fmt.Println("type: unknown, value:", val)
		}

	}
}

func main() {
	detectType(163)
	detectType("hello")
	detectType(true)
	detectType(make(chan int))
	detectType(make(chan string))
	detectType([]int{1, 2, 3})
}
