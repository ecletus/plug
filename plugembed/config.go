package plugembed

import (
	"os"
	"os/exec"

	"github.com/moisespsena-go/httpu"
)

type IOConfig struct {
	Stdin  string
	Stdout string
	Stderr string

	Addr httpu.Addr
}

func (c *IOConfig) Sys() {
	c.Stdout = "-"
	c.Stderr = "-"
}

func (c *IOConfig) Prepare(cmd *exec.Cmd) error {
	if c.Stdout == "-" {
		cmd.Stdout = os.Stdout
	}
	if c.Stderr == "-" {
		cmd.Stderr = os.Stderr
	}
	return nil
}
