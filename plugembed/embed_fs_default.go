// +build !assetfs_bindata

package plugembed

import (
	"os"
	"path/filepath"

	"github.com/ecletus/plug"
	"github.com/moisespsena-go/build-dest"
	"github.com/moisespsena-go/assetfs/assetfsapi"
	"github.com/moisespsena-go/path-helpers"
)

func (p *PluginWithEmbededCommands) _setup(ev *plug.FSEvent) {
	var (
		wd, _     = os.Getwd()
		binDir    = bdest.BuildDir(filepath.Join(wd, ".embed", "bin", ev.Plugin().NameSpace))
		staticDir = filepath.Join(wd, ".embed", "static", ev.Plugin().NameSpace)
	)

	if len(p.executables) > 0 {
		if err := path_helpers.MkdirAllIfNotExists(binDir); err != nil {
			panic(err)
		}
		p.EmbedBinFS().(assetfsapi.PathRegistrator).RegisterPath(binDir)
		p.executablesSetup(binDir)
	}

	if len(p.statics) > 0 {
		if err := path_helpers.MkdirAllIfNotExists(staticDir); err != nil {
			panic(err)
		}
		p.EmbedStaticFS().(assetfsapi.PathRegistrator).RegisterPath(staticDir)
		p.staticsSetup(staticDir)
	}
}
