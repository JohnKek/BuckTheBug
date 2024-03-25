package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"image/color"
	"log"
)

type Game struct {
	images  []*ebiten.Image
	options []*ebiten.DrawImageOptions
}

const (
	WIDTH int = 320
	HIGH  int = 240
)

func (g *Game) Update() error {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		if len(g.images) == 0 {
			mouseX, mouseY := ebiten.CursorPosition()
			fmt.Printf("%v %v\n", mouseX, mouseY)
			redImage := ebiten.NewImage(5, 5)
			redImage.Fill(color.RGBA{255, 0, 0, 255}) // Красный цвет
			opts := &ebiten.DrawImageOptions{}
			opts.GeoM.Translate(float64(mouseX), float64(mouseY))
			g.images = append(g.images, redImage)
			g.options = append(g.options, opts)
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyW) {
		fmt.Println(*g.images[0])
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for i, img := range g.images {
		opts := g.options[i]
		screen.DrawImage(img, opts)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return WIDTH, HIGH
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
