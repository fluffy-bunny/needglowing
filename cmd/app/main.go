package main

import (
	"github.com/fluffy-bunny/glowing/pkg/junk"
	"github.com/fluffy-bunny/glowing/pkg/utils"
)

func main() {

	println("Hello, world." + utils.GenerateRandomServiceName("hello", 32))
	println("Hello, world." + junk.GenerateRandomServiceName("hello", 32))

}
