package main

import logger "github.com/thanover/gologger"

var AppLogger logger.Logger

func main() {
	logger.Initialize("./log")
	AppLogger = logger.GetLogger()
	AppLogger.Info("Test Info")
	AppLogger.Warn("Test Warning")
	AppLogger.Error("Test Error")
	AppLogger.Success("Test Success")
}
