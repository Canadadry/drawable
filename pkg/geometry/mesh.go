package geometry

import (
	"github.com/go-gl/gl/v4.1-core/gl"
)

type MeshDimmension struct {
	Width      int
	Height     int
	CellWidth  int
	CellHeight int
}

const floatPerVertice = 5
const verticePerTriangle = 3
const verticePerCell = verticePerTriangle * 2

func NewMesh(md MeshDimmension) Geometry {
	vertices := make([]float32, 0, md.Width*md.Height*verticePerCell*floatPerVertice)
	var left, right, top, bottom, z float32
	z = 0
	for y := 0; y < md.Height; y++ {
		bottom = float32(y * md.Height)
		top = bottom + float32(md.Height)
		for x := 0; x < md.Width; x++ {
			left = float32(x * md.Width)
			right = left + float32(md.Width)

			vertices = append(vertices,
				//  X, Y, Z, U, V
				left, top, z, 0, 1,
				right, top, z, 1, 1,
				left, bottom, z, 0, 0,
				left, bottom, z, 0, 0,
				right, top, z, 1, 1,
				right, bottom, z, 1, 0,
			)
		}
	}

	return Geometry{
		Buf:   vertices,
		Width: floatPerVertice,
		Mode:  gl.TRIANGLES,
	}
}
