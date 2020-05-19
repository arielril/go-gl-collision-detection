# Basic Golang OpenGL example

This project uses [Go-Gl](https://github.com/go-gl/gl) to create a scene with OpenGL and play with it.

## Dependencies
This package uses [go modules](https://github.com/golang/go/wiki/Modules). Whenever you run `go build` or `go test` all dependencies will be downloaded automatically. See [go.mod](./go.mod) and [go.sum](go.sum) to see all modules used by this package.

- OpenGL for Go: `github.com/go-gl/gl/v2.1/gl`
- GLFW: `github.com/go-gl/glfw/v3.3/glfw`

## Usage
To execute this example program you can use the `make run` command, it will build the golang executable and run the example program. After the program execution the command will remove the golang executable it has created.

### Commands

- `make run`: build and run the program
- `make`: will build all the files needed

## Result

When running this example program, the program will create an white window and present 2 squares (cars) with spinning sirens (X). One car is blue and the other is red.

You can press `q` or `ESC` to quit the program.
