package print

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/brahama/cf-doc/doc"
)

//Pretty prints pretty json of Doc
func Pretty(d *doc.Doc) (string, error) {
	j, err := json.MarshalIndent(d, "", "    ")

	return string(j), err
}

// Markdown prints the doc as Markdown
func Markdown(d *doc.Doc) (string, error) {
	var buf bytes.Buffer

	if len(d.Usage) > 0 {
		buf.WriteString(fmt.Sprintf("# %s\n", d.Usage))
	}

	if len(d.Parameters) > 0 {
		buf.WriteString("\n## Parameters\n\n")
		buf.WriteString("| Name | Description | Type | Default | Allowed Values |\n")
		buf.WriteString("|------|-------------|:----:|:-------:|:---------------|\n")

		for k := range d.Parameters {
			buf.WriteString(fmt.Sprintf("| %s | %s | %s | %s | %s |\n",
				d.Parameters[k].Name,
				d.Parameters[k].Description,
				d.Parameters[k].Type,
				d.Parameters[k].Default,
				d.Parameters[k].AllowedValues))
		}

	}

	if len(d.Outputs) > 0 {
		buf.WriteString("\n## Outputs\n\n")
		buf.WriteString("| Name | Description | Export |\n")
		buf.WriteString("|------|-------------|--------|\n")

		for k := range d.Outputs {
			buf.WriteString(fmt.Sprintf("| %s | %s | %s |\n",
				d.Outputs[k].Name,
				d.Outputs[k].Description,
				d.Outputs[k].Export))
		}
	}

	return buf.String(), nil

}
