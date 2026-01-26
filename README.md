# thorvg

This library provides Go bindings for the
[thorvg](https://www.thorvg.org/) library.
There is no use of cgo.

## image loaders

The only supported image loaders are those for svg and lottie files.
For other image formats use Go's image packages to load the raw data.

## AI

No AI was used in the creation of this library.

## Platforms

The following platforms are currently suppported.

- `linux/amd64`
- `macos/amd64`
- `macos/arm64`

## no cgo

thorvg is a C++ library, with an optional C api.
The https://github.com/pekim/thorvg-binaries repo has a github worklfow
that builds thorvg for multiple platforms, with the C api.
For each platform a shared library file is produced.
The library files, such as `libthorvg_linux_amd64`, are copied in to this repo.

When an application uses this library, one of the library files is included at build time
courtesy of Go's embedding support.
At runtime the embed library is written to a temporary file.
The [purego](https://github.com/ebitengine/purego) library is then used to open the file and
register its api functions.

## development

### upgrading thorvg

See https://github.com/pekim/thorvg-binaries for details of how to
upgrade the thorvg library.

### pre-commit hook

- install `goimports` if not already installed
  - https://pkg.go.dev/golang.org/x/tools/cmd/goimports
- install `golangci-lint` (v2.x) if not already installed
  - https://golangci-lint.run/docs/welcome/install/#binaries
- install the `pre-commit` application if not already installed
  - https://pre-commit.com/index.html#install
- install pre-commit hook in this repo's workspace
  - `pre-commit install`
