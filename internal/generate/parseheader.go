package generate

import (
	"os"
	"os/exec"
	"path"
	"strings"
	"sync"

	"github.com/go-clang/clang-v15/clang"
)

var clangResourceDir = sync.OnceValue(func() string {
	out, err := exec.Command("clang", "-print-resource-dir").Output()
	fatalOnError(err)

	resDir := strings.TrimSpace(string(out))
	parts := strings.Split(resDir, "\n")
	resDir = parts[0]

	if resDir == "" {
		fatal("no output when getting clang resource dir")
	}
	if !strings.HasPrefix(resDir, "/") {
		fatalf("expected clang resource dir to start with '/', but it %s", resDir)
	}

	return resDir
})

func parseHeaderFile(filename string) clang.TranslationUnit {
	headerData, err := os.ReadFile(filename)
	fatalOnError(err)
	headerContents := string(headerData)
	srcFile := clang.NewUnsavedFile(filename, string(headerContents))

	resourcesDir := clangResourceDir()
	parseArgs := []string{
		"-I", path.Join(resourcesDir, "include"),
		"-x", "c-header",
	}

	index := clang.NewIndex(0, 1)
	var tu clang.TranslationUnit
	errCode := index.ParseTranslationUnit2(
		filename,
		parseArgs,
		[]clang.UnsavedFile{srcFile},
		uint32(clang.TranslationUnit_SkipFunctionBodies|clang.TranslationUnit_DetailedPreprocessingRecord),
		&tu,
	)
	if errCode != clang.Error_Success {
		fatal(errCode)
	}

	return tu
}
