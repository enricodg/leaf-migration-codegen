package handler

import (
	"github.com/paulusrobin/leaf-utilities/leafMigration/logger"
	leafLogger "github.com/paulusrobin/leaf-utilities/logger/logger"
	"sync"
)

var (
	instance handler
	once     sync.Once
)

type (
	handler struct {
		log leafLogger.Logger
	}
)

func GetHandler() handler {
	once.Do(func() {
		instance = handler{
			log: logger.GetLogger(),
		}
	})
	return instance
}
