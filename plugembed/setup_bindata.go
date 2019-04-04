// +build assetfs_bindata

package plugembed

import assetfs "github.com/moisespsena-go/assetfs"

func (e *Embed) Setup(dest string) {
	e.DestPath = func() string {
		return dest
	}
	if err := assetfs.SaveExecutable(e.Name, e.FS, dest, false); err != nil {
		panic(err)
	}
}
