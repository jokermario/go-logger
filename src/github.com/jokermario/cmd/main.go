package main

import (
	"time"

	"github.com/jokermario/logging"
)

func main()  {
	logger := logging.New(time.RFC3339, true)

	logger.Log(logging.Info, "The INFO log statement")
	logger.Log(logging.Warn, "The WARN log statement")
	logger.Log(logging.Error, "The ERROR log statement")
}