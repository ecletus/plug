// +build !assetfs_bindata

package plugprivate

import (
	"github.com/ecletus/plug"
)

func (p *PluginWithPrivateFS) onFS() {
	p.Trigger(&PrivatePathRegisterEvent{
		plug.NewEvent(E_PRIVATE_PATH_REGISTER),
		p.Plugin(),
	})
}

func (p *PluginWithPrivateFS) OnPrivatePathRegister(cb func(e *PrivatePathRegisterEvent)) {
	p.On(E_PRIVATE_PATH_REGISTER, func(e plug.EventInterface) {
		cb(e.(*PrivatePathRegisterEvent))
	})
}
