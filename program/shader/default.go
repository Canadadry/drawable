package shader

type Shaders struct {
	Vertex    string
	Fragment  string
	Output    string
	Uniform   []string
	Attribute []string
}

var Basic = Shaders{
	Vertex: `
#version 330

uniform mat4 projection;
uniform mat4 camera;
uniform mat4 model;

in vec3 vert;
in vec2 vertTexCoord;

out vec2 fragTexCoord;

void main() {
    fragTexCoord = vertTexCoord;
    gl_Position = projection * camera * model * vec4(vert, 1);
}`,
	Fragment: `
#version 330

uniform sampler2D tex;

in vec2 fragTexCoord;

out vec4 outputColor;

void main() {
    outputColor = texture(tex, fragTexCoord);
}`,
	Output:    "outputColor",
	Uniform:   []string{"projection", "camera", "model"},
	Attribute: []string{"vert", "vertTexCoord"},
}
