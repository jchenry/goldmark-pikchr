package pikchr

import "github.com/yuin/goldmark/ast"

// Kind is the kind of a pikchr block.
var Kind = ast.NewNodeKind("PikchrBlock")

// Block is a pikchr block.
//
//  ```pikchr
//  graph TD;
//      A-->B;
//      A-->C;
//      B-->D;
//      C-->D;
//  ```
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
