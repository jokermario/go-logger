package main

import (
	"time"

	"github.com/jokermario/go_logger/go_logger"
)

func main()  {
	logger := go_logger.New(time.RFC3339, true)

	logger.Log(go_logger.Info, "The INFO log statement")
	logger.Log(go_logger.Warn, "The WARN log statement")
	logger.Log(go_logger.Error, "The ERROR log statement")
}