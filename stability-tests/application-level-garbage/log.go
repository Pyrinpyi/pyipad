package main

import (
	"github.com/Pyrinpyi/pyipad/infrastructure/logger"
	"github.com/Pyrinpyi/pyipad/util/panics"
)

var (
	backendLog = logger.NewBackend()
	log        = backendLog.Logger("APLG")
	spawn      = panics.GoroutineWrapperFunc(log)
)
