package main

import (
	"os"

	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/formatters/html"
	"github.com/alecthomas/chroma/styles"
	"github.com/spf13/cobra"
)

var (
	_ cmder = (*genChromaStyles)(nil)
)

type baseCmd struct {
	cmd *cobra.Command
}

type genChromaStyles struct {
	style          string
	highlightStyle string
	linesStyle     string
	*baseCmd
}

// TODO(bep) highlight
func main() *genChromaStyles { //createGenChromaStyles
	g := &genChromaStyles{
		baseCmd: newBaseCmd(&cobra.Command{
			Use:   "chromastyles",
			Short: "Generate CSS stylesheet for the Chroma code highlighter",
			Long: `Generate CSS stylesheet for the Chroma code highlighter for a given style. This stylesheet is needed if pygmentsUseClasses is enabled in config.
See https://help.farbox.com/pygments.html for preview of available styles`,
		}),
	}

	g.cmd.RunE = func(cmd *cobra.Command, args []string) error {
		return g.generate()
	}

	g.cmd.PersistentFlags().StringVar(&g.style, "style", "friendly", "highlighter style (see https://help.farbox.com/pygments.html)")
	g.cmd.PersistentFlags().StringVar(&g.highlightStyle, "highlightStyle", "bg:#ffffcc", "style used for highlighting lines (see https://github.com/alecthomas/chroma)")
	g.cmd.PersistentFlags().StringVar(&g.linesStyle, "linesStyle", "", "style used for line numbers (see https://github.com/alecthomas/chroma)")

	return g
}

func (g *genChromaStyles) generate() error {
	builder := styles.Get(g.style).Builder()
	if g.highlightStyle != "" {
		builder.Add(chroma.LineHighlight, g.highlightStyle)
	}
	if g.linesStyle != "" {
		builder.Add(chroma.LineNumbers, g.linesStyle)
	}
	style, err := builder.Build()
	if err != nil {
		return err
	}
	formatter := html.New(html.WithClasses(true))
	formatter.WriteCSS(os.Stdout, style)
	return nil
}
