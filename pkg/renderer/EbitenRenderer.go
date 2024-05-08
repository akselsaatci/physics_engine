package renderer

import (
	"github.com/akselsaatci/physics_engine/pkg/objects"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type RendererInterface interface {
	Render(screen *ebiten.Image, objects []objects.ObjectInterface)
}

type EbitenRenderer struct{}

func (renderer *EbitenRenderer) Render(screen *ebiten.Image, objects []objects.ObjectInterface) {
	for _, object := range objects {
		pos_x, pos_y := object.GetPosition()
		vector.DrawFilledCircle(screen, pos_x, pos_y, object.GetWidth(), object.GetColor(), false)
	}
}

func MakeNewEbitenRenderer() *EbitenRenderer {
	return &EbitenRenderer{}
}
