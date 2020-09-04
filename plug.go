package plug

import (
	"github.com/ecletus/assets"
	"github.com/moisespsena-go/assetfs/assetfsapi"
	"github.com/moisespsena-go/pluggable"
)

type Plugins struct {
	*pluggable.I18nPlugins
}

func New(fs assetfsapi.Interface) *Plugins {
	p := &Plugins{pluggable.NewI18nPlugins(fs, assets.NS_LOCALE)}
	p.SetDispatcher(p)
	return p
}
