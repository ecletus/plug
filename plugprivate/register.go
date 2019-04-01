package plugprivate

import (
	"github.com/ecletus/plug"
	"github.com/moisespsena-go/path-helpers"
)

var E_PRIVATE_PATH_REGISTER = path_helpers.GetCalledDir() + ":privatePathRegister"

type PrivatePathRegisterEvent struct {
	plug.EventInterface
	Plugin *plug.Plugin
}
