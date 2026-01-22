# thorvg

This library provides Go bindings for the
[thorvg](https://www.thorvg.org/) library.

## image loaders

The only supported image loaders are those for svg and lottie files.
For other image formats use Go's image packages to load the raw data.

## AI

No AI was used in the creation of this library.

## Platforms

`linux/amd64` was used for intial development, and is currently the only supported platform.

## Linker errors

If using the `mold` linker, then the linker will emit many warnings when
building the glfw examples. They are annoying but harmless.

The warnings look like this.

```
  mold: warning: symbol type mismatch: glDisableVertexAttribArray
  >>> defined in /tmp/go-link-1307360853/000266.o as STT_OBJECT
  >>> defined in /usr/lib64/libGL.so as STT_FUNC
```

## development

### updating thorvg

- Update the value of the `THORVG_COMMIT` variable in the script if required.
- Run `go generate` to execute the shell script `./internal/build.sh`.

### pre-commit hook

- install `goimports` if not already installed
  - https://pkg.go.dev/golang.org/x/tools/cmd/goimports
- install `golangci-lint` (v2.x) if not already installed
  - https://golangci-lint.run/docs/welcome/install/#binaries
- install the `pre-commit` application if not already installed
  - https://pre-commit.com/index.html#install
- install pre-commit hook in this repo's workspace
  - `pre-commit install`
