package styles

import (
	"github.com/alecthomas/chroma"
)

// Colorful style.
var Colorful = Register(chroma.MustNewStyle("colorful", chroma.StyleEntries{
	chroma.TextWhitespace:        "#bbbbbb",
	chroma.Comment:               "#888",
	chroma.CommentPreproc:        "#579",
	chroma.CommentSpecial:        "bold #cc0000",
	chroma.Keyword:               "bold #080",
	chroma.KeywordPseudo:         "#038",
	chroma.KeywordType:           "#339",
	chroma.Operator:              "#333",
	chroma.OperatorWord:          "bold #000",
	chroma.NameBuiltin:           "#007020",
	chroma.NameFunction:          "bold #06B",
	chroma.NameClass:             "bold #B06",
	chroma.NameNamespace:         "bold #0e84b5",
	chroma.NameException:         "bold #F00",
	chroma.NameVariable:          "#963",
	chroma.NameVariableInstance:  "#33B",
	chroma.NameVariableClass:     "#369",
	chroma.NameVariableGlobal:    "bold #d70",
	chroma.NameConstant:          "bold #036",
	chroma.NameLabel:             "bold #970",
	chroma.NameEntity:            "bold #800",
	chroma.NameAttribute:         "#00C",
	chroma.NameTag:               "#070",
	chroma.NameDecorator:         "bold #555",
	chroma.LiteralString:         "bg:#fff0f0",
	chroma.LiteralStringChar:     "#04D bg:",
	chroma.LiteralStringDoc:      "#D42 bg:",
	chroma.LiteralStringInterpol: "bg:#eee",
	chroma.LiteralStringEscape:   "bold #666",
	chroma.LiteralStringRegex:    "bg:#fff0ff #000",
	chroma.LiteralStringSymbol:   "#A60 bg:",
	chroma.LiteralStringOther:    "#D20",
	chroma.LiteralNumber:         "bold #60E",
	chroma.LiteralNumberInteger:  "bold #00D",
	chroma.LiteralNumberFloat:    "bold #60E",
	chroma.LiteralNumberHex:      "bold #058",
	chroma.LiteralNumberOct:      "bold #40E",
	chroma.GenericHeading:        "bold #000080",
	chroma.GenericSubheading:     "bold #800080",
	chroma.GenericDeleted:        "#A00000",
	chroma.GenericInserted:       "#00A000",
	chroma.GenericError:          "#FF0000",
	chroma.GenericEmph:           "italic",
	chroma.GenericStrong:         "bold",
	chroma.GenericPrompt:         "bold #c65d09",
	chroma.GenericOutput:         "#888",
	chroma.GenericTraceback:      "#04D",
	chroma.GenericUnderline:      "underline",
	chroma.Error:                 "#F00 bg:#FAA",
	chroma.Background:            " bg:#ffffff",
}))
