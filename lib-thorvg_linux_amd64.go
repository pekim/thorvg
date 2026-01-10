package thorvg

import _ "embed"

//go:embed internal/lib/libthorvg_linux_amd64
var sharedObject []byte
