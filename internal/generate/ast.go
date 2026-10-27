package generate

import (
	"bytes"
	"go/format"
	"go/parser"
	"go/token"
	"os"
	"strings"

	"github.com/dave/dst"
	"github.com/dave/dst/decorator"
)

func amendFunctionComments(functionDocs map[string]comment) {
	fileset := token.NewFileSet()
	packages, err := decorator.ParseDir(fileset, ".", nil, parser.ParseComments)
	fatalOnError(err)

	for _, pkg := range packages {
		for filename, file := range pkg.Files {
			if !strings.HasPrefix(filename, "api-") ||
				strings.HasSuffix(filename, "_test.go") {
				continue
			}

			changed := false
			for _, decl := range file.Decls {
				if fn, isFuncDecl := decl.(*dst.FuncDecl); isFuncDecl {
					apiFuncName, found := funcApiFuncName(fn)
					if found {
						fn.Decorations().Start.Clear()

						comment := functionDocs[apiFuncName]
						for _, line := range comment.lines {
							fn.Decorations().Start.Append("// " + line)
						}
						changed = true
					}
				}
			}

			if changed {
				var buf bytes.Buffer
				err = decorator.Fprint(&buf, file)
				fatalOnError(err)

				// Format, because otherwise the decorator package can sometimes produce a
				// file with no newline at the end.
				formatted, err := format.Source(buf.Bytes())
				fatalOnError(err)

				err = os.WriteFile(filename, formatted, 0644)
				fatalOnError(err)
			}
		}
	}
}

type apiFuncNameVisitor struct {
	funcName *string
}

func (v apiFuncNameVisitor) Visit(node dst.Node) dst.Visitor {
	if ident, isIdent := node.(*dst.Ident); isIdent {
		if strings.HasPrefix(ident.Name, "tvg_") {
			*v.funcName = ident.Name
			return nil
		}
	}

	return v
}

func funcApiFuncName(funcNode dst.Node) (string, bool) {
	var apiFuncName string
	visitor := apiFuncNameVisitor{funcName: &apiFuncName}
	dst.Walk(visitor, funcNode)

	return apiFuncName, apiFuncName != ""
}

// type dumpVisitor struct {
// 	indent string
// }

// func (v dumpVisitor) Visit(node dst.Node) dst.Visitor {
// 	fmt.Printf("%s%T\n", v.indent, node)

// 	return dumpVisitor{
// 		indent: v.indent + "  ",
// 	}
// }

// func dumpNode(n dst.Node) {
// 	dst.Walk(dumpVisitor{}, n)
// }
