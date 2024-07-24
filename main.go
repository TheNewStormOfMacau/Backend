package main

import (
	"backend/core"
	"backend/eth"
)

func main() {
	go eth.Init()
	core.Init()
}
