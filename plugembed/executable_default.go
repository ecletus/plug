// +build !assetfs_bindata

package plugembed

func (e *Executable) BinPath() string {
	return e.SrcPath()
}
