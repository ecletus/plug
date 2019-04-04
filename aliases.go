package plug

import (
	"github.com/moisespsena-go/pluggable"
)

type (
	CallbackFunc             = pluggable.CallbackFunc
	CallbackFuncE            = pluggable.CallbackFuncE
	EventDispatcherInterface = pluggable.EventDispatcherInterface
	EventDispatcher          = pluggable.EventDispatcher
	PluginEventDispatcher    = pluggable.PluginEventDispatcher
	EventInterface           = pluggable.EventInterface
	Event                    = pluggable.Event

	Options = pluggable.Options

	PluginEventDispatcherInterface = pluggable.PluginEventDispatcherInterface
	PluginEventInterface           = pluggable.PluginEventInterface
	PluginEvent                    = pluggable.PluginEvent
	PluginsFS                      = pluggable.PluginsFS

	FSEvent       = pluggable.FSEvent
	LocaleFSEvent = pluggable.LocaleFSEvent

	PluginAccess = pluggable.PluginAccess
	Accessible   = pluggable.Accessible

	Plugin = pluggable.Plugin
)

var (
	NewEvent       = pluggable.NewEvent
	NewPluginEvent = pluggable.NewPluginEvent
	EAll           = pluggable.EAll
	UID            = pluggable.UID
	UIDs           = pluggable.UIDs
	OnInit         = pluggable.OnInit
	OnPostInit     = pluggable.OnPostInit

	OnFS       = pluggable.OnFS
	OnLocaleFS = pluggable.OnLocaleFS
)
