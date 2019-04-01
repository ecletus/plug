package plugembed

import (
	"os/exec"
	"path/filepath"
	"syscall"

	"github.com/ecletus/ecletus"
	"github.com/moisespsena-go/task"
	"github.com/moisespsena-go/error-wrap"
	"github.com/spf13/cobra"

	"github.com/moisespsena-go/path-helpers"
)

type ExecutableSlice []*Executable

func (es ExecutableSlice) CreateCobraCommands() (cmds []*cobra.Command) {
	for _, emb := range es {
		if emb.Enabled == nil || emb.Enabled() {
			cmds = append(cmds, emb.CobraCommandCreate())
		}
	}
	return cmds
}

type Executable struct {
	Embed

	IO IOConfig

	Args    func(args []string) []string
	Prepare func(cmd *exec.Cmd) error

	CobraCommandFactory func() *cobra.Command
	CobraCommandCreated func(cmd *cobra.Command)
}

func (e *Executable) MakeCmd(args ...string) (cmd *exec.Cmd, err error) {
	binPath := e.BinPath()
	dir := filepath.Dir(binPath)
	if err = path_helpers.MkdirAllIfNotExists(dir); err != nil {
		return
	}
	cmd = exec.Command(binPath, args...)
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}

	if e.Prepare != nil {
		if err = e.Prepare(cmd); err != nil {
			return nil, err
		}
	}

	return cmd, nil
}

func (e *Executable) CreateTask(args ...string) (t task.Task, err error) {
	var c *exec.Cmd
	if c, err = e.MakeCmd(args...); err != nil {
		return nil, err
	}

	if err = e.IO.Prepare(c); err != nil {
		err = errwrap.Wrap(err, "prepare io")
		return
	}

	if e.Args != nil {
		args = e.Args(args)
	}

	return task.NewCmdTask(c).SetLog(e.Log), nil
}

func (e *Executable) CobraCommandSetup(cmd *cobra.Command) *cobra.Command {
	if cmd.Use == "" {
		cmd.Use = e.Name
	}

	var runE = cmd.RunE
	if runE == nil && cmd.Run != nil {
		run := cmd.Run
		runE = func(cmd *cobra.Command, args []string) error {
			run(cmd, args)
			return nil
		}
	}

	cmd.RunE = func(cmd *cobra.Command, args []string) (err error) {
		if runE != nil {
			if err = runE(cmd, args); err != nil {
				return err
			}
		}
		var t task.Task
		if t, err = e.CreateTask(args...); err != nil {
			return errwrap.Wrap(err, "Create task")
		}

		ect := ecletus.Get()
		return errwrap.Wrap(ect.AddTask(t), "Ecletus: Add Task")
	}
	return cmd
}

func (e *Executable) CobraCommandCreate() (cmd *cobra.Command) {
	if e.CobraCommandFactory == nil {
		cmd = &cobra.Command{}
	} else {
		cmd = e.CobraCommandFactory()
	}

	cmd = e.CobraCommandSetup(cmd)
	if e.CobraCommandCreated != nil {
		e.CobraCommandCreated(cmd)
	}
	cmd.SetHelpFunc(func(cmd *cobra.Command, i []string) {
		if err := cmd.RunE(cmd, []string{"-h"}); err != nil {
			e.Log.Error(err)
		}
	})
	return
}
