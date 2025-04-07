package main

import logger "github.com/thanover/gologger"

func main() {
	logger := logger.MakeLogger("./log")
	logger.Info("Test Info")
	logger.Warn("Test Warning")
	logger.Error("Test Error")
	logger.Success("Test Success")
}
