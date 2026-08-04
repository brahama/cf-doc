/*
* Splices generated docs into an existing file (e.g. a Readme) between a
* pair of marker comments, leaving the rest of the file untouched. Modeled
* after terraform-docs' inject output mode.
*
 */

package inject

import (
	"fmt"
	"strings"
)

// StartMarker and EndMarker delimit the region of a file that gets
// replaced with generated content. Both must already be present in the
// target file.
const (
	StartMarker = "<!-- cf-doc:start -->"
	EndMarker   = "<!-- cf-doc:end -->"
)

// Content returns existing with the region between StartMarker and
// EndMarker replaced by content. The markers themselves are preserved.
func Content(existing, content string) (string, error) {
	start := strings.Index(existing, StartMarker)
	if start == -1 {
		return "", fmt.Errorf("could not find %q marker", StartMarker)
	}

	end := strings.Index(existing, EndMarker)
	if end == -1 {
		return "", fmt.Errorf("could not find %q marker", EndMarker)
	}

	if end < start {
		return "", fmt.Errorf("%q marker found before %q marker", EndMarker, StartMarker)
	}

	before := existing[:start+len(StartMarker)]
	after := existing[end:]

	return fmt.Sprintf("%s\n%s\n%s", before, strings.TrimSpace(content), after), nil
}
