# thorvg

## building thorvg

```sh
git clone https://github.com/thorvg/thorvg.git
cd thorvg
meson setup build -Dbindings=capi -Dengines=sw,gl
ninja -C build
cp build/src/libthorvg-1.so.1.0.0 $THIS_REPO_DIR/internal/lib/
cp src/bindings/capi/thorvg_capi.h $THIS_REPO_DIR/internal/lib/
```

## pre-commit hook

- install `goimports` if not already installed
  - https://pkg.go.dev/golang.org/x/tools/cmd/goimports
- install `golangci-lint` (v2.x) if not already installed
  - https://golangci-lint.run/docs/welcome/install/#binaries
- install the `pre-commit` application if not already installed
  - https://pre-commit.com/index.html#install
- install pre-commit hook in this repo's workspace
  - `pre-commit install`
