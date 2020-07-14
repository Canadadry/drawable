package program

import (
	"app/program/shader"
	"fmt"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
	"strings"
)

type Program interface {
	Use()
	Uniform(name string, value interface{}) error
	Attribute(name string) uint32
}

type implProgram struct {
	glId     uint32
	location map[string]int32
}

func (ip implProgram) Use() {
	gl.UseProgram(ip.glId)
}

func (ip implProgram) Uniform(name string, value interface{}) error {
	l, ok := ip.location[name]
	if !ok {
		l = gl.GetUniformLocation(ip.glId, gl.Str(name+"\x00"))
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

func (ip implProgram) Attribute(name string) uint32 {
	return uint32(gl.GetAttribLocation(ip.glId, gl.Str(name+"\x00")))
}

func New(s shader.Shaders) (implProgram, error) {
	vertexShader, err := compileShader(s.Vertex+"\x00", gl.VERTEX_SHADER)
	if err != nil {
		return implProgram{}, err
	}

	fragmentShader, err := compileShader(s.Fragment+"\x00", gl.FRAGMENT_SHADER)
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

	gl.BindFragDataLocation(program, 0, gl.Str(s.Output+"\x00"))

	return implProgram{glId: program, location: map[string]int32{}}, nil
}

func compileShader(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)

	csources, free := gl.Strs(source)
	gl.ShaderSource(shader, 1, csources, nil)
	free()
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to compile %v: %v", source, log)
	}

	return shader, nil
}
