// +build assetfs_bindataCompile

package plugembed

import (
	"github.com/moisespsena-go/error-wrap"
)

func (e *Embed) Setup(dest string) {
	if e.Build != nil {
		err := e.Build(dest)
		if err != nil {
			panic(errwrap.Wrap(err, "Build Embed %q", e.Name))
		}
	}
}
