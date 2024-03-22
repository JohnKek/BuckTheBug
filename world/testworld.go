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

func (g *Game) Update() error {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		mouseX, mouseY := ebiten.CursorPosition()
		fmt.Printf("%v %v\n", mouseX, mouseY)

		// Создание изображения с красным пикселем в координатах mouseX, mouseY
		redImage := ebiten.NewImage(1, 1)
		redImage.Fill(color.RGBA{255, 0, 0, 255}) // Красный цвет

		// Преобразование координат из окна в координаты изображения
		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Translate(float64(mouseX), float64(mouseY))

		// Добавление изображения и его опций в соответствующие списки для отображения
		g.images = append(g.images, redImage)
		g.options = append(g.options, opts)
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
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
