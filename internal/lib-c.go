package internal

import (
	"github.com/ebitengine/purego"
)

var free func(*byte)
var malloc func(length int) *byte

func initLibc() error {
	lib, err := purego.Dlopen("libc.so.6", purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		return err
	}

	purego.RegisterLibFunc(&free, lib, "free")
	purego.RegisterLibFunc(&malloc, lib, "malloc")

	return nil
}
