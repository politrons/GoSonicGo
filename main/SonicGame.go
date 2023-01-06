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

type SonicSprites struct {
	picture pixel.Picture
}

type DrawFeature interface {
	drawPicture(window pixelgl.Window, minX float64, maxX float64)
}

//Initial Pictures(Sprites) of the game
var logoPic,
	greenHillPic,
	sonicStopPic,
	sonicLeftPic,
	sonicRightPic,
	sonicUpPic,
	sonicDownPic,
	sonicWaitingPic = loadGamePictures()

func main() {
	pixelgl.Run(run)
}

func run() {
	win := createWindow()
	greenHillBackground := pixel.NewSprite(greenHillPic.picture, pixel.R(-1200, -800, 1200, 800))
	logoPicture := pixel.NewSprite(logoPic.picture, pixel.R(-824, -768, 1300, 768))
	var logoTime = time.Now().Second() + introTime
	for time.Now().Second() < logoTime {
		win.Update()
		win.Clear(colornames.Grey)
		logoPicture.Draw(win, pixel.IM)
	}
	animateGame(greenHillBackground, win)
}

/**
Function where we implement the key event handler over the window using [win.Pressed(Button)].
Every time we press a key we start animation for the left, right, up and down, using the sprites per
each Picture configured before.
*/
func animateGame(greenHillBackground *pixel.Sprite, win *pixelgl.Window) {
	for !win.Closed() {
		win.Update()
		win.Clear(colornames.Grey)
		greenHillBackground.Draw(win, pixel.IM)
		if win.Pressed(pixelgl.KeyRight) {
			checkLastKeyPressed(pixelgl.KeyRight)
			rightAnimation(win)
			resetWaitTime()
		} else if win.Pressed(pixelgl.KeyLeft) {
			checkLastKeyPressed(pixelgl.KeyLeft)
			leftAnimation(win)
			resetWaitTime()
		} else if win.Pressed(pixelgl.KeyUp) {
			checkLastKeyPressed(pixelgl.KeyUp)
			upAnimation(win)
			resetWaitTime()
		} else if win.Pressed(pixelgl.KeyDown) {
			checkLastKeyPressed(pixelgl.KeyDown)
			downAnimation(win)
			resetWaitTime()
		} else {
			if (waitingTime + 3) < time.Now().Second() {
				waitingAnimation(win)
			} else {
				stopAnimation(win)
				resetUpVector()
				resetDownVector()
			}
		}
		time.Sleep(40 * time.Millisecond)
	}
}

func stopAnimation(win *pixelgl.Window) {
	sonicStopPic.drawPicture(win, 0, pixelRightRect)
}

func waitingAnimation(win *pixelgl.Window) {
	if frame < 2 {
		minWaitX = minWaitX + pixelWaitRect
		maxWaitX = maxWaitX + pixelWaitRect
		frame = frame + 1
	} else {
		frame = 0
		minWaitX = 0
		maxWaitX = pixelWaitRect
	}
	sonicWaitingPic.drawPicture(win, minWaitX, maxWaitX)
	time.Sleep(200 * time.Millisecond)
}

func upAnimation(win *pixelgl.Window) {
	if frame < 5 {
		minUpX = minUpX + pixelUpRect
		maxUpX = maxUpX + pixelUpRect
		frame = frame + 1
	}
	sonicUpPic.drawPicture(win, minUpX, maxUpX)
}

func downAnimation(win *pixelgl.Window) {
	if frame < 5 {
		minDownX = minDownX + pixelDownRect
		maxDownX = maxDownX + pixelDownRect
		frame = frame + 1
	}
	sonicDownPic.drawPicture(win, minDownX, maxDownX)
}

func leftAnimation(win *pixelgl.Window) {
	if frame < 9 {
		minLeftX = minLeftX - (pixelLeftRect + 1.8)
		maxLeftX = maxLeftX - pixelLeftRect
	} else {
		frame = 0
		minLeftX = (pixelLeftRect + 1.8) * 8
		maxLeftX = pixelLeftRect * 9
	}
	frame = frame + 1
	step = step - 5
	sonicLeftPic.drawPicture(win, minLeftX, maxLeftX)
}

func rightAnimation(win *pixelgl.Window) {
	if frame < 9 {
		minRightX = minRightX + pixelRightRect
		maxRightX = maxRightX + pixelRightRect
	} else {
		frame = 0
		minRightX = 0
		maxRightX = pixelRightRect
	}
	frame = frame + 1
	step = step + 5
	sonicRightPic.drawPicture(win, minRightX, maxRightX)
}

/**
Implementation of [SonicSprites] to draw the specific sprite in the window.
Using interface each group of sprites invokes this method, avoiding type mismatch
*/
func (pic SonicSprites) drawPicture(win *pixelgl.Window, minX float64, maxX float64) {
	fmt.Printf("Moving... frame:%d .xMin:%f xMax:%f\n", frame, minRightX, maxRightX)
	sonicSprite := pixel.NewSprite(pic.picture, pixel.R(minX, 0, maxX, 50))
	movement = pixel.Vec{X: step, Y: 220}
	sonicSprite.Draw(win, pixel.IM.Moved(movement))
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
func loadGamePictures() (SonicSprites, SonicSprites, SonicSprites, SonicSprites, SonicSprites, SonicSprites, SonicSprites, SonicSprites) {
	logoPic, err := loadPicture("sprites/logo.png")
	greenHillPic, err := loadPicture("sprites/green_hill.png")
	sonicStopPic, err := loadPicture("sprites/sonic_stop.png")
	sonicLeftPic, err := loadPicture("sprites/sonic_left.png")
	sonicRightPic, err := loadPicture("sprites/sonic_right.png")
	sonicUpPic, err := loadPicture("sprites/sonic_up.png")
	sonicDownPic, err := loadPicture("sprites/sonic_down.png")
	sonicWaitingPic, err := loadPicture("sprites/sonic_waiting.png")

	if err != nil {
		panic(err)
	}
	return logoPic, greenHillPic, sonicStopPic, sonicLeftPic, sonicRightPic, sonicUpPic, sonicDownPic, sonicWaitingPic
}

/**
Function to load a game Pictures from a path using [os.Open] to get a file, [image] Decode to obtain an image
and from Pixel library [PictureDataFromImage] to create a [Picture] type from [Image]
*/
func loadPicture(path string) (SonicSprites, error) {
	file, err := os.Open(path)
	if err != nil {
		return SonicSprites{}, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return SonicSprites{}, err
	}
	return SonicSprites{pixel.PictureDataFromImage(img)}, nil
}
