#! /usr/bin/env bash
set -eo pipefail

# change to project's root dir, regardless of the dir where the script is run from
SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
cd $SCRIPT_DIR/..

THORVG_COMMIT=1a43240ec3ffdaa689412e7cd52e83cf8118e2b9
THORVG_DIR=internal/thorvg-src
GOOS=$(go env GOOS)
GOARCH=$(go env GOARCH)
LIBRARY_DIR=internal/lib
LIBRARY_FILE=$LIBRARY_DIR/libthorvg_${GOOS}_${GOARCH}

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
meson setup build -Dbindings=capi -Dengines=sw,gl -Dsimd=true
ninja -C build
popd

# copy library and C header
cp $THORVG_DIR/build/src/libthorvg-1.so.1.0.0 $LIBRARY_FILE
cp $THORVG_DIR/src/bindings/capi/thorvg_capi.h $LIBRARY_DIR

# strip symbols from the library, reducing the size
# from ~11M to ~1.2M
strip $LIBRARY_FILE

# generate a Go file that embeds the GOOS/GOARCH specific library
cat << EOF > lib-thorvg_${GOOS}_${GOARCH}.go
package thorvg

import _ "embed"

//go:embed $LIBRARY_FILE
var sharedObject []byte
EOF

# generate a Go file with a hash value of the library
HASH=$(sha256sum $LIBRARY_FILE | cut -d " " -f 1)
cat << EOF > lib-thorvg-constant.go
package thorvg

const sharedObjectHash = "$HASH"
const libthorvgCommit = "$THORVG_COMMIT"
EOF
