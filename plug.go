package plug

import (
	"github.com/moisespsena-go/assetfs/assetfsapi"
	"github.com/moisespsena-go/pluggable"
)

type Plugins struct {
	*pluggable.I18nPlugins
}

func New(fs assetfsapi.Interface) *Plugins {
	p := &Plugins{pluggable.NewI18nPlugins(fs)}
	p.SetDispatcher(p)
	return p
}
