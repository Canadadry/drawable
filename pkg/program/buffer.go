package program

import (
	"github.com/go-gl/gl/v4.1-core/gl"
)

const (
	float32size = 4
)

type Buffer interface {
	Bind()
}

type buffer struct {
	vao uint32
	vbo uint32
}

func (b buffer) Bind() {
	gl.BindVertexArray(b.vao)
}

type VBOPart struct {
	Name string
	Len  int32
}

func getVBOWidth(desc []VBOPart) int32 {
	width := int32(0)
	for _, p := range desc {
		width += p.Len
	}
	return width
}

func NewBuffer(p Program, desc []VBOPart, buf []float32) buffer {
	w := getVBOWidth(desc)

	b := buffer{}
	gl.GenVertexArrays(1, &b.vao)
	gl.BindVertexArray(b.vao)

	gl.GenBuffers(1, &b.vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, b.vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(buf)*float32size, gl.Ptr(buf), gl.STATIC_DRAW)

	offset := int32(0)
	for _, d := range desc {
		attrib := p.Attribute(d.Name)
		gl.EnableVertexAttribArray(attrib)
		gl.VertexAttribPointer(attrib, d.Len, gl.FLOAT, false, w*float32size, gl.PtrOffset(int(offset*float32size)))

		offset += d.Len
	}
	return b
}
