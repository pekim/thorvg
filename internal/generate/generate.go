package generate

import (
	"github.com/go-clang/clang-v15/clang"
)

func Docs() {
	functionDocs := extractDocComments()
	amendFunctionComments(functionDocs)
}

func extractDocComments() map[string]comment {
	headerFilename := "thorvg_capi.h"

	functionDocs := make(map[string]comment)
	tu := parseHeaderFile(headerFilename)
	tu.TranslationUnitCursor().Visit(func(cursor, _parent clang.Cursor) clang.ChildVisitResult {
		// ignore cursors that are not from the header file
		file, _, _, _ := cursor.Location().FileLocation()
		if file.Name() != headerFilename {
			return clang.ChildVisit_Continue
		}

		switch cursor.Kind() {

		case clang.Cursor_FunctionDecl:
			functionDocs[cursor.Spelling()] = newComment(cursor)
		}

		return clang.ChildVisit_Continue
	})

	return functionDocs
}
