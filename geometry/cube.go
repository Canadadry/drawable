package geometry

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

func NewCube(at mgl32.Vec3, size float32) Geometry {
	cubeVertices := []float32{
		//  X, Y, Z, U, V
		// Bottom
		-size + at[0], -size + at[1], -size + at[2], 0.0, 0.0,
		size + at[0], -size + at[1], -size + at[2], 1.0, 0.0,
		-size + at[0], -size + at[1], size + at[2], 0.0, 1.0,
		size + at[0], -size + at[1], -size + at[2], 1.0, 0.0,
		size + at[0], -size + at[1], size + at[2], 1.0, 1.0,
		-size + at[0], -size + at[1], size + at[2], 0.0, 1.0,

		// Top
		-size + at[0], size + at[1], -size + at[2], 0.0, 0.0,
		-size + at[0], size + at[1], size + at[2], 0.0, 1.0,
		size + at[0], size + at[1], -size + at[2], 1.0, 0.0,
		size + at[0], size + at[1], -size + at[2], 1.0, 0.0,
		-size + at[0], size + at[1], size + at[2], 0.0, 1.0,
		size + at[0], size + at[1], size + at[2], 1.0, 1.0,

		// Front
		-size + at[0], -size + at[1], size + at[2], 1.0, 0.0,
		size + at[0], -size + at[1], size + at[2], 0.0, 0.0,
		-size + at[0], size + at[1], size + at[2], 1.0, 1.0,
		size + at[0], -size + at[1], size + at[2], 0.0, 0.0,
		size + at[0], size + at[1], size + at[2], 0.0, 1.0,
		-size + at[0], size + at[1], size + at[2], 1.0, 1.0,

		// Back
		-size + at[0], -size + at[1], -size + at[2], 0.0, 0.0,
		-size + at[0], size + at[1], -size + at[2], 0.0, 1.0,
		size + at[0], -size + at[1], -size + at[2], 1.0, 0.0,
		size + at[0], -size + at[1], -size + at[2], 1.0, 0.0,
		-size + at[0], size + at[1], -size + at[2], 0.0, 1.0,
		size + at[0], size + at[1], -size + at[2], 1.0, 1.0,

		// Left
		-size + at[0], -size + at[1], size + at[2], 0.0, 1.0,
		-size + at[0], size + at[1], -size + at[2], 1.0, 0.0,
		-size + at[0], -size + at[1], -size + at[2], 0.0, 0.0,
		-size + at[0], -size + at[1], size + at[2], 0.0, 1.0,
		-size + at[0], size + at[1], size + at[2], 1.0, 1.0,
		-size + at[0], size + at[1], -size + at[2], 1.0, 0.0,

		// Right
		size + at[0], -size + at[1], size + at[2], 1.0, 1.0,
		size + at[0], -size + at[1], -size + at[2], 1.0, 0.0,
		size + at[0], size + at[1], -size + at[2], 0.0, 0.0,
		size + at[0], -size + at[1], size + at[2], 1.0, 1.0,
		size + at[0], size + at[1], -size + at[2], 0.0, 0.0,
		size + at[0], size + at[1], size + at[2], 0.0, 1.0,
	}
	return Geomtry{
		Buf:   cubeVertices,
		Width: 5,
		Mode:  gl.TRIANGLES,
	}

}
