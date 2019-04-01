// +build !assetfs_bindata,!assetfs_bindataClean,!assetfs_bindataCompile

package plugembed

import (
	"github.com/moisespsena-go/error-wrap"
	"github.com/moisespsena-go/path-helpers"
)

func (e *Embed) Setup(dest string) {
	pth := e.SrcPath()
	if !path_helpers.IsExistingRegularFile(pth) {
		err := e.Build(dest)
		if err != nil {
			panic(errwrap.Wrap(err, "Build Embed %q", e.Name))
		}
	}
}
