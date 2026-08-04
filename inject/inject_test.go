package inject

import "testing"

func TestContentReplacesRegionBetweenMarkers(t *testing.T) {
	existing := `# My Stack

Some intro text.

<!-- cf-doc:start -->
old docs
<!-- cf-doc:end -->

Footer text.
`
	expected := `# My Stack

Some intro text.

<!-- cf-doc:start -->
new docs
<!-- cf-doc:end -->

Footer text.
`

	actual, err := Content(existing, "new docs")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if actual != expected {
		t.Errorf("Test failed, got %q, expected %q", actual, expected)
	}
}

func TestContentMissingStartMarker(t *testing.T) {
	_, err := Content("no markers here\n<!-- cf-doc:end -->", "docs")
	if err == nil {
		t.Fatal("expected an error, got nil")
	}
}

func TestContentMissingEndMarker(t *testing.T) {
	_, err := Content("<!-- cf-doc:start -->\nno end marker", "docs")
	if err == nil {
		t.Fatal("expected an error, got nil")
	}
}

func TestContentMarkersOutOfOrder(t *testing.T) {
	_, err := Content("<!-- cf-doc:end -->\n<!-- cf-doc:start -->", "docs")
	if err == nil {
		t.Fatal("expected an error, got nil")
	}
}
