package main

import (
	"github.com/fluffy-bunny/glowing/pkg/utils"
)

func main() {

	d := utils.GenerateRandomServiceName("hello", 32)
	println("Hello, world." + d)
}
