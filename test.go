package main

import "fmt"

func printHi() {
	fmt.Print("Hi~\n")
}

func sum(a int, b int) int {
	result := a + b
	return result
}

func main() {
	var name string

	i := 0
	for {
		fmt.Println(i)
		if i >= 10 {
			break
		}
		i++
	}

	for i := 0; i <= 10; i++ {
		fmt.Print(i)
	}
	fmt.Print("\n")

	fmt.Println("hello world")
	fmt.Println(sum(3, 5))
	printHi()

	fmt.Print("이름 입력 : ")
	fmt.Scanln(&name)
	fmt.Println(name)
}
