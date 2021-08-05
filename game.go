package snake

import (
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Point struct {
	x, y float32
}

type Snake struct {
	Head  Point
	parts []*Point
}

type Game struct {
	Snakes       []*Snake
	screenWidth  int
	screenHeight int

	// var red = loadSprite("red.png")
	bgTile *ebiten.Image
}

func NewGame(screenWidth, screenHeight float64) *Game {
	ratio := screenWidth / screenHeight
	w := 1000
	h := int(1000 / ratio)
	// w := int(screenWidth)
	// h := int(screenHeight)
	g := Game{
		screenWidth:  w,
		screenHeight: h,
	}

	fmt.Println("Virtual size", w, h)
	g.loadAssets()

	return &g
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.drawBackground(screen)
	g.drawSnakes(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.screenWidth, g.screenHeight
}

func (g *Game) loadAssets() {
	tile := loadSprite("bg_tile.png")
	w, h := tile.Size()
	w /= 2
	h /= 2
	g.bgTile = ebiten.NewImage(w, h)
	op := ebiten.DrawImageOptions{}
	op.GeoM.Scale(0.5, 0.5)
	op.Filter = ebiten.FilterLinear
	g.bgTile.DrawImage(tile, &op)
	red := color.RGBA{255, 0, 0, 64}
	ebitenutil.DrawLine(g.bgTile, 0, 0, float64(w), 0, red)
	ebitenutil.DrawLine(g.bgTile, float64(w), 0, float64(w), float64(h), red)
	ebitenutil.DrawLine(g.bgTile, float64(w), float64(h), 0, float64(h), red)
	ebitenutil.DrawLine(g.bgTile, 0, float64(h), 0, 0, red)
	// op.GeoM.Scale()
	// g.bgTile.DrawImage(tile, &op)
}

func loadSprite(name string) *ebiten.Image {
	fh, err := os.OpenFile("assets/"+name, os.O_RDONLY, 0600)
	if err != nil {
		panic(err)
	}
	img, _, err := image.Decode(fh)
	if err != nil {
		panic(err)
	}
	return ebiten.NewImageFromImage(img)
}

// var x float64
var px, py float64

func (g *Game) drawBackground(screen *ebiten.Image) {
	px -= 1
	py -= 1
	// x += 1
	// Draw the tiled background
	w, h := g.bgTile.Size()
	maxX, maxY := screen.Size()
	ops := ebiten.DrawImageOptions{}
	for x := int(px)%w - w; x < maxX; x += w {
		for y := int(py)%h - h; y < maxY; y += h {
			ops.GeoM.Reset()
			ops.GeoM.Translate(float64(x), float64(y))
			screen.DrawImage(g.bgTile, &ops)
		}
	}
	// ops.GeoM.Scale(0.5, 0.5)

	// w, _ := g.bgTile.Size()
	// ops = ebiten.DrawImageOptions{}
	// ops.GeoM.Apply(0.5, 0.5)
	// fmt.Println(x)
	// ops.GeoM.Scale(0.5, 0.5)
	// ops.GeoM.Translate(float64(w)/2, 0)
	// screen.DrawImage(tile, &ops)
}

func (g *Game) drawSnakes(screen *ebiten.Image) {
	// ops := ebiten.DrawImageOptions{}
	// ops.GeoM.Scale(0.5, 0.5)
	// screen.DrawImage(red, &ops)
}
