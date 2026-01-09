#! /usr/bin/env bash
set -eo pipefail

# change to project's root dir, regardless of the dir where the script is run from
SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
cd $SCRIPT_DIR/..

THORVG_COMMIT=0a680b13d1753afdb85b498f26a03e605efe7c2e
THORVG_DIR=internal/thorvg-src
LIBRARY_DIR=internal/lib
LIBRARY_FILE=$LIBRARY_DIR/libthorvg-1.so.1.0.0

# clone if not already cloned
if [ ! -e $THORVG_DIR ]; then
	git clone https://github.com/thorvg/thorvg.git $THORVG_DIR
fi

# ensure desired commit is checked out
pushd $THORVG_DIR
HEAD=$(git rev-parse HEAD)
if [ "$HEAD" != "$THORVG_COMMIT" ]; then
  git pull
  git checkout "${THORVG_COMMIT}"
fi
popd

# build thorvg with C bindings
pushd $THORVG_DIR
meson setup build -Dbindings=capi -Dengines=sw,gl
ninja -C build
popd

# copy library and C header
cp $THORVG_DIR/build/src/libthorvg-1.so.1.0.0 $LIBRARY_DIR
cp $THORVG_DIR/src/bindings/capi/thorvg_capi.h $LIBRARY_DIR

# strip symbols from the library, reducing the size
# from ~11M to ~1.2M
strip $LIBRARY_FILE

# generate a Go file with a hash value of the library
HASH=$(sha256sum $LIBRARY_FILE | cut -d " " -f 1)
echo $HASH
cat << EOF > internal/lib-thorvg-constant.go
package internal

const sharedObjectHash = "$HASH"
const libthorvgCommit = "$THORVG_COMMIT"
EOF
