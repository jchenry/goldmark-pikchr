package pikchr

import (
	"bytes"

	"github.com/gebv/pikchr"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/util"
)

// const _defaultpikchrJS = "https://cdn.jsdelivr.net/npm/pikchr/dist/pikchr.min.js"

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
	// if entering {
	// 	w.WriteString(`<div class="pikchr">`)
	// 	lines := n.Lines()
	// 	for i := 0; i < lines.Len(); i++ {
	// 		line := lines.At(i)
	// 		template.HTMLEscape(w, line.Value(src))
	// 	}
	// } else {
	// 	w.WriteString("</div>")
	// }

	if !entering {
		raw := r.getLines(src, n)
		// h := sha1.New()
		// h.Write(raw)
		// hash := string(h.Sum([]byte{}))
		// if result, exist := u.buf[hash]; exist {
		// _, _ = w.Write(result)
		// } else {
		res, _ := pikchr.Render(
			string(raw),
			pikchr.HTMLError(),
		)

		svg := []byte(res.Data)
		// if len(u.buf) >= u.MaxLength {
		// u.buf = make(map[string][]byte)
		// }
		// u.buf[hash] = svg
		_, _ = w.Write(svg)
		// }
	}

	return ast.WalkContinue, nil
}

func (r *Renderer) getLines(source []byte, n ast.Node) []byte {
	buf := bytes.NewBuffer([]byte{})
	l := n.Lines().Len()
	for i := 0; i < l; i++ {
		line := n.Lines().At(i)
		buf.Write(line.Value(source))
	}
	return buf.Bytes()
}
