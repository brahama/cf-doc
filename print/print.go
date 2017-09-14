package print

import "github.com/brahama/cf-doc/doc"
import "encoding/json"

//Pretty prints pretty json of Doc
func Pretty(d *doc.Doc) (string, error) {
	j, err := json.MarshalIndent(d, "", "    ")

	return string(j), err
}
