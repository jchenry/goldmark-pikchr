package pikchr

import (
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/util"
)

// Extender adds support for pikchr diagrams to a Goldmark Markdown parser.
//
// Use it by installing it to the goldmark.Markdown object upon creation.
//
//	goldmark.New(
//	  // ...
//	  goldmark.WithExtensions(
//	    // ...
//	    &pikchr.Exender{},
//	  ),
//	)
type Extender struct {
	DarkMode bool
}

// Extend extends the provided Goldmark parser with support for pikchr
// diagrams.
func (e *Extender) Extend(md goldmark.Markdown) {
	md.Parser().AddOptions(
		parser.WithASTTransformers(
			util.Prioritized(&Transformer{}, 100),
		),
	)
	md.Renderer().AddOptions(
		renderer.WithNodeRenderers(
			util.Prioritized(&Renderer{DarkMode: e.DarkMode}, 100),
		),
	)
}
