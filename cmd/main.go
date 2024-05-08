package main

import (
	"image/color"
	"log"

	Engine "github.com/akselsaatci/physics_engine/pkg/engine"
	"github.com/akselsaatci/physics_engine/pkg/helper"
	"github.com/akselsaatci/physics_engine/pkg/objects"
	Renderer "github.com/akselsaatci/physics_engine/pkg/renderer"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	engine   *Engine.Engine
	renderer Renderer.RendererInterface
}


// this is delta time and i hope engine renders at 60fps
// TODO calculate this later
const dt float32 = 1.0 / 60

func (g *Game) Update() error {
    g.engine.CalculateObjectsNextPosition()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
    g.renderer.Render(screen, g.engine.Objects)
}

// TODO should look what is this
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(1412, 832)
	ebiten.SetWindowTitle("Hello, World!")

	r := helper.CustomColor{R: 255, G: 0, B: 0, A: 255}

	var gameObjects []objects.ObjectInterface
	gameObjects = append(gameObjects, objects.MakeNewObject(100, 100, 50, color.White,true))
	gameObjects = append(gameObjects, objects.MakeNewObject(100, 100, 20, r,false))

	eng := Engine.MakeNewEngine(dt, 1.0, gameObjects)
	renderer := Renderer.MakeNewEbitenRenderer()

	game := &Game{engine: eng, renderer: renderer}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
