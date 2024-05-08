package Engine

import (
	obj "github.com/akselsaatci/physics_engine/pkg/objects"
)

type EngineInterface interface {
	CalculateObjectsNextPosition()
    ApplyConstraints()
}

type Engine struct {
	gravity_y float32
	dt        float32
	Objects   []obj.ObjectInterface
    constraints []obj.Constraint
}

func (engine *Engine) CalculateObjectsNextPosition() {
	for _, object := range engine.Objects {
		object.CalculateNextPosition(engine.dt, engine.gravity_y)
	}
}

func (engine *Engine) ApplyConstraints() {
	for _, object := range engine.Objects {
		for _, constraint := range engine.constraints {
			constraint.Apply(object)
		}
	}
}

func MakeNewEngine(gravity_y, dt float32, objects []obj.ObjectInterface,constraints []obj.Constraint) *Engine {
	return &Engine{
		gravity_y:   gravity_y,
		dt:          dt,
		Objects:     objects,
		constraints: constraints,
	}
}
