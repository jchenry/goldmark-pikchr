package pikchr

import (
	"bytes"
	"fmt"

	"github.com/gebv/pikchr"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/util"
)

// Renderer renders pikchr diagrams as HTML, to be rendered into images
// client side.
type Renderer struct{}

// RegisterFuncs registers the renderer for pikchr blocks with the provided
// Goldmark Registerer.
func (r *Renderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	reg.Register(Kind, r.Render)
}

// Render renders pikchr.Block nodes.
func (r *Renderer) Render(w util.BufWriter, src []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	n := node.(*Block)

	if !entering {
		raw := getLines(src, n)
		res, _ := pikchr.Render(
			string(raw),
			pikchr.HTMLError(),
		)

		_, _ = w.WriteString(fmt.Sprintf(`<div class="pikchr-svg" style="max-width:%dpx">`, res.Width))
		_, _ = w.Write([]byte(res.Data))
		_, _ = w.WriteString("</div>")
	}

	return ast.WalkContinue, nil
}

func getLines(source []byte, n ast.Node) []byte {
	buf := bytes.NewBuffer([]byte{})
	lines := n.Lines()
	for i := 0; i < lines.Len(); i++ {
		line := lines.At(i)
		buf.Write(line.Value(source))
	}
	return buf.Bytes()
}
