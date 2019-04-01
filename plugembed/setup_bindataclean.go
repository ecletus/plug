// +build assetfs_bindataClean

package plugembed

import (
	"path/filepath"
	"syscall"

	"github.com/moisespsena-go/error-wrap"
	"github.com/moisespsena-go/path-helpers"
)

func (e *Embed) Setup(dest string) {
	binPath := filepath.Join(dest, e.DestPath())
	if e.BindaClean && path_helpers.IsExistingRegularFile(binPath) {
		if err := syscall.Unlink(binPath); err != nil {
			panic(errwrap.Wrap(err, "Unlink Embed %q at %q", e.Name, binPath))
		}
	}
}
