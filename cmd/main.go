package main

import (
	// "replenish"
	"web"
	// "network"
	"webcam"
)

func main() {
	// replenish.ReplenishImageSet(5)
	//replenish.RemoveAdditional()
    web.Listen()
	// network.Start()
	webcam.Connect()
}