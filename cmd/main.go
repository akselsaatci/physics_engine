package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Game struct {
	Circles []*CircleObject
}

const gravity_y float32 = 1

const dt float32 = 1.0 / 60

type CircleObject struct {
	pos_x_old      float32
	pos_y_old      float32
	pos_x          float32
	pos_y          float32
	width          float32
	color          color.Color
	acceleration_x float32
	acceleration_y float32
}

func (object *CircleObject) Update() {
	vel_x := object.pos_x - object.pos_x_old
	vel_y := object.pos_y - object.pos_y_old

	object.pos_x_old = object.pos_x
	object.pos_y_old = object.pos_y

	object.pos_x = object.pos_x + vel_x + object.acceleration_x*dt*dt
	object.pos_y = object.pos_y + vel_y + object.acceleration_y*dt*dt

	fmt.Println(dt)
	object.acceleration_x = 0
	object.acceleration_y = 0
}
func (object *CircleObject) accelerate(acc_x, acc_y float32) {
	object.acceleration_x = acc_x
	object.acceleration_y = acc_y

}

func (object *CircleObject) Draw(screen *ebiten.Image) {
	vector.DrawFilledCircle(screen, object.pos_x, object.pos_y, object.width, object.color, false)
}

func MakeNewCircleObject(loc_x, loc_y, width float32, color color.Color) *CircleObject {
	return &CircleObject{
		pos_x_old:      loc_x,
		pos_y_old:      loc_y,
		pos_x:          loc_x,
		pos_y:          loc_y,
		width:          width,
		color:          color,
		acceleration_x: 0,
		acceleration_y: 0,
	}
}
func (g *Game) Update() error {
	for _, circle := range g.Circles {
		circle.accelerate(0, gravity_y)
		circle.Update()
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, circle := range g.Circles {
		circle.Draw(screen)

	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(1412, 832)
	ebiten.SetWindowTitle("Hello, World!")

	game := &Game{
		Circles: []*CircleObject{
			MakeNewCircleObject(160, 120, 5, color.White),
		},
	}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
