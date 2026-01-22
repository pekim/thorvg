#! /usr/bin/env bash
set -eo pipefail

# change to project's root dir, regardless of the dir where the script is run from
SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
cd $SCRIPT_DIR/..

THORVG_COMMIT=bf5fe14c1eb61377f34a48e49f91be61374c36c6
THORVG_DIR=internal/thorvg-src

# clone if not already cloned
if [ ! -e $THORVG_DIR ]; then
	git clone https://github.com/thorvg/thorvg.git $THORVG_DIR
fi

# ensure desired commit is checked out
pushd $THORVG_DIR
HEAD=$(git rev-parse HEAD)
if [ "$HEAD" != "$THORVG_COMMIT" ]; then
  git fetch
  git checkout "${THORVG_COMMIT}"
fi
popd

# configure thorvg with C bindings
pushd $THORVG_DIR
meson setup build --reconfigure -Dbindings=capi -Dengines=sw,gl -Dsimd=true
popd

# copy required thorvg source files
cp -r $THORVG_DIR/inc internal/cgo/
cp -r $THORVG_DIR/src internal/cgo/

# generate a Go file with the thorvg library commit
HASH=$(sha256sum $LIBRARY_FILE | cut -d " " -f 1)
cat << EOF > lib-thorvg-constant.go
package thorvg

const libthorvgCommit = "$THORVG_COMMIT"
EOF

# generated cgo code
echo -e "\nGenerate code"
rm -f internal/cgo/*.cpp
go run internal/generate/main.go
