package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"time"
)

//Intro time
var introTime = 1

//Initial sprite position in the window
var step = float64(650)
var movement = pixel.Vec{X: step, Y: step}
var frame = 0
var waitingTime = time.Now().Second()

//Persist and check the latest keyPressed to restart frame.
var lastKeyPressed pixelgl.Button

func checkLastKeyPressed(button pixelgl.Button) {
	if lastKeyPressed != button {
		frame = 0
	}
	lastKeyPressed = button
}

//Initial sprite rect min,max vector for Right animation
var pixelRightRect = 36.6
var minRightX = float64(0)
var maxRightX = pixelRightRect

//Initial sprite rect min,max vector for Left animation
var pixelLeftRect = 37.9
var minLeftX = (pixelLeftRect + 1.8) * 8
var maxLeftX = pixelLeftRect * 9

//Initial sprite rect min,max vector for Up animation
var pixelUpRect = float64(30)
var minUpX = float64(0)
var maxUpX = pixelUpRect

func resetUpVector() {
	frame = 0
	minUpX = float64(0)
	maxUpX = pixelUpRect
}

//Initial sprite rect min,max vector for Down animation
var pixelDownRect = float64(31)
var minDownX = float64(0)
var maxDownX = pixelDownRect

func resetDownVector() {
	frame = 0
	minDownX = float64(0)
	maxDownX = pixelDownRect
}

//Initial sprite rect min,max vector for Up animation
var pixelWaitRect = float64(33)
var minWaitX = float64(0)
var maxWaitX = pixelWaitRect

func resetWaitTime() {
	waitingTime = time.Now().Second()
}

//Initial sprite rect min,max vector for Down animation
var pixelBallRect = float64(30)
var minBallX = float64(0)
var maxBallX = pixelBallRect
