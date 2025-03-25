package main

import (
	"github.com/ChenYJ1016/puppy"

	"fmt"
)

func main() {
	s1, s2 := puppy.Bark(), puppy.Barks()

	fmt.Println(s1)
	fmt.Println(s2)
}
