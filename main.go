package main

import (
	"fmt"
	"go_project/core"
)

func main() {
	config, err := core.SetConfig()
	if err != nil {
		fmt.Println(err)
		return
	}
	core.SetMap(config)
}
