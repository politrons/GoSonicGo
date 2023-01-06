package main

import "github.com/faiface/pixel"

//Intro time
var introTime = 1

//Initial sprite position in the window
var step = float64(650)
var movement = pixel.Vec{X: step, Y: step}
var frame = 0

//Initial sprite rect min,max vector for Right animation
var pixelRightRect = 36.6
var minRightX = float64(0)
var maxRightX = pixelRightRect

//Initial sprite rect min,max vector for Left animation
var pixelLeftRect = 37.9
var minLeftX = (pixelLeftRect + 1.8) * 8
var maxLeftX = pixelLeftRect * 9

//Initial sprite rect min,max vector for Right animation
var pixelUpRect = float64(30)
var minUpX = float64(0)
var maxUpX = pixelUpRect

func resetUpVector() {
	frame = 0
	minUpX = float64(0)
	maxUpX = pixelUpRect
}
