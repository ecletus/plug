package plugprivate

import (
	"github.com/ecletus/plug"
	"github.com/moisespsena-go/assetfs/assetfsapi"
	"github.com/moisespsena-go/pluggable"
)

type PluginWithPrivateFS struct {
	plug.EventDispatcher
	privateFS assetfsapi.Interface
	plug.Accessible
}

func (p *PluginWithPrivateFS) OnRegister() {
	plug.OnFS(p, func(e *pluggable.FSEvent) {
		p.privateFS = e.PrivateFS
		p.onFS()
	})
}

func (p *PluginWithPrivateFS) PrivateFS() assetfsapi.Interface {
	return p.privateFS
}
