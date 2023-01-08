package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"time"
)

//Intro time
var introTime = 3

//Initial sprite position in the window foe Sonic
var originalX = float64(600)
var originalY = float64(220)
var xSonic = originalX
var ySonic = originalY
var movement = pixel.Vec{X: xSonic, Y: ySonic}
var sonicFrame = 0
var waitingTime = time.Now().Second()

//Initial sprite position in the window for Fly
var originalXFlyEnemy = float64(700)
var originalYFlyEnemy = float64(280)
var xFlyEnemy = originalXFlyEnemy
var yFlyEnemy = originalYFlyEnemy
var flyFrame = 0

//Persist and check the latest keyPressed to restart sonicFrame.
var lastKeyPressed pixelgl.Button

func checkLastKeyPressed(button pixelgl.Button) {
	if lastKeyPressed.String() != button.String() {
		sonicFrame = 0
		lastKeyPressed = button
	}
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
	sonicFrame = 0
	minUpX = float64(0)
	maxUpX = pixelUpRect
}

//Initial sprite rect min,max vector for Down animation
var pixelDownRect = float64(31)
var minDownX = float64(0)
var maxDownX = pixelDownRect

func resetDownVector() {
	sonicFrame = 0
	minDownX = float64(0)
	maxDownX = pixelDownRect
}

//Initial sprite rect min,max vector for Wait animation
var pixelWaitRect = float64(33)
var minWaitX = float64(0)
var maxWaitX = pixelWaitRect

func resetWaitTime() {
	waitingTime = time.Now().Second()
}

//Initial sprite rect min,max vector for Ball animation
var pixelBallRect = float64(30)
var minBallX = float64(0)
var maxBallX = pixelBallRect

//Initial sprite rect min,max vector for Jump animation
var pixelJumpRect = float64(31)
var minJumpX = float64(0)
var maxJumpX = pixelJumpRect

//Initial sprite rect min,max vector for Jump animation
var pixelFlyRect = float64(53)
var minFlyX = float64(0)
var maxFlyX = pixelFlyRect
