package main

import (
	"fmt"
)

var s = "fuck man"

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	array := [2]int{1, 2}
	b := &array
	b[0] = 2
	s := array[1:]
	c := s
	c[0] = 1
	fmt.Println(c)

}

func f() {
	fmt.Println(s)
}
