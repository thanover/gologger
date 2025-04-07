package main

import "github.com/thanover/go-logger"

func main() {
	logger := logger.MakeLogger("./log")
	logger.Info("Test Info")
	logger.Warn("Test Warning")
	logger.Error("Test Error")
	logger.Success("Test Success")
}
