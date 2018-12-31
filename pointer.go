package main

import (
	"fmt"
)

//usually the increment function will make a copy
//not modify the original, but this method can
func main() {
	i := 7
	inc(&i) //& points at the reference address
	fmt.Println(i) //will print 8
}

func inc(x *int) {
	*x++  //* de-reference the address, if not the address will be incremented
}