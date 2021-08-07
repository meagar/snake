package main

import (
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/meagar/snake"
)

const (
	screenWidth  = float64(1024.0)
	screenHeight = float64(768.0)
)

func main() {
	ebiten.SetWindowSize(int(screenWidth), int(screenHeight))
	ebiten.SetWindowTitle("Snake")
	if err := ebiten.RunGame(snake.NewGame(screenWidth, screenHeight)); err != nil {
		log.Fatal(err)
	}
}
