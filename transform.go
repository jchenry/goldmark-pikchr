package pikchr

import (
	"bytes"

	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
)

// Transformer transforms a Goldmark Markdown AST with support for pikchr
// diagrams by replacing pikchr code blocks with [Block] nodes.
type Transformer struct{}

var _pikchr = []byte("pikchr")

// Transform transforms the provided Markdown AST.
func (*Transformer) Transform(doc *ast.Document, reader text.Reader, pctx parser.Context) {
	var (
		pikchrBlocks []*ast.FencedCodeBlock
	)

	// Collect all blocks to be replaced without modifying the tree.
	_ = ast.Walk(doc, func(node ast.Node, enter bool) (ast.WalkStatus, error) {
		if !enter {
			return ast.WalkContinue, nil
		}

		cb, ok := node.(*ast.FencedCodeBlock)
		if !ok {
			return ast.WalkContinue, nil
		}

		lang := cb.Language(reader.Source())
		if !bytes.Equal(lang, _pikchr) {
			return ast.WalkContinue, nil
		}

		pikchrBlocks = append(pikchrBlocks, cb)
		return ast.WalkContinue, nil
	})

	// Nothing to do.
	if len(pikchrBlocks) == 0 {
		return
	}

	for _, cb := range pikchrBlocks {
		b := new(Block)
		b.SetLines(cb.Lines())

		parent := cb.Parent()
		if parent != nil {
			parent.ReplaceChild(parent, cb, b)
		}
	}
}
