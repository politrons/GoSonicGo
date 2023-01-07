package main

import (
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

type EnemySprites struct {
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
	sonicWaitingPic,
	sonicBallPic,
	sonicJumpPic = loadGamePictures()

var flyPic = loadEnemyPictures()

func main() {
	pixelgl.Run(run)
}

func run() {
	win := createWindow()
	greenHillBackground := pixel.NewSprite(greenHillPic.picture, pixel.R(-1200, -800, 1200, 800))
	logoPicture := pixel.NewSprite(logoPic.picture, pixel.R(-824, -768, 1300, 768))
	//startIntroTheme()
	var logoTime = time.Now().Second() + introTime
	for time.Now().Second() < logoTime {
		win.Update()
		win.Clear(colornames.Grey)
		logoPicture.Draw(win, pixel.IM)
	}
	go startMainTheme()
	go gravity()
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
		go animateEnemies(win)
		if win.Pressed(pixelgl.KeyDown) && win.Pressed(pixelgl.KeySpace) {
			checkLastKeyPressed(pixelgl.KeyDown)
			ballAnimation(win)
			ySonic = ySonic + 2
			resetWaitTime()
		} else if win.Pressed(pixelgl.KeySpace) {
			checkLastKeyPressed(pixelgl.KeySpace)
			jumpAnimation(win)
			ySonic = ySonic + 10
			resetWaitTime()
		} else if win.Pressed(pixelgl.KeyRight) {
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
			if originalY < ySonic {
				jumpAnimation(win)
			} else if (waitingTime + 3) < time.Now().Second() {
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
func loadGamePictures() (SonicSprites, SonicSprites, SonicSprites, SonicSprites, SonicSprites, SonicSprites, SonicSprites, SonicSprites, SonicSprites, SonicSprites) {
	logoPic, err := loadPicture("sprites/logo.png")
	greenHillPic, err := loadPicture("sprites/green_hill.png")
	sonicStopPic, err := loadPicture("sprites/sonic_stop.png")
	sonicLeftPic, err := loadPicture("sprites/sonic_left.png")
	sonicRightPic, err := loadPicture("sprites/sonic_right.png")
	sonicUpPic, err := loadPicture("sprites/sonic_up.png")
	sonicDownPic, err := loadPicture("sprites/sonic_down.png")
	sonicWaitingPic, err := loadPicture("sprites/sonic_waiting.png")
	sonicBallPic, err := loadPicture("sprites/sonic_ball.png")
	sonicJumpPic, err := loadPicture("sprites/sonic_jump.png")

	if err != nil {
		panic(err)
	}
	return logoPic, greenHillPic, sonicStopPic, sonicLeftPic, sonicRightPic, sonicUpPic, sonicDownPic, sonicWaitingPic, sonicBallPic, sonicJumpPic
}

func loadEnemyPictures() EnemySprites {
	flyPic, err := loadPicture("sprites/fly.png")

	if err != nil {
		panic(err)
	}
	return EnemySprites{flyPic.picture}
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
