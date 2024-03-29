package pikchr

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/yuin/goldmark/text"
)

func hijackStdout(t testing.TB) (path string, close func() error) {
	stdout := os.Stdout
	t.Cleanup(func() {
		os.Stdout = stdout
	})

	path = filepath.Join(t.TempDir(), "stdout")
	f, err := os.Create(path)
	require.NoError(t, err)
	os.Stdout = f
	return path, f.Close
}

func TestBlock(t *testing.T) {
	src := []byte("foo\n")

	lines := text.NewSegments()
	lines.Append(text.NewSegment(0, len(src)))

	var b Block
	b.SetLines(lines)

	t.Run("Raw", func(t *testing.T) {
		assert.True(t, b.IsRaw())
	})

	t.Run("Dump", func(t *testing.T) {
		stdout, closeStdout := hijackStdout(t)

		b.Dump(src, 0)
		require.NoError(t, closeStdout())

		got, err := os.ReadFile(stdout)
		require.NoError(t, err)
		require.Equal(t, unlines(
			"PikchrBlock {",
			"    RawText: \"foo\n\"",
			"    HasBlankPreviousLines: false",
			"}",
		), string(got))
	})
}
