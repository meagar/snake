package snake

import (
	"fmt"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"math/rand"
	"strings"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func init() {
	seed := time.Now().Unix()
	fmt.Println("Seeding rand:", seed)
	rand.Seed(seed)
}

type Pill struct {
	Point
	Size  float64
	Eaten bool
}

type Game struct {
	Snake        Snake
	screenWidth  int
	screenHeight int

	cameraX float64
	cameraY float64

	pills []Pill
	keys  []ebiten.Key

	// var red = loadSprite("red.png")
	bgTile *ebiten.Image
}

func NewGame(screenWidth, screenHeight float64) *Game {
	ratio := screenWidth / screenHeight
	w := 800
	h := int(600 / ratio)
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

	g.pills = make([]Pill, 1000)
	for i := range g.pills {
		g.pills[i].X = (rand.Float64() * 4000) - 2000
		g.pills[i].Y = (rand.Float64() * 4000) - 2000
		g.pills[i].Size = rand.Float64() * 10
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

	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		g.Snake.Grow()
	}

	// for i := range g.pills {
	// 	if g.pills[i].Eaten == false {
	// 		if math.Abs(g.Snake.head.X-g.pills[i].X) < (g.Snake.size+10) && math.Abs(g.Snake.head.Y-g.pills[i].Y) < (g.Snake.size+10) {
	// 			g.pills[i].Eaten = true
	// 			g.Snake.Grow()
	// 			g.Snake.Grow()
	// 			g.Snake.Grow()
	// 		}
	// 	}
	// }

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
	// g.drawPills(screen)
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

func (g *Game) drawPills(screen *ebiten.Image) {
	op := ebiten.DrawImageOptions{}
	red := loadSprite("red.png")
	for i := range g.pills {
		if g.pills[i].Eaten == false {
			op.GeoM.Reset()
			op.GeoM.Scale(0.2, 0.2)
			op.GeoM.Translate(-10, -10)
			op.GeoM.Translate(float64(g.screenWidth/2), float64(g.screenHeight/2))
			op.GeoM.Translate(g.pills[i].X-g.cameraX, g.pills[i].Y-g.cameraY)
			screen.DrawImage(red, &op)
		}
	}
}

func loadSprite(name string) *ebiten.Image {
	img, _, err := ebitenutil.NewImageFromFile("assets/" + name)
	fmt.Println("Loading asset", name)
	// fh, err := os.OpenFile("assets/"+name, os.O_RDONLY, 0600)
	if err != nil {
		panic(err)
	}
	return img
	// img, _, err := image.Decode(fh)
	// if err != nil {
	// 	panic(err)
	// }
	// return ebiten.NewImageFromImage(img)
}

func createSphere() *ebiten.Image {
	c := color.RGBA{
		uint8(rand.Intn(200)),
		uint8(rand.Intn(200)),
		uint8(rand.Intn(200)),
		255}
	mask := loadSprite("circle_mask.png")
	w, h := mask.Size()
	dest := ebiten.NewImage(w, h)
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			_, _, _, a := mask.At(x, y).RGBA()
			c.A = byte(a)
			// c.R = 255
			dest.Set(x, y, c)
		}
	}

	shadow := loadSprite("shadow.png")
	dest.DrawImage(shadow, nil)

	highlight := loadSprite("highlight.png")
	dest.DrawImage(highlight, nil)
	return dest
}

// var x float64
var px, py float64

var drawBackgroundOp ebiten.DrawImageOptions

func (g *Game) drawBackground(screen *ebiten.Image) {
	// x += 1
	// Draw the tiled background
	w, h := g.bgTile.Size()
	// scale := 1 / (g.Snake.size * 0.01)
	// fw, fh := float64(w)*scale, float64(h)*scale
	// w, h = int(fw), int(fh)
	maxX, maxY := screen.Size()
	for x := int(-g.cameraX)%w - w; x < maxX; x += w {
		for y := int(-g.cameraY)%h - h; y < maxY; y += h {
			drawBackgroundOp.GeoM.Reset()
			// ops.GeoM.Scale(scale, scale)
			drawBackgroundOp.GeoM.Translate(float64(x), float64(y))
			screen.DrawImage(g.bgTile, &drawBackgroundOp)
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

var drawSnakeOp ebiten.DrawImageOptions

var red *ebiten.Image

func (g *Game) drawSnake(s *Snake, screen *ebiten.Image) {
	if red == nil {
		red = createSphere() //loadSprite("red.png")
	}
	w, h := red.Size()
	fw, fh := float64(w), float64(h)
	scale := 0.25 + (s.size * 0.001)
	scaledWidth := fw * scale
	scaledHeight := fh * scale

	drawSnakeOp.Filter = ebiten.FilterLinear
	for i := g.Snake.Length() - 1; i >= 0; i-- {
		p := g.Snake.body[i]
		drawSnakeOp.GeoM.Reset()
		drawSnakeOp.GeoM.Scale(scale, scale)
		drawSnakeOp.GeoM.Translate(-scaledWidth/2, -scaledHeight/2)

		// Move the segment to its world coordinates
		drawSnakeOp.GeoM.Translate(p.X, p.Y)
		// Move the snake back into screen coordinates
		drawSnakeOp.GeoM.Translate(-g.cameraX, -g.cameraY)
		// Center the camera
		drawSnakeOp.GeoM.Translate(float64(g.screenWidth/2), float64(g.screenHeight/2))

		screen.DrawImage(red, &drawSnakeOp)
	}

}
