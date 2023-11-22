package handshake

import (
	"github.com/Pyrinpyi/pyipad/infrastructure/logger"
	"github.com/Pyrinpyi/pyipad/util/panics"
)

var log = logger.RegisterSubSystem("PROT")
var spawn = panics.GoroutineWrapperFunc(log)
