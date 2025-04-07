package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type logger struct {
	Dir  string
	File *os.File
}

type msg struct {
	Raw        string
	Level      string
	ConsoleMsg string
	FileMsg    string
}

const (
	green  = "\033[97;42m"
	yellow = "\033[90;43m"
	red    = "\033[97;41m"
	blue   = "\033[97;44m"
	// magenta = "\033[97;45m"
	// cyan    = "\033[97;46m"
	reset = "\033[0m"
)

func MakeLogger(dir string) logger {
	makeDir(dir)
	file := makeFile(dir)
	return logger{Dir: dir, File: file}
}

func makeDir(dir string) {
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		fmt.Println("Error creating directory:", err)
		return
	}
}

func makeFile(dir string) *os.File {
	filePath := filepath.Join(dir, "log.txt")
	file, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	return file
}

func (l logger) Info(input string) {
	l.log(input, "INFO", blue)
}

func (l logger) Warn(input string) {
	l.log(input, "WARN", yellow)
}

func (l logger) Error(input string) {
	l.log(input, "ERROR", red)
}

func (l logger) Success(input string) {
	l.log(input, "SUCCESS", green)
}

func (l logger) log(input, level, color string) {
	m := msg{Raw: input, Level: level}
	m.fmtMsg(color)
	m.logToConsole()
	if l.Dir != "" {
		l.writeToFile(m)
	}
}

func (m msg) logToConsole() {
	fmt.Println(m.ConsoleMsg)
}

func (l logger) writeToFile(m msg) {
	_, err := l.File.WriteString(fmt.Sprintln(m.FileMsg))
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}

func (msg *msg) fmtMsg(color string) {
	t := time.Now().Local().Format("02 Jan 06 15:04.00 MST")

	msg.FileMsg = fmt.Sprintf("[%s][%s]: %s", t, msg.Level, msg.Raw)

	lvlWithColor := fmt.Sprintf("%s %s %s", color, msg.Level, reset)
	msg.ConsoleMsg = fmt.Sprintf("[%s][%s]: %s", t, lvlWithColor, msg.Raw)
}
