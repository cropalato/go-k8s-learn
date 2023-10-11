//
// logger.go
// Copyright (C) 2023 rmelo <Ricardo Melo <rmelo@ludia.com>>
//
// Distributed under terms of the MIT license.
//

package main

import (
	"log"
)

type aggregatedLogger struct {
  infoLogger  *log.Logger
  warnLogger  *log.Logger
  errorLogger *log.Logger
}

func (l *aggregatedLogger) info(v ...interface{}) {
  l.infoLogger.Println(v...)
}

func (l *aggregatedLogger) warn(v ...interface{}) {
  l.warnLogger.Println(v...)
}
func (l *aggregatedLogger) error(v ...interface{}) {
  l.errorLogger.Println(v...)
}
