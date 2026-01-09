# thorvg

## updating thorvg

The shell script `build.sh` can be used to update the thorvg
shared object and header files.

Update the value of the `THORVG_COMMIT` variable in the script if required.
Then run the script `./internal/build.sh`.

## pre-commit hook

- install `goimports` if not already installed
  - https://pkg.go.dev/golang.org/x/tools/cmd/goimports
- install `golangci-lint` (v2.x) if not already installed
  - https://golangci-lint.run/docs/welcome/install/#binaries
- install the `pre-commit` application if not already installed
  - https://pre-commit.com/index.html#install
- install pre-commit hook in this repo's workspace
  - `pre-commit install`
