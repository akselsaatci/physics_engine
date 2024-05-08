package Engine

import (
	obj "github.com/akselsaatci/physics_engine/pkg/objects"
)

type EngineInterface interface {
	CalculateObjectsNextPosition()
}

type Engine struct {
	gravity_y float32
	dt        float32
	Objects   []obj.ObjectInterface
}

func (engine *Engine) CalculateObjectsNextPosition() {
    for _, object := range engine.Objects {
        object.CalculateNextPosition(engine.dt, engine.gravity_y)
    }
}

func MakeNewEngine(gravity_y, dt float32, objects []obj.ObjectInterface) *Engine {
    return &Engine{
        gravity_y: gravity_y,
        dt:        dt,
        Objects:   objects,
    }
}

