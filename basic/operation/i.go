package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := 10
	b := 20

	fmt.Printf(strconv.Itoa(a+b) + "\r\n")
	fmt.Print(strconv.Itoa(a-b) + "\r\n")
	fmt.Print(strconv.Itoa(a*b) + "\r\n")
	fmt.Print(strconv.Itoa(a/b) + "\r\n")
	fmt.Print(strconv.Itoa(a%b) + "\r\n")

	/* golang不同
	* 在golang中的++和--并不是操作而是表达式，所以不可以像php一样赋值
	* 同时也没有--x和++x
	 */
	a++
	fmt.Print(strconv.Itoa(a) + "\r\n")

	a--
	fmt.Print(strconv.Itoa(a) + "\r\n")

}
