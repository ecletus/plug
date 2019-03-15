// +build assetfs_bindata

package plugprivate

func (p *PluginWithPrivateFS) onFS() {
}

func (p *PluginWithPrivateFS) OnPrivatePathRegister(cb func(e *PrivatePathRegisterEvent)) {
}
