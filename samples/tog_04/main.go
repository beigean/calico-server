// ref. https://go-tour-jp.appspot.com/moretypes/26
package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	f1 := 0
	f2 := 1
	return func() int {
		ret := f1
		f1 = f2
		f2 = ret + f2
		return ret
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 30; i++ {
		fmt.Println(f())
	}
}
