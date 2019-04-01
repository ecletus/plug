package plug

import (
	"github.com/moisespsena-go/default-logger"
	"github.com/moisespsena-go/path-helpers"
)

var (
	PREFIX = path_helpers.GetCalledDir()
	log    = defaultlogger.NewLogger(PREFIX)
)
