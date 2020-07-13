package program

import (
	"fmt"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
	"strings"
)

type Program interface {
	Use()
	Uniform(name string, value interface{})
}

type implProgram struct {
	GlId     uint32
	location map[string]int32
}

func (ip implProgram) Use() {
	gl.UseProgram(ip.GlId)
}

func (ip implProgram) Uniform(name string, value interface{}) error {
	l, ok := ip.location[name]
	if !ok {
		l = gl.GetUniformLocation(ip.GlId, gl.Str(name+"\x00"))
		if l < 0 {
			return fmt.Errorf("Cannont found %s in program", name)
		}
		ip.location[name] = l
	}
	switch value.(type) {
	case mgl32.Mat4:
		v, _ := value.(mgl32.Mat4)
		gl.UniformMatrix4fv(l, 1, false, &v[0])
	default:
		return fmt.Errorf("Un handled type %T", value)
	}
	return nil
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

	return implProgram{GlId: program, location: map[string]int32{}}, nil
}
