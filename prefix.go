package plug

import (
	"github.com/moisespsena/go-default-logger"
	"github.com/op/go-logging"
	"github.com/aghape/helpers"
)

var PREFIX string

var log *logging.Logger

func init() {
	PREFIX = helpers.GetCalledDir()
	log = defaultlogger.NewLogger(PREFIX)
}
