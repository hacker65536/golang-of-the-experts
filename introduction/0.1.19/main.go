package main

import (
	"fmt"

	"github.com/hacker65536/0.1.19/pkg/syncutil"
)

func main() {

	c := &syncutil.Counter{
		Name: "Access",
	}

	fmt.Println(c.Increment()) // Output: 1
	fmt.Println(c.View())      // Output: 1
}
