package logger

import (
	"testing"
)

func TestInfo(t *testing.T) {
	//arrange
	logger := MakeLogger("./log")
	msg := "Test Message"
	//act
	logger.Info(msg)

	//assert

}
