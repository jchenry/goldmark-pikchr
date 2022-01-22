package pikchr_test

import (
	"testing"

	pikchr "github.com/jchenry/goldmark-pikchr"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/testutil"
)

func XTestIntegration(t *testing.T) {
	t.Parallel()

	testutil.DoTestCaseFile(
		goldmark.New(goldmark.WithExtensions(&pikchr.Extender{
			// pikchrJS: "pikchr.js",
		})),
		"testdata/tests.txt",
		t,
	)
}
