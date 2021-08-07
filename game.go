package snake

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	Snake        Snake
	screenWidth  int
	screenHeight int

	cameraX float64
	cameraY float64

	keys []ebiten.Key

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
		Snake: Snake{
			size:    5,
			heading: 0,
			speed:   1,
		},
	}

	g.Snake.head = &g.Snake.body[0]
	fmt.Println("Virtual size", w, h)
	g.loadAssets()

	return &g
}

func (g *Game) Update() error {
	g.keys = inpututil.PressedKeys()
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		g.Snake.TurnLeft()
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		g.Snake.TurnRight()
	}

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		g.Snake.Grow()
	}

	g.Snake.Move()

	// Seek the camera towards the snake
	// TODO: scale to tps
	g.cameraX += ((g.Snake.head.X - g.cameraX) / 20.0)
	g.cameraY += ((g.Snake.head.Y - g.cameraY) / 20.0)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.drawBackground(screen)
	g.drawSnake(&g.Snake, screen)
	g.drawDebug(screen)
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
	// red := color.RGBA{255, 0, 0, 64}
	// ebitenutil.DrawLine(g.bgTile, 0, 0, float64(w), 0, red)
	// ebitenutil.DrawLine(g.bgTile, float64(w), 0, float64(w), float64(h), red)
	// ebitenutil.DrawLine(g.bgTile, float64(w), float64(h), 0, float64(h), red)
	// ebitenutil.DrawLine(g.bgTile, 0, float64(h), 0, 0, red)
	// op.GeoM.Scale()
	// g.bgTile.DrawImage(tile, &op)
}

func (g *Game) drawDebug(screen *ebiten.Image) {
	keys := []string{}
	for _, p := range g.keys {
		keys = append(keys, p.String())
	}
	ebitenutil.DebugPrint(screen, fmt.Sprintf("Snake: %0.2f, %0.f, %0.2f\nCamera: %0.2f,%0.2f\nKeys: %s\nFPS: %0.2f",
		g.Snake.head.X, g.Snake.head.Y, g.Snake.size,
		g.cameraX, g.cameraY,
		strings.Join(keys, ", "),
		ebiten.CurrentFPS(),
	))
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
	// x += 1
	// Draw the tiled background
	w, h := g.bgTile.Size()
	// scale := 1 / (g.Snake.size * 0.01)
	// fw, fh := float64(w)*scale, float64(h)*scale
	// w, h = int(fw), int(fh)
	maxX, maxY := screen.Size()
	ops := ebiten.DrawImageOptions{}
	for x := int(-g.cameraX)%w - w; x < maxX; x += w {
		for y := int(-g.cameraY)%h - h; y < maxY; y += h {
			ops.GeoM.Reset()
			// ops.GeoM.Scale(scale, scale)
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

func (g *Game) drawSnake(s *Snake, screen *ebiten.Image) {
	red := loadSprite("red.png")
	w, h := red.Size()
	fw, fh := float64(w), float64(h)
	scale := 0.25 + (s.size * 0.001)
	scaledWidth := fw * scale
	scaledHeight := fh * scale

	ops := ebiten.DrawImageOptions{}
	ops.Filter = ebiten.FilterLinear
	for i := g.Snake.Length() - 1; i >= 0; i-- {
		p := g.Snake.body[i]
		ops.GeoM.Reset()
		ops.GeoM.Scale(scale, scale)
		ops.GeoM.Translate(-scaledWidth/2, -scaledHeight/2)

		// Move the segment to its world coordinates
		ops.GeoM.Translate(p.X, p.Y)
		// Move the snake back into screen coordinates
		ops.GeoM.Translate(-g.cameraX, -g.cameraY)
		// Center the camera
		ops.GeoM.Translate(float64(g.screenWidth/2), float64(g.screenHeight/2))

		screen.DrawImage(red, &ops)
	}

}
