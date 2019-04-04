package plugembed

import (
	"github.com/op/go-logging"

	"github.com/moisespsena-go/assetfs/assetfsapi"
)

type EmbedInfo struct {
	Path string
}

type Embed struct {
	Name       string
	FS         assetfsapi.Interface
	Log        *logging.Logger
	BindaClean bool

	SrcPath  func() string
	DestPath func() string
	Build    func(dest string) error
	Clean    func() error

	Enabled func() bool
}

func (e *Embed) GetAsset() assetfsapi.AssetInterface {
	a, _ := e.FS.Asset(e.Name)
	return a
}

type EmbedSlice []*Embed
