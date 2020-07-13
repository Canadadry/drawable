package texture

import (
	"github.com/go-gl/gl/v4.1-core/gl"
)

type Texture interface {
	Bind()
}

type texture struct {
	glId uint32
}

func (t texture) Bind() {
	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, t.glId)
}
