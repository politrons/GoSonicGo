package main

import "time"

func gravity() {
	for true {
		if originalY < yStep {
			yStep = yStep - 1
		}
		time.Sleep(10 * time.Millisecond)
	}
}
