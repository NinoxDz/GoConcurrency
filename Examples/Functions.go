package Examples

import (
	"fmt"
	"time"
)

func Count(something string) {
	for i := 1; true; i++ {
		fmt.Println(i, something)
		time.Sleep(time.Millisecond * 500)
	}
}

func CountChannels(something string, c chan string) {
	for i := 1; i < 5; i++ {
		c <- something
		time.Sleep(time.Millisecond * 500)
	}
	close(c)
}

func Fib(n int) int {
	if n <= 1 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}
func DynamicFib(n int) int {
	fibMap := make(map[int]int)
	for i := 0; i < n; i++ {
		if i < 3 {
			fibMap[i] = i
		} else {
			fibMap[i] = fibMap[i-1] + fibMap[i-2]
		}
	}
	return fibMap[n-1]
}
