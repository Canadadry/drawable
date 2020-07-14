package drawable

import (
	"app/pkg/geometry"
	"app/pkg/program"
	"app/pkg/program/shader"
	"app/pkg/texture"
	"github.com/go-gl/mathgl/mgl32"
)

type MeshParam struct {
	Dimmension geometry.MeshDimmension
	Texture    string
	Projection mgl32.Mat4
	Camera     mgl32.Mat4
	Model      mgl32.Mat4
}

type Mesh struct {
	Param MeshParam
	p     program.Program
	g     geometry.Geometry
	b     program.Buffer
	t     texture.Texture
}

func (m *Mesh) Prepare() error {
	p, err := program.New(shader.Basic)
	if err != nil {
		return err
	}
	p.Use()

	err = p.Uniform(shader.Basic.Uniform[0], m.Param.Projection)
	if err != nil {
		return err
	}
	err = p.Uniform(shader.Basic.Uniform[1], m.Param.Camera)
	if err != nil {
		return err
	}
	err = p.Uniform(shader.Basic.Uniform[2], m.Param.Model)
	if err != nil {
		return err
	}

	t, err := texture.FromImage(m.Param.Texture)
	if err != nil {
		return err
	}

	part := []program.VBOPart{
		{shader.Basic.Attribute[0], 3},
		{shader.Basic.Attribute[1], 2},
	}

	m.p = p
	m.g = geometry.NewMesh(m.Param.Dimmension)
	m.b = program.NewBuffer(p, part, m.g.Buf)
	m.t = t

	return nil
}

func (m *Mesh) Draw() {
	m.p.Use()
	m.p.Uniform(shader.Basic.Uniform[2], m.Param.Model)
	m.b.Bind()
	m.t.Bind()
	m.g.Draw()
}
