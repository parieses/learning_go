package main

import (
	"B"
	"fmt"
)
// var (
// 	a int = 1
// 	b string
// )
const a  = iota
const name  = "亮"
const names string  = "亮"
const strs  = len(names)
const b  = iota
const(
	c = iota//0
	d = iota//1
	_
	e = iota//3
	f = iota *2 //8
	g  //10
)

func main() {
	//var a = 25

	//print('\n')
	//fmt.Print(reflect.TypeOf(a))
	//print('\n')
	fmt.Print(name)
	fmt.Print(B.B)
	fmt.Print("\n")
	fmt.Print(a)
	fmt.Print("\n")
	fmt.Print(b)
	fmt.Print("\n")
	fmt.Print(c)
	fmt.Print("\n")
	fmt.Print(d)
	fmt.Print("\n")
	fmt.Print(e)
	fmt.Print("\n")
	fmt.Print(f)
	fmt.Print("\n")
	fmt.Print(g)



}
