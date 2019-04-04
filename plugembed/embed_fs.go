package plugembed

import (
	"github.com/ecletus/plug"
	"github.com/moisespsena-go/assetfs/assetfsapi"
)

func (p *PluginWithEmbededCommands) EmbedFS() assetfsapi.Interface {
	return p.PrivateFS().NameSpace("@embed")
}

func (p *PluginWithEmbededCommands) EmbedBinFS() assetfsapi.Interface {
	return p.EmbedFS().NameSpace("bin")
}

func (p *PluginWithEmbededCommands) EmbedStaticFS() assetfsapi.Interface {
	return p.EmbedFS().NameSpace("static")
}

func (p *PluginWithEmbededCommands) setup(ev *plug.FSEvent) {
	for _, e := range p.executables {
		e.FS = p.EmbedBinFS()
	}
	for _, e := range p.statics {
		e.FS = p.EmbedStaticFS()
	}

	p._setup(ev)
}
