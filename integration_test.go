package pikchr_test

import (
	"bytes"
	"flag"
	"os"
	"strings"
	"testing"

	pikchr "github.com/jchenry/goldmark-pikchr"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/yuin/goldmark"
	"gopkg.in/yaml.v3"
)

const _fixtures = "testdata/tests.yaml"

// Run 'go test -update-fixtures' to update "want"s in tests.yaml.
var _updateFixtures = flag.Bool("update-fixtures", false,
	"if true, update integration test fixtures")

func TestIntegration(t *testing.T) {
	t.Parallel()

	testdata, err := os.ReadFile(_fixtures)
	require.NoError(t, err)

	var tests []struct {
		Desc string `yaml:"desc"`
		Give string `yaml:"give"`
		Want string `yaml:"want"`
	}
	require.NoError(t, yaml.Unmarshal(testdata, &tests))

	for idx, tt := range tests {
		t.Run(tt.Desc, func(t *testing.T) {
			var ext pikchr.Extender
			md := goldmark.New(goldmark.WithExtensions(&ext))

			var got bytes.Buffer
			require.NoError(t, md.Convert([]byte(tt.Give), &got))

			if *_updateFixtures {
				// If run with --update-fixtures,
				// the test will always pass.
				tt.Want = ensureNewline(got.String())
				tests[idx] = tt
			} else {
				assert.Equal(t,
					ensureNewline(tt.Want),
					ensureNewline(got.String()),
				)
			}
		})
	}

	if *_updateFixtures {
		testdata, err := yaml.Marshal(tests)
		require.NoError(t, err)
		require.NoError(t, os.WriteFile(_fixtures, testdata, 0o644))
	}
}

// Ensures that the given string has a trailing newline.
func ensureNewline(s string) string {
	if strings.HasSuffix(s, "\n") {
		return s
	}
	return s + "\n"
}
