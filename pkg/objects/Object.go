package objects

import (
	"image/color"
)

type ObjectInterface interface {
	Accelerate(acc_x, acc_y float32)
	updatePosition(new_x, new_y float32)
	GetPosition() (X float32, Y float32)
	resetAcceleration()
	CalculateNextPosition(dt float32, gravity_y float32)
	GetColor() color.Color
	GetWidth() float32
}

type Object struct {
	pos_x_old        float32
	pos_y_old        float32
	pos_x            float32
	pos_y            float32
	width            float32
	color            color.Color // TODO could change to a hex code
	acceleration_x   float32
	acceleration_y   float32
	mass             float32 // TODO
	isGravityApplied bool
}

func (object *Object) Accelerate(acc_x, acc_y float32) {
	object.acceleration_x += acc_x
	object.acceleration_y += acc_y
}

func (object *Object) updatePosition(new_x, new_y float32) {
	object.pos_x_old = object.pos_x
	object.pos_y_old = object.pos_y
	object.pos_x = new_x
	object.pos_y = new_y
}
func (object *Object) GetPosition() (float32, float32) {
	return object.pos_x, object.pos_y
}

func (object *Object) GetColor() color.Color {
	return object.color
}

func (object *Object) GetWidth() float32 {
	return object.width
}

func (object *Object) CalculateNextPosition(dt float32, gravity_y float32) {

	if object.isGravityApplied {
		object.Accelerate(0, gravity_y)
	}

	vel_x := object.pos_x - object.pos_x_old
	vel_y := object.pos_y - object.pos_y_old

	calc_x := object.pos_x + vel_x + object.acceleration_x*dt*dt
	calc_y := object.pos_y + vel_y + object.acceleration_y*dt*dt

	object.updatePosition(calc_x, calc_y)

	object.resetAcceleration()

}

func (object *Object) resetAcceleration() {
	object.acceleration_x = 0
	object.acceleration_y = 0
}

func MakeNewObject(loc_x, loc_y, width float32, color color.Color, isGravityApplied bool) *Object {
	return &Object{
		pos_x_old:        loc_x,
		pos_y_old:        loc_y,
		pos_x:            loc_x,
		pos_y:            loc_y,
		width:            width,
		color:            color,
		acceleration_x:   0,
		acceleration_y:   0,
		mass:             1,
		isGravityApplied: isGravityApplied,
	}
}
