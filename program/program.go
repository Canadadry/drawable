package program

import (
	"fmt"
	"github.com/go-gl/gl/v4.1-core/gl"
	"strings"
)

type Program interface {
	Use()
}

type implProgram struct {
	GlId uint32
}

func (ip implProgram) Use() {
	gl.UseProgram(ip.GlId)
}

func New(vertexShaderSource, fragmentShaderSource string) (implProgram, error) {
	vertexShader, err := compileShader(vertexShaderSource, gl.VERTEX_SHADER)
	if err != nil {
		return implProgram{}, err
	}

	fragmentShader, err := compileShader(fragmentShaderSource, gl.FRAGMENT_SHADER)
	if err != nil {
		return implProgram{}, err
	}

	program := gl.CreateProgram()

	gl.AttachShader(program, vertexShader)
	gl.AttachShader(program, fragmentShader)
	gl.LinkProgram(program)

	var status int32
	gl.GetProgramiv(program, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(program, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(program, logLength, nil, gl.Str(log))

		return implProgram{}, fmt.Errorf("failed to link program: %v", log)
	}

	gl.DeleteShader(vertexShader)
	gl.DeleteShader(fragmentShader)

	return implProgram{GlId: program}, nil
}
