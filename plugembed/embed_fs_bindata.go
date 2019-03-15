// +build assetfs_bindata

package plugembed

import (
	"path/filepath"

	"github.com/aghape/plug"
)

func (p *PluginWithEmbededCommands) _setup(ev *plug.FSEvent) {
	for _, e := range p.executables {
		e.Setup(filepath.Join(".embed", "bin", e.Name))
	}
	for _, e := range p.statics {
		e.Setup(filepath.Join(".embed", "static", e.Name))
	}
}
