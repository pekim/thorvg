#! /usr/bin/env bash
set -eo pipefail

THORVG_COMMIT=0a680b13d1753afdb85b498f26a03e605efe7c2e
LIBRARY_FILE=internal/lib/libthorvg-1.so.1.0.0

# clone if not already cloned
if [ ! -e thorvg ]; then
	git clone https://github.com/thorvg/thorvg.git
fi

# ensure desired commit is checked out
pushd thorvg
HEAD=$(git rev-parse HEAD)
if [ "$HEAD" != "$THORVG_COMMIT" ]; then
  git pull
  git checkout "${THORVG_COMMIT}"
fi
popd

# build thorvg with C bindings
pushd thorvg
meson setup build -Dbindings=capi -Dengines=sw,gl
ninja -C build
popd

# copy library and C header
cp thorvg/build/src/libthorvg-1.so.1.0.0 $LIBRARY_FILE
cp thorvg/src/bindings/capi/thorvg_capi.h internal/lib/thorvg_capi.h

# generate a Go file with a hash value of the library
HASH=$(sha256sum $LIBRARY_FILE | cut -d " " -f 1)
echo $HASH
cat << EOF > internal/lib-thorvg-hash.go
package internal

const sharedObjectHash = "$HASH"
EOF
