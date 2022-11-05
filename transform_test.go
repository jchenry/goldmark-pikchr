package pikchr

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
	"github.com/yuin/goldmark/util"
)

func TestTransformer(t *testing.T) {
	t.Parallel()

	tests := []struct {
		desc       string
		give       string
		wantBodies []string
	}{
		{
			desc: "empty",
			give: "",
		},
		{
			desc: "pikchr",
			give: unlines(
				"```pikchr",
				"foo",
				"```",
			),
			wantBodies: []string{"foo\n"},
		},
		{
			desc: "pikchr and not",
			give: unlines(
				"Foo",
				"",
				"```pikchr",
				"foo",
				"```",
				"",
				"Bar",
				"",
				"```go",
				"bar",
				"",
				"Baz",
				"",
			),
			wantBodies: []string{"foo\n"},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.desc, func(t *testing.T) {
			t.Parallel()

			p := goldmark.New().Parser()
			p.AddOptions(
				parser.WithASTTransformers(
					util.Prioritized(&Transformer{}, 100),
				),
			)

			src := []byte(tt.give)
			got := p.Parse(text.NewReader(src))

			var gotBodies []string
			_ = ast.Walk(got, func(node ast.Node, enter bool) (ast.WalkStatus, error) {
				if !enter {
					return ast.WalkContinue, nil
				}

				switch n := node.(type) {
				case *Block:
					gotBodies = append(gotBodies, string(getLines(src, n)))
				}

				return ast.WalkContinue, nil
			})

			assert.Equal(t, tt.wantBodies, gotBodies)
		})
	}
}

func TestTransformer_RepeatedTransformations(t *testing.T) {
	t.Parallel()

	src := []byte(unlines(
		"```pikchr",
		"foo",
		"```",
	))
	r := text.NewReader(src)

	pctx := parser.NewContext()
	doc := goldmark.New().Parser().
		Parse(r, parser.WithContext(pctx)).(*ast.Document)

	var trans Transformer
	for i := 0; i < 10; i++ {
		trans.Transform(doc, r, pctx)
	}
}
