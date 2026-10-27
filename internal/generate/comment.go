package generate

import (
	"fmt"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"github.com/go-clang/clang-v15/clang"
)

type comment struct {
	cursor  clang.Cursor
	rawText string
	lines   []string
	inAtSee bool
}

func newComment(cursor clang.Cursor) comment {
	c := comment{
		cursor:  cursor,
		rawText: cursor.RawCommentText(),
	}

	var builder strings.Builder
	c.processComment(&builder, c.cursor.ParsedComment())
	text := builder.String()
	lines := strings.Split(text, "\n")
	for i, line := range lines {
		if len(line) > 1 && line[0] == ' ' && line[1] != ' ' {
			// Remove a lone space at the start of the line.
			// This is a space that follows the asterisk in a block comment
			// line sucha as this "* comment text...".
			lines[i] = line[1:]
		}

		// There are some block comments for api sections that immediately
		// precede a doxygen comment. These result in lines with lots of asterisks
		// appearing at the start of a clang parsed comment.
		if strings.HasPrefix(line, "**********") && strings.TrimSpace(strings.ReplaceAll(line, "*", " ")) == "" {
			lines[i] = ""
		}
	}

	c.lines = lines

	return c
}

func (c *comment) processComment(builder *strings.Builder, comment clang.Comment) {
	switch comment.Kind() {

	case clang.Comment_BlockCommand:
		switch comment.BlockCommandComment_getCommandName() {
		case "brief":
			c.processChildren(builder, comment)
			builder.WriteRune('\n')

		case "deprecated":
			builder.WriteString("\nDeprecate:")
			c.processChildren(builder, comment)

		case "note":
			builder.WriteString("\nNote:")
			c.processChildren(builder, comment)

		case "return":
			builder.WriteString("\nReturn:")
			c.processChildren(builder, comment)

		case "retval":
			retval := comment.BlockCommandComment_getArgText(0)
			retval = strings.TrimSuffix(retval, ".")
			// All @retval instances refer to enum members.
			memberName := strings.TrimPrefix(retval, "TVG_")
			fmt.Fprintf(builder, "\n\nReturns [%s] :", memberName)
			c.processChildren(builder, comment)

		case "see":
			c.inAtSee = true
			c.processChildren(builder, comment)

		case "since":
			builder.WriteString("Since:")
			c.processChildren(builder, comment)

		case "warning":
			builder.WriteString("\nWarning:")
			c.processChildren(builder, comment)

		default:
			fmt.Println(c.rawText)
			fatalf("unhandled block command : %s\n", comment.BlockCommandComment_getCommandName())
		}

	case clang.Comment_InlineCommand:
		switch comment.InlineCommandComment_getCommandName() {
		case "c": // fixed font
			builder.WriteString(comment.InlineCommandComment_getArgText(0))

		case "n": // newline
			builder.WriteRune('\n')
			builder.WriteString(comment.InlineCommandComment_getArgText(0))

		case "p": // fixed font
			builder.WriteString(comment.InlineCommandComment_getArgText(0))

		case "ref":
			ref := comment.InlineCommandComment_getArgText(0)
			builder.WriteString(c.resolveReference(ref))

		default:
			fmt.Println(c.rawText)
			fatalf("unhandled inline command : %s\n", comment.InlineCommandComment_getCommandName())
		}

	case clang.Comment_ParamCommand:
		name := comment.ParamCommandComment_getParamName()
		fmt.Fprintf(builder, "  - %s :", name)
		c.processChildren(builder, comment)

	case clang.Comment_FullComment:
		c.processChildren(builder, comment)

	case clang.Comment_Null:
		c.processChildren(builder, comment)

	case clang.Comment_Paragraph:
		c.processChildren(builder, comment)
		builder.WriteRune('\n')

	case clang.Comment_Text:
		commentText := comment.TextComment_getText()
		if commentText != "" {
			if c.inAtSee {
				c.inAtSee = false
				fmt.Fprintf(builder, "\n\nsee %s", c.resolveReference(commentText))
			} else {
				builder.WriteString(commentText)
			}
		}
		c.processChildren(builder, comment)

	case clang.Comment_VerbatimLine:
		c.processChildren(builder, comment)

	default:
		fmt.Println(c.rawText)
		fatalf("unhandled comment type : %s\n", comment.Kind().Spelling())
	}
}

func (c *comment) processChildren(builder *strings.Builder, comment clang.Comment) {
	for i := range comment.NumChildren() {
		c.processComment(builder, comment.Child(i))
	}
}

var classes = []string{
	"accessor",
	"animation",
	"canvas",
	"conic_gradient",
	"linear_gradient",
	"paint",
	"picture",
	"radial_gradient",
	"saver",
	"scene",
	"shape",
	"text",
	"video",
}

var titleCaser = cases.Title(language.English, cases.NoLower)

func (c *comment) resolveReference(ref string) string {
	ref = strings.TrimSpace(ref)
	if !strings.HasPrefix(ref, "tvg_") {
		return ref
	}

	ref = strings.TrimPrefix(ref, "tvg_")
	ref = strings.TrimSuffix(ref, ".")
	ref = strings.TrimSuffix(ref, "()")
	if strings.HasPrefix(ref, "lottie_animation") {
		ref = strings.TrimPrefix(ref, "lottie_")
	}

	for _, class := range classes {
		if strings.HasPrefix(ref, class+"_") {
			classname := goName(class)
			methodName := goName(strings.TrimPrefix(ref, class+"_"))

			if methodName != "" {
				return fmt.Sprintf("[%s.%s]", classname, methodName)
			}
		}
	}

	return fmt.Sprintf("[%s]", goName(ref))
}

func goName(name string) string {
	parts := strings.Split(name, "_")
	for i, part := range parts {
		parts[i] = titleCaser.String(part)
	}
	return strings.Join(parts, "")
}

func (c comment) Dump() {
	c.dumpComment(c.cursor.ParsedComment(), "")
}

func (c comment) dumpComment(comment clang.Comment, indent string) {
	fmt.Println(indent, comment.Kind().Spelling(), "=>", comment.TextComment_getText())

	switch comment.Kind() {
	case clang.Comment_BlockCommand:
		fmt.Println(indent, "  Command : Block :", comment.BlockCommandComment_getCommandName())
	case clang.Comment_InlineCommand:
		fmt.Println(indent, "  Command : Inline :", comment.InlineCommandComment_getCommandName())
	case clang.Comment_ParamCommand:
		fmt.Println(indent, "  Command : Param :", comment.ParamCommandComment_getDirection(), comment.ParamCommandComment_getParamName())

	}

	for i := range comment.NumChildren() {
		c.dumpComment(comment.Child(i), indent+"  ")
	}
}
