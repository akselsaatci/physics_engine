package objects

import "math"

type Constraint interface {
	Apply(object ObjectInterface)
}

type CircleConstraint struct {
	radius float32
	pos_x  float32
	pos_y  float32
}

func (c CircleConstraint) Apply(object ObjectInterface) {
	x, y := object.GetPosition()
	minus := float32(-5.0)
	distance := (x+object.GetWidth() / 2 -c.pos_x-minus)*(x+object.GetWidth() / 2 -c.pos_x-minus) + (y+object.GetWidth() / 2-c.pos_y-minus)*(y+object.GetWidth()/ 2 -minus-c.pos_y)
	if distance > c.radius*c.radius { 
		// TODO
		scale := c.radius / float32(math.Sqrt(float64(distance)))
		new_x := c.pos_x + (x-c.pos_x)*scale
		new_y := c.pos_y + (y-c.pos_y)*scale
		object.updatePosition(new_x, new_y)
	}
}

func MakeNewCircleConstraint(pos_x, pos_y, radius float32) CircleConstraint {
	return CircleConstraint{
		radius: radius,
		pos_x:  pos_x,
		pos_y:  pos_y,
	}
}
