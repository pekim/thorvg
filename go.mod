module github.com/pekim/thorvg

go 1.26.0

// Address deprecation of clang_getDiagnosticCategoryName.
// Avoids a noisy message when generating api.
replace github.com/go-clang/clang-v15 => github.com/pekim/clang-v15 v0.0.0-20240830114552-c0d27ccce9ec

require (
	github.com/dave/dst v0.28.0
	github.com/ebitengine/purego v0.11.1
	github.com/go-clang/clang-v15 v0.0.0-20230222085438-ee3102fa0c71
	github.com/go-gl/glfw/v3.3/glfw v0.0.0-20260823155953-d41da22a9587
	github.com/jupiterrider/purego-sdl3 v0.0.0-20260514083405-8523da70a041
	github.com/stretchr/testify v1.12.1
	golang.org/x/sys v0.49.0
	golang.org/x/text v0.43.0
)

require (
	go.yaml.in/yaml/v3 v3.0.5 // indirect
	golang.org/x/mod v0.41.0 // indirect
	golang.org/x/sync v0.24.0 // indirect
	golang.org/x/tools v0.51.0 // indirect
)
