package Examples

import "fmt"

func Channels() {
	c := make(chan string)
	go CountChannels("ship", c)

	for msg := range c {
		fmt.Println(msg)
	}

}
