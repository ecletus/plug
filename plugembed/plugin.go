package plugembed

import (
	"path/filepath"

	"github.com/ecletus/cli"
	"github.com/ecletus/plug"
	"github.com/ecletus/plug/plugprivate"
	"github.com/moisespsena-go/pluggable"
)

type PluginWithEmbededCommands struct {
	plugprivate.PluginWithPrivateFS
	executables ExecutableSlice
	statics     EmbedSlice
	postSetup   []func()
}

func (p *PluginWithEmbededCommands) EmbedPostSetup(f ...func()) {
	p.postSetup = append(p.postSetup, f...)
}

func (p *PluginWithEmbededCommands) EmbedRegisterExecutable(e ...*Executable) {
	p.executables = append(p.executables, e...)
}

func (p *PluginWithEmbededCommands) EmbedRegisterStatic(emb ...*Embed) {
	p.statics = append(p.statics, emb...)
}

func (p *PluginWithEmbededCommands) EmbedExecutables() ExecutableSlice {
	return p.executables
}

func (p *PluginWithEmbededCommands) EmbedStatics() EmbedSlice {
	return p.statics
}

func (p *PluginWithEmbededCommands) executablesSetup(pth string) {
	for _, e := range p.executables {
		e.Setup(filepath.Join(pth, e.Name))
	}
}

func (p *PluginWithEmbededCommands) staticsSetup(pth string) {
	for _, e := range p.statics {
		e.Setup(filepath.Join(pth, e.Name))
	}
}

func (p *PluginWithEmbededCommands) OnRegister() {
	p.PluginWithPrivateFS.OnRegister()
	plug.OnFS(p, func(e *pluggable.FSEvent) {
		p.setup(e)
		defer func() {
			p.postSetup = nil
		}()
		for _, f := range p.postSetup {
			f()
		}
	})
	cli.OnRegister(p, func(e *cli.RegisterEvent) {
		e.RootCmd.AddCommand(p.executables.CreateCobraCommands()...)
	})
}
