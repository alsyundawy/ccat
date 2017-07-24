package main

import (
	"fmt"
	"io"
	// "sort"
	// "strings"

	"github.com/nsf/termbox-go"
	"github.com/sourcegraph/syntaxhighlight"
)

var (
	stringKind        = kind{"String", syntaxhighlight.String}
	keywordKind       = kind{"Keyword", syntaxhighlight.Keyword}
	commentKind       = kind{"Comment", syntaxhighlight.Comment}
	typeKind          = kind{"Type", syntaxhighlight.Type}
	literalKind       = kind{"Literal", syntaxhighlight.Literal}
	punctuationKind   = kind{"Punctuation", syntaxhighlight.Punctuation}
	plaintextKind     = kind{"Plaintext", syntaxhighlight.Plaintext}
	tagKind           = kind{"Tag", syntaxhighlight.Tag}
	htmlTagKind       = kind{"HTMLTag", syntaxhighlight.HTMLTag}
	htmlAttrNameKind  = kind{"HTMLAttrName", syntaxhighlight.HTMLAttrName}
	htmlAttrValueKind = kind{"HTMLAttrValue", syntaxhighlight.HTMLAttrValue}
	decimalKind       = kind{"Decimal", syntaxhighlight.Decimal}

	kinds = []kind{
		stringKind,
		keywordKind,
		commentKind,
		typeKind,
		literalKind,
		punctuationKind,
		plaintextKind,
		tagKind,
		htmlTagKind,
		htmlAttrNameKind,
		htmlAttrValueKind,
		decimalKind,
	}

	LightColorPalettes = ColorPalettes{
		stringKind:        termbox.ColorYellow,
		keywordKind:       termbox.ColorBlue,
		commentKind:       termbox.Attribute(243),
		typeKind:          termbox.ColorMagenta,
		literalKind:       termbox.ColorMagenta,
		punctuationKind:   termbox.ColorRed,
		plaintextKind:     termbox.ColorRed,
		tagKind:           termbox.ColorBlue,
		htmlTagKind:       termbox.ColorGreen,
		htmlAttrNameKind:  termbox.ColorBlue,
		htmlAttrValueKind: termbox.ColorGreen,
		decimalKind:       termbox.ColorBlue,
	}

	DarkColorPalettes = ColorPalettes{
		stringKind:        termbox.Attribute(20),
		keywordKind:       termbox.ColorBlue,
		commentKind:       termbox.Attribute(200),
		typeKind:          termbox.Attribute(2000),
		literalKind:       termbox.Attribute(2000),
		punctuationKind:   termbox.ColorRed,
		plaintextKind:     termbox.ColorBlue,
		tagKind:           termbox.ColorBlue,
		htmlTagKind:       termbox.Attribute(20000),
		htmlAttrNameKind:  termbox.ColorBlue,
		htmlAttrValueKind: termbox.ColorGreen,
		decimalKind:       termbox.ColorBlue,
	}

	// cache kind name and syntax highlight kind
	// for faster lookup
	kindsByName map[string]kind
	kindsByKind map[syntaxhighlight.Kind]kind
)

func init() {
	kindsByName = make(map[string]kind)
	for _, k := range kinds {
		kindsByName[k.Name] = k
	}

	kindsByKind = make(map[syntaxhighlight.Kind]kind)
	for _, k := range kinds {
		kindsByKind[k.Kind] = k
	}
}

type kind struct {
	Name string
	Kind syntaxhighlight.Kind
}

type ColorPalettes map[kind]termbox.Attribute

func (c ColorPalettes) Set(k string, v termbox.Attribute) bool {
	kind, ok := kindsByName[k]
	if ok {
		c[kind] = v
	}

	return ok
}

func (c ColorPalettes) Get(k syntaxhighlight.Kind) termbox.Attribute {
	// ignore whitespace kind
	if k == syntaxhighlight.Whitespace {
		return termbox.ColorRed
	}

	kind, ok := kindsByKind[k]
	if !ok {
		panic(fmt.Sprintf("Unknown syntax highlight kind %d\n", k))
	}

	return c[kind]
}

// func (c ColorPalettes) String() string {
// 	var s []string
// 	for _, k := range kinds {
// 		color := c[k]
// 		s = append(s, fmt.Sprintf("%13s\t%s", k.Name, Colorize(color, color)))
// 	}
//
// 	return strings.Join(s, "\n")
// }

func CPrint(r io.Reader, w io.Writer, palettes ColorPalettes) error {
	return syntaxhighlight.Print(
		syntaxhighlight.NewScannerReader(r),
		w,
		Printer{palettes},
	)
}

type Printer struct {
	ColorPalettes ColorPalettes
}

func (p Printer) Print(w io.Writer, kind syntaxhighlight.Kind, tokText string) error {
	// c := p.ColorPalettes.Get(kind)
	// if len(c) > 0 {
	// 	tokText = Colorize(c, tokText)
	// }

	_, err := io.WriteString(w, tokText)

	return err
}
