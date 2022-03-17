[![Go Reference](https://pkg.go.dev/badge/github.com/jchenry/goldmark-pikchr.svg)](https://pkg.go.dev/github.com/jchenry/goldmark-pikchr)
[![Go](https://github.com/jchenry/goldmark-pikchr/actions/workflows/go.yml/badge.svg)](https://github.com/jchenry/goldmark-pikchr/actions/workflows/go.yml)

goldmark-pikchr is an extension for the [goldmark] Markdown parser that adds
support for [pikchr] diagrams.

  [goldmark]: http://github.com/yuin/goldmark
  [pikchr]: https://pikchr.org/home/doc/trunk/homepage.md

# Usage

To use goldmark-pikchr, import the `pikchr` package.

```go
import pikchr "github.com/jchenry/goldmark-pikchr"
```

Then include the `pikchr.Extender` in the list of extensions you build your
[`goldmark.Markdown`] with.

  [`goldmark.Markdown`]: https://pkg.go.dev/github.com/yuin/goldmark#Markdown

```go
goldmark.New(
  &pikchr.Extender{}
  // ...
).Convert(src, out)
```

The package supports pikchr diagrams inside fenced code blocks with the language `pikchr`. For example,

    ```pikchr
    arrow right 200% "Markdown" "Source"
    box rad 10px "Markdown" "Formatter" "(markdown.c)" fit
    arrow right 200% "HTML+SVG" "Output"
    arrow <-> down 70% from last box.s
    box same "Pikchr" "Formatter" "(pikchr.c)" fit
    ```

When you render the Markdown as HTML, these will be replaced with SVG blocks.
[pikchr] will render these into diagrams client-side.

Code is based on [github.com/abhinav/goldmark-mermaid](https://github.com/abhinav/goldmark-mermaid) and licensed under the same terms. 
