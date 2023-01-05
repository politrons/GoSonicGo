package main

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"image"
	_ "image/png"
	"os"
	"time"
)

//Intro time
var introTime = 5

//Initial sprite position in the window
var step = float64(650)
var movement = pixel.Vec{step, step}
var frame = 0

//Initial sprite rect min,max vector for Right animation
var pixelRightRect = 36.6
var minRightX = float64(0)
var maxRightX = pixelRightRect

//Initial sprite rect min,max vector for Left animation
var pixelLeftRect = float64(37.9)
var minLeftX = (pixelLeftRect + 1.8) * 8
var maxLeftX = pixelLeftRect * 9

//Initial Pictures(Sprites) of the game
var logoPic, greenHillPic, sonicStopPic, sonicLeftPic, sonicRightPic = loadGamePictures()

func main() {
	pixelgl.Run(run)
}

func run() {
	win := createWindow()
	greenHillBackground := pixel.NewSprite(greenHillPic, pixel.R(-1200, -800, 1200, 800))
	logoPicture := pixel.NewSprite(logoPic, pixel.R(-824, -768, 1300, 768))
	var logoTime = time.Now().Second() + introTime
	for time.Now().Second() < logoTime {
		win.Update()
		win.Clear(colornames.Grey)
		logoPicture.Draw(win, pixel.IM)
	}
	for !win.Closed() {
		win.Update()
		win.Clear(colornames.Grey)
		animateGame(greenHillBackground, win)
	}
}

/**
Function where we implement the key event handler over the window using [win.Pressed(Button)].
Every time we press a key we start animation for the left, right, up and down, using the sprites per
each Picture configured before.
*/
func animateGame(
	greenHillBackground *pixel.Sprite,
	win *pixelgl.Window,
) {
	greenHillBackground.Draw(win, pixel.IM)
	if win.Pressed(pixelgl.KeyRight) {
		if frame < 9 {
			minRightX = minRightX + pixelRightRect
			maxRightX = maxRightX + pixelRightRect
		} else {
			frame = 0
			minRightX = 0
			maxRightX = pixelRightRect
		}
		frame = frame + 1
		fmt.Printf("Moving right... frame:%d .xMin:%f xMax:%f\n", frame, minRightX, maxRightX)
		sonicSprite := pixel.NewSprite(sonicRightPic, pixel.R(minRightX, 0, maxRightX, 50))

		step = step + 5
		movement = pixel.Vec{step, 220}
		sonicSprite.Draw(win, pixel.IM.Moved(movement))
	} else if win.Pressed(pixelgl.KeyLeft) {
		if frame < 9 {
			minLeftX = minLeftX - (pixelLeftRect + 1.8)
			maxLeftX = maxLeftX - pixelLeftRect
		} else {
			frame = 0
			minLeftX = (pixelLeftRect + 1.8) * 8
			maxLeftX = pixelLeftRect * 9
		}
		frame = frame + 1
		fmt.Printf("Moving left... frame:%d .xMin:%f xMax:%f\n", frame, minLeftX, maxLeftX)
		sonicSprite := pixel.NewSprite(sonicLeftPic, pixel.R(minLeftX, 0, maxLeftX, 50))

		step = step - 5
		movement = pixel.Vec{X: step, Y: 220}
		sonicSprite.Draw(win, pixel.IM.Moved(movement))
	} else {
		movement = pixel.Vec{X: step, Y: 220}
		sonicSprite := pixel.NewSprite(sonicStopPic, pixel.R(0, 0, pixelRightRect, 50))
		sonicSprite.Draw(win, pixel.IM.Moved(movement))
	}
	time.Sleep(40 * time.Millisecond)
}

/**
Using Pixel Library we create Windows using a previous [WindowConfig] with the dimension, vsync, title.
We also set smooth to true to make the animation more smooth and less pixely.
*/
func createWindow() *pixelgl.Window {
	cfg := pixelgl.WindowConfig{
		Title:  "Go Sonic Go!",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	win.SetSmooth(true)
	return win
}

/**
Function to load all game Pictures
*/
func loadGamePictures() (pixel.Picture, pixel.Picture, pixel.Picture, pixel.Picture, pixel.Picture) {
	logoPic, err := loadPicture("sprites/logo.png")
	greenHillPic, err := loadPicture("sprites/green_hill.png")
	sonicStopPic, err := loadPicture("sprites/sonic_stop.png")
	sonicLeftPic, err := loadPicture("sprites/sonic_left.png")
	sonicRightPic, err := loadPicture("sprites/sonic_right.png")
	if err != nil {
		panic(err)
	}
	return logoPic, greenHillPic, sonicStopPic, sonicLeftPic, sonicRightPic
}

/**
Function to load a game Pictures from a path using [os.Open] to get a file, [image] Decode to obtain an image
and from Pixel library [PictureDataFromImage] to create a [Picture] type from [Image]
*/
func loadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return pixel.PictureDataFromImage(img), nil
}
