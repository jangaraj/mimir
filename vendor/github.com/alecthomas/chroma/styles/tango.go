package styles

import (
	"github.com/alecthomas/chroma"
)

// Tango style.
var Tango = Register(chroma.MustNewStyle("tango", chroma.StyleEntries{
	chroma.TextWhitespace:           "underline #f8f8f8",
	chroma.Error:                    "#a40000 border:#ef2929",
	chroma.Other:                    "#000000",
	chroma.Comment:                  "italic #8f5902",
	chroma.CommentMultiline:         "italic #8f5902",
	chroma.CommentPreproc:           "italic #8f5902",
	chroma.CommentSingle:            "italic #8f5902",
	chroma.CommentSpecial:           "italic #8f5902",
	chroma.Keyword:                  "bold #204a87",
	chroma.KeywordConstant:          "bold #204a87",
	chroma.KeywordDeclaration:       "bold #204a87",
	chroma.KeywordNamespace:         "bold #204a87",
	chroma.KeywordPseudo:            "bold #204a87",
	chroma.KeywordReserved:          "bold #204a87",
	chroma.KeywordType:              "bold #204a87",
	chroma.Operator:                 "bold #ce5c00",
	chroma.OperatorWord:             "bold #204a87",
	chroma.Punctuation:              "bold #000000",
	chroma.Name:                     "#000000",
	chroma.NameAttribute:            "#c4a000",
	chroma.NameBuiltin:              "#204a87",
	chroma.NameBuiltinPseudo:        "#3465a4",
	chroma.NameClass:                "#000000",
	chroma.NameConstant:             "#000000",
	chroma.NameDecorator:            "bold #5c35cc",
	chroma.NameEntity:               "#ce5c00",
	chroma.NameException:            "bold #cc0000",
	chroma.NameFunction:             "#000000",
	chroma.NameProperty:             "#000000",
	chroma.NameLabel:                "#f57900",
	chroma.NameNamespace:            "#000000",
	chroma.NameOther:                "#000000",
	chroma.NameTag:                  "bold #204a87",
	chroma.NameVariable:             "#000000",
	chroma.NameVariableClass:        "#000000",
	chroma.NameVariableGlobal:       "#000000",
	chroma.NameVariableInstance:     "#000000",
	chroma.LiteralNumber:            "bold #0000cf",
	chroma.LiteralNumberFloat:       "bold #0000cf",
	chroma.LiteralNumberHex:         "bold #0000cf",
	chroma.LiteralNumberInteger:     "bold #0000cf",
	chroma.LiteralNumberIntegerLong: "bold #0000cf",
	chroma.LiteralNumberOct:         "bold #0000cf",
	chroma.Literal:                  "#000000",
	chroma.LiteralDate:              "#000000",
	chroma.LiteralString:            "#4e9a06",
	chroma.LiteralStringBacktick:    "#4e9a06",
	chroma.LiteralStringChar:        "#4e9a06",
	chroma.LiteralStringDoc:         "italic #8f5902",
	chroma.LiteralStringDouble:      "#4e9a06",
	chroma.LiteralStringEscape:      "#4e9a06",
	chroma.LiteralStringHeredoc:     "#4e9a06",
	chroma.LiteralStringInterpol:    "#4e9a06",
	chroma.LiteralStringOther:       "#4e9a06",
	chroma.LiteralStringRegex:       "#4e9a06",
	chroma.LiteralStringSingle:      "#4e9a06",
	chroma.LiteralStringSymbol:      "#4e9a06",
	chroma.Generic:                  "#000000",
	chroma.GenericDeleted:           "#a40000",
	chroma.GenericEmph:              "italic #000000",
	chroma.GenericError:             "#ef2929",
	chroma.GenericHeading:           "bold #000080",
	chroma.GenericInserted:          "#00A000",
	chroma.GenericOutput:            "italic #000000",
	chroma.GenericPrompt:            "#8f5902",
	chroma.GenericStrong:            "bold #000000",
	chroma.GenericSubheading:        "bold #800080",
	chroma.GenericTraceback:         "bold #a40000",
	chroma.GenericUnderline:         "underline",
	chroma.Background:               " bg:#f8f8f8",
}))
