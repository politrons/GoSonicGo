package main

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

func animateEnemies(win *pixelgl.Window) {
	animateFly(win)
}

func animateFly(win *pixelgl.Window) {
	if flyFrame < 3 {
		minFlyX = minFlyX + pixelFlyRect
		maxFlyX = maxFlyX + pixelFlyRect
		flyFrame = flyFrame + 1
	} else {
		flyFrame = 0
		minFlyX = 0
		maxFlyX = pixelFlyRect
	}
	xFlyEnemy = xFlyEnemy - 5
	flyPic.drawPicture(win, minFlyX, maxFlyX)
}

/**
Implementation of [SonicSprites] to draw the specific sprite in the window.
Using interface each group of sprites invokes this method, avoiding type mismatch
*/
func (pic EnemySprites) drawPicture(win *pixelgl.Window, minX float64, maxX float64) {
	fmt.Printf("Moving... EnemyFrame:%d .xMin:%f xMax:%f\n", flyFrame, minX, maxX)
	enemySprite := pixel.NewSprite(pic.picture, pixel.R(minX, 0, maxX, 50))
	movement := pixel.Vec{X: xFlyEnemy, Y: yFlyEnemy}
	enemySprite.Draw(win, pixel.IM.Moved(movement))
}
