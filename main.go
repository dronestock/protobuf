package main

import (
	"fmt"
	"os"

	"github.com/dronestock/drone"
)

func main() {
	for _, env := range os.Environ() {
		fmt.Println(env)
	}
	panic(drone.Bootstrap(newPlugin))
}
