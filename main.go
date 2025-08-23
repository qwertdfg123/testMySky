package main

import (
	"fmt"
)

var s = "fuck man"

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	var s1 []int = make([]int, 0)
	fmt.Println(s1)
	var s2 []int = make([]int, 0)
	fmt.Println(s2)
	var s3 []int = make([]int, 0)
	fmt.Println(s3)

}

func f() {
	fmt.Println(s)
}
