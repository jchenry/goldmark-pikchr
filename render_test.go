package pikchr

import (
	"bytes"
	"testing"

	"github.com/gebv/pikchr"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
)

// func TestRenderer_Block(t *testing.T) {
// 	t.Parallel()

// 	tests := []struct {
// 		desc string
// 		give string
// 		want string
// 	}{
// 		{
// 			desc: "empty",
// 			give: "",
// 			want: `<div class="pikchr"></div>`,
// 		},
// 		{
// 			desc: "graph",
// 			give: "graph TD;",
// 			want: `<div class="pikchr">graph TD;</div>`,
// 		},
// 		{
// 			desc: "newlines",
// 			give: unlines("foo", "bar"),
// 			want: `<div class="pikchr">foo` + "\nbar" + "\n</div>",
// 		},
// 		{
// 			desc: "escaping",
// 			give: "A -> B",
// 			want: `<div class="pikchr">A -&gt; B</div>`,
// 		},
// 	}

// 	for _, tt := range tests {
// 		tt := tt
// 		t.Run(tt.desc, func(t *testing.T) {
// 			t.Parallel()

// 			r := renderer.NewRenderer(
// 				renderer.WithNodeRenderers(
// 					util.Prioritized(&Renderer{}, 100),
// 				),
// 			)

// 			reader := text.NewReader([]byte(tt.give))
// 			segs := text.NewSegments()
// 			for {
// 				line, seg := reader.PeekLine()
// 				if line == nil {
// 					break
// 				}

// 				segs.Append(seg)
// 				reader.AdvanceLine()
// 			}
// 			give := new(Block)
// 			give.SetLines(segs)

// 			var buff bytes.Buffer
// 			w := bufio.NewWriter(&buff)

// 			assert.NoError(t, r.Render(w, reader.Source(), give), "Render")
// 			assert.Equal(t, tt.want, buff.String())
// 		})
// 	}
// }

func TestDefault(t *testing.T) {
	var buf bytes.Buffer

	backticks := "```"

	raw := `arrow right 200% "Markdown" "Source"
box rad 10px "Markdown" "Formatter" "(markdown.c)" fit
arrow right 200% "HTML+SVG" "Output"
arrow <-> down 70% from last box.s
box same "Pikchr" "Formatter" "(pikchr.c)" fit
`

	source := backticks + ` pikchr
` + raw + backticks
	want := func() string {
		res, _ := pikchr.Render(
			raw,
			pikchr.HTMLError(),
		)
		return res.Data
	}()
	md := goldmark.New(
		goldmark.WithExtensions(
			extension.GFM,
			&Extender{},
		),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(),
	)

	if err := md.Convert([]byte(source), &buf); err != nil {
		t.Error(err)
	}
	if !bytes.Equal(buf.Bytes(), []byte(want)) {
		t.Errorf("got %s, excepted %s\n", buf.Bytes(), want)
	}

	buf.Reset()
}
