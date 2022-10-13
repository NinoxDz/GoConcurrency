package main

import (
	"GoConcurrency/Examples"
	"fmt"
)

func main() {

	println("Hello Go Concurrency")
	println("Which example you would like to run")
	println("1 => GoRoutine")
	println("2 => WaitGroup")
	println("3 => Channels")
	println("4 => Buffer")
	println("5 => SelectStatement")
	println("6 => WorkerPool")
	fmt.Println("Enter example Number: ")
	var input int
	_, err := fmt.Scanf("%d", &input)

	switch input {
	case 1:
		Examples.GoRoutine()
		break
	case 2:
		Examples.WaitGroup()
		break
	case 3:
		Examples.Channels()
		break
	case 4:
		Examples.Buffer()
		break
	case 5:
		Examples.SelectStatement()
		break
	case 6:
		Examples.WorkerPool()
		break
	default:
		fmt.Println("Wrong input")
		fmt.Println(err)
	}
}
