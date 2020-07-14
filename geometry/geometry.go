package geometry

import (
	"github.com/go-gl/gl/v4.1-core/gl"
)

type Geomtry struct {
	Buf   []float32
	Width int32
	Mode  uint32
}

func (g Geomtry) Draw() {
	gl.DrawArrays(g.Mode, 0, int32(len(g.Buf))/g.Width)
}
