package connmanager

import (
	"github.com/Pyrinpyi/pyipad/infrastructure/logger"
	"github.com/Pyrinpyi/pyipad/util/panics"
)

var log = logger.RegisterSubSystem("CMGR")
var spawn = panics.GoroutineWrapperFunc(log)
