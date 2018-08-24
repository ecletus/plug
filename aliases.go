package plug

import (
	"github.com/moisespsena/go-pluggable"
)

type (
	CallbackFunc             = pluggable.CallbackFunc
	CallbackFuncE            = pluggable.CallbackFuncE
	EventDispatcherInterface = pluggable.EventDispatcherInterface
	EventDispatcher          = pluggable.EventDispatcher
	EventInterface           = pluggable.EventInterface
	Event                    = pluggable.Event

	Options = pluggable.Options

	PluginEventDispatcherInterface = pluggable.PluginEventDispatcherInterface
	PluginEventInterface           = pluggable.PluginEventInterface
	PluginEvent                    = pluggable.PluginEvent
	PluginsFS                      = pluggable.PluginsFS

	AssetFSEvent  = pluggable.AssetFSEvent
	LocaleFSEvent = pluggable.LocaleFSEvent
)

var (
	NewEvent       = pluggable.NewEvent
	NewPluginEvent = pluggable.NewPluginEvent
	EAll           = pluggable.EAll
	UID            = pluggable.UID
	UIDs           = pluggable.UIDs
	EInit          = pluggable.EInit
	EPostInit      = pluggable.EPostInit

	OnAssetFS  = pluggable.OnAssetFS
	OnLocaleFS = pluggable.OnLocaleFS
)
