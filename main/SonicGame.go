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

func main() {
	pixelgl.Run(run)
}

func run() {

	win := createWindow()
	greenHillPic, sonicStopPic, sonicRightPic := loadGamePictures()
	greenHillBackground := drawBackground(greenHillPic, win)

	var step = float64(300)
	var movement = pixel.Vec{step, step}
	var frame = 0
	var pixelRec = float64(36.6)
	var minX = float64(0)
	var maxX = pixelRec

	for !win.Closed() {
		win.Update()
		win.Clear(colornames.Grey)
		greenHillBackground.Draw(win, pixel.IM)
		if win.Pressed(pixelgl.KeyRight) {
			if frame < 9 {
				minX = minX + pixelRec
				maxX = maxX + pixelRec
			} else {
				frame = 1
				minX = 0
				maxX = pixelRec
			}
			frame = frame + 1
			fmt.Printf("Moving......xMin:%f xMax:%f\n", minX, maxX)
			sonicSprite := pixel.NewSprite(sonicRightPic, pixel.R(minX, 0, maxX, 50))

			step = step + 5
			movement = pixel.Vec{step, 220}
			sonicSprite.Draw(win, pixel.IM.Moved(movement))
		} else {
			fmt.Printf("Stoped......xMin:%f xMax:%f\n", minX, maxX)
			movement = pixel.Vec{step, 220}
			sonicSprite := pixel.NewSprite(sonicStopPic, pixel.R(0, 0, pixelRec, 50))
			sonicSprite.Draw(win, pixel.IM.Moved(movement))
		}
		time.Sleep(50 * time.Millisecond)
	}

}

/**
Create the background sprite and draw in the window with the dimensions specify in pixel.R Rect
*/
func drawBackground(greenHillPic pixel.Picture, win *pixelgl.Window) *pixel.Sprite {
	greenHillBackground := pixel.NewSprite(greenHillPic, pixel.R(-1200, -800, 1200, 800))
	greenHillBackground.Draw(win, pixel.IM)
	return greenHillBackground
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
func loadGamePictures() (pixel.Picture, pixel.Picture, pixel.Picture) {
	greenHillPic, err := loadPicture("green_hill.png")
	sonicStopPic, err := loadPicture("sonic_stop.png")
	sonicRightPic, err := loadPicture("sonic_right.png")
	if err != nil {
		panic(err)
	}
	return greenHillPic, sonicStopPic, sonicRightPic
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
