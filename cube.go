package drawable

import (
	"github.com/canadadry/drawable/geometry"
	"github.com/canadadry/drawable/program"
	"github.com/canadadry/drawable/program/shader"
	"github.com/canadadry/drawable/texture"
	"github.com/go-gl/mathgl/mgl32"
)

type CubeParam struct {
	Pos        mgl32.Vec3
	Size       float32
	Texture    string
	Projection mgl32.Mat4
	Camera     mgl32.Mat4
	Model      mgl32.Mat4
}

type Cube struct {
	Param CubeParam
	p     program.Program
	g     geometry.Geometry
	b     program.Buffer
	t     texture.Texture
}

func (c *Cube) Prepare() error {
	p, err := program.New(shader.Basic)
	if err != nil {
		return err
	}
	p.Use()

	err = p.Uniform(shader.Basic.Uniform[0], c.Param.Projection)
	if err != nil {
		return err
	}
	err = p.Uniform(shader.Basic.Uniform[1], c.Param.Camera)
	if err != nil {
		return err
	}
	err = p.Uniform(shader.Basic.Uniform[2], c.Param.Model)
	if err != nil {
		return err
	}

	t, err := texture.FromImage(c.Param.Texture)
	if err != nil {
		return err
	}

	part := []program.VBOPart{
		{shader.Basic.Attribute[0], 3},
		{shader.Basic.Attribute[1], 2},
	}

	c.p = p
	c.g = geometry.NewCube(c.Param.Pos, c.Param.Size)
	c.b = program.NewBuffer(p, part, c.g.Buf)
	c.t = t

	return nil
}

func (c *Cube) Draw() {
	c.p.Use()
	c.p.Uniform(shader.Basic.Uniform[2], c.Param.Model)
	c.b.Bind()
	c.t.Bind()
	c.g.Draw()
}
