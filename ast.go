package pikchr

import "github.com/yuin/goldmark/ast"

// Kind is the kind of a pikchr block.
var Kind = ast.NewNodeKind("PikchrBlock")

// Block is a pikchr block.
//
//	```pikchr
//	arrow right 200% "Markdown" "Source"
//	box rad 10px "Markdown" "Formatter" "(markdown.c)" fit
//	arrow right 200% "HTML+SVG" "Output"
//	arrow <-> down 70% from last box.s
//	box same "Pikchr" "Formatter" "(pikchr.c)" fi	t
//	```
//
// Its raw contents are the plain text of the pikchr diagram.
type Block struct {
	ast.BaseBlock
}

// IsRaw reports that this block should be rendered as-is.
func (*Block) IsRaw() bool { return true }

// Kind reports that this is a pikchrBlock.
func (*Block) Kind() ast.NodeKind { return Kind }

// Dump dumps the contents of this block to stdout.
func (b *Block) Dump(src []byte, level int) {
	ast.DumpHelper(b, src, level, nil, nil)
}
