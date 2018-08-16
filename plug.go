package plug

import (
	"github.com/moisespsena/go-assetfs/api"
	"github.com/moisespsena/go-pluggable"
)

type Plugins struct {
	*pluggable.I18nPlugins
}

func New(fs api.Interface) *Plugins {
	p := &Plugins{pluggable.NewI18nPlugins(fs)}
	p.SetDispatcher(p)
	return p
}
