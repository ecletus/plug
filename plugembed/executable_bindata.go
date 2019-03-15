// +build assetfs_bindata

package plugembed

import (
	"os"
	"path/filepath"
)

func (e *Executable) BinPath() string {
	var cwd, _ = os.Getwd()
	return filepath.Join(cwd, ".embed", "bin", e.Name)
}