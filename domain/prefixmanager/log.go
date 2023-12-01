package prefixmanager

import (
	"github.com/Pyrinpyi/pyipad/infrastructure/logger"
	"github.com/Pyrinpyi/pyipad/util/panics"
)

var log = logger.RegisterSubSystem("PRFX")
var spawn = panics.GoroutineWrapperFunc(log)
