package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

const (
	width  = 160
	height = 144
	scale  = 10
)

var (
	pixelData [width][height][3]byte
)

func update(screen *ebiten.Image) error {
	p := make([]byte, width*height*4)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			position := y*width + x
			p[(position*4 + 0)] = pixelData[x][y][0]
			p[(position*4 + 1)] = pixelData[x][y][1]
			p[(position*4 + 2)] = pixelData[x][y][2]
			p[(position*4 + 3)] = 255
		}
	}
	screen.ReplacePixels(p)
	ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %f", ebiten.CurrentFPS()))
	return nil
}

func main() {
	pixelData[0][0][0] = 255
	pixelData[0][0][1] = 255
	pixelData[0][0][2] = 255
	pixelData[width-1][height-1][0] = 255
	pixelData[width-1][height-1][1] = 255
	pixelData[width-1][height-1][2] = 255
	pixelData[width/2][height/2][0] = 255
	pixelData[width/2][height/2][1] = 255
	pixelData[width/2][height/2][2] = 255
	if err := ebiten.Run(update, width, height, scale, "go-ebiten-example"); err != nil {
		panic(err)
	}
}
