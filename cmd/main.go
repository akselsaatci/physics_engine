package main

import (
	"image/color"
	"log"

	"github.com/akselsaatci/physics_engine/pkg/objects"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	Circles []*objects.CircleObject
}

const gravity_y float32 = 1

// this is delta time and i hope engine renders at 60fps
// TODO calculate this later
const dt float32 = 1.0 / 60

func (g *Game) Update() error {
	for _, circle := range g.Circles {
		circle.Accelerate(0, gravity_y)
		circle.Update()
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, circle := range g.Circles {
		circle.Draw(screen)

	}
}

// TODO should look what is this
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(1412, 832)
	ebiten.SetWindowTitle("Hello, World!")

//	r := helper.CustomColor{R: 255, G: 0, B: 0, A: 255}

	game := &Game{
		Circles: []*objects.CircleObject{
			objects.MakeNewCircleObject(160, 120, 5, color.White),
		},
	}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
