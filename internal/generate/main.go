package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type introTargets []introTarget

type introTarget struct {
	TargetSources []struct {
		Compiler   []string `json:"compiler"`
		Linker     []string `json:"linker"`
		Parameters []string `json:"parameters"`
		Sources    []string `json:"sources"`
	} `json:"target_sources"`
}

func main() {
	bytes, err := os.ReadFile("internal/thorvg-src/build/meson-info/intro-targets.json")
	if err != nil {
		panic(err)
	}
	var introTargets introTargets
	err = json.Unmarshal(bytes, &introTargets)
	if err != nil {
		panic(err)
	}
	if len(introTargets) != 1 {
		panic(fmt.Sprintf("expected 1 intro target, but have %d\n", len(introTargets)))
	}
	introTarget := introTargets[0]

	copyConfigHeader()
	generateCppFiles(introTarget)
	generateCgo(introTarget)
}

func copyConfigHeader() {
	bytes, err := os.ReadFile("internal/thorvg-src/build/config.h")
	if err != nil {
		panic(err)
	}
	content := string(bytes)

	// trim multiple trailing newlines, to satisfy end-of-file-fixer pre-commit hook
	for strings.HasSuffix(content, "\n\n") {
		content = strings.TrimSuffix(content, "\n")
	}

	err = os.WriteFile("internal/cgo/config.h", []byte(content), 0644)
	if err != nil {
		panic(err)
	}
}

func generateCppFiles(introTarget introTarget) {
	workingDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	srcDir := filepath.Join(workingDir, "internal/thorvg-src/src") + "/"

	for _, targetSource := range introTarget.TargetSources {
		for _, source := range targetSource.Sources {
			sourceFile := strings.TrimPrefix(source, srcDir)
			outPath := filepath.Join("internal/cgo", strings.ReplaceAll(sourceFile, "/", "_"))
			err := os.WriteFile(outPath, []byte(fmt.Sprintf("#include \"%s\"\n", sourceFile)), 0644)
			if err != nil {
				panic(err)
			}
		}
	}
}

func generateCgo(introTarget introTarget) {
	workingDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	srcDir := filepath.Join(workingDir, "internal/thorvg-src")

	var flagLines []string
	for _, targetSource := range introTarget.TargetSources {
		if len(targetSource.Compiler) > 0 {
			for _, param := range targetSource.Parameters {
				if strings.HasPrefix(param, "-I") {
					file := param
					file = strings.TrimPrefix(file, "-I")
					file = strings.TrimPrefix(file, srcDir)
					if !strings.HasPrefix(file, "/build") {
						flagLines = append(flagLines, fmt.Sprintf("#cgo CXXFLAGS: -I${SRCDIR}%s", file))
					}
				} else {
					// Skip parameters that are not supported by cgo.
					// At least not supported without use of CGO_CXXFLAGS_ALLOW environment variable.
					if strings.HasPrefix(param, "-fdiagnostics-color") ||
						strings.HasPrefix(param, "-fno-math-errno") ||
						strings.HasPrefix(param, "-fno-unwind-tables") ||
						strings.HasPrefix(param, "-DTEST_DIR") {
						continue
					}

					flagLines = append(flagLines, fmt.Sprintf("#cgo CXXFLAGS: %s", param))
				}
			}
		}

		if len(targetSource.Linker) > 0 {
			for _, param := range targetSource.Parameters {
				if strings.Contains(param, "-soname") {
					continue
				}
				flagLines = append(flagLines, fmt.Sprintf("#cgo LDFLAGS: %s", param))
			}
		}

		flagLines = append(flagLines, "")
	}

	template := `
package cgo

/*
%s
*/
import "C"
`

	content := fmt.Sprintf(template,
		strings.Join(flagLines, "\n"),
	)

	err = os.WriteFile("internal/cgo/cgo.go", []byte(content), 0644)
	if err != nil {
		panic(err)
	}

	cmd := exec.Command("goimports", "-w", "internal/cgo/cgo.go")
	output, err := cmd.CombinedOutput()
	if len(output) > 0 {
		fmt.Println(string(output))
	}
	if err != nil {
		panic(err)
	}
}
