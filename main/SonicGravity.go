package main

import "time"

func gravity() {
	for true {
		if originalY < ySonic {
			ySonic = ySonic - 1
		}
		time.Sleep(10 * time.Millisecond)
	}
}
