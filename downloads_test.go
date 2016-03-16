package ghdownloads

import (
	"fmt"
	"testing"

	"github.com/cheekybits/is"
)

func TestDownloads(t *testing.T) {
	is := is.New(t)
	downloads, err := Count("matryer/bitbar")
	is.NoErr(err)
	fmt.Println(downloads)
}
