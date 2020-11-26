package logging

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

type Level int

var (
	F                  *os.File
	DefaultPrefix      = ""
	DefaultCallerDepth = 2
	logger             *log.Logger
	logPrefix          = ""
	levelFlags         = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
)

const (
	colorRed     = uint8(iota + 91) // 红
	colorGreen                      // 绿
	colorYellow                     // 黄
	colorBlue                       // 蓝
	colorMagenta                    // 洋红
)

const (
	DEBUG Level = iota
	INFO
	WARNING
	ERROR
	FATAL
)

func init() {
	filePath := getLogFileFullPath()
	F = openLogFile(filePath)
	mw := io.MultiWriter(os.Stdout, F)
	logger = log.New(mw, DefaultPrefix, log.LstdFlags)
}
func Debug(v ...interface{}) {
	setPrefix(DEBUG, green)
	logger.Println(v)
}
func Infoo(v ...interface{}) {
	setPrefix(INFO, blue)
	logger.Println(v)
}
func Warn(v ...interface{}) {
	setPrefix(WARNING, yellow)
	logger.Println(v)
}
func Error(v ...interface{}) {
	setPrefix(ERROR, magenta)
	logger.Println(v)
}
func Fatal(v ...interface{}) {
	setPrefix(FATAL, red)
	logger.Fatalln(v)
}
func setPrefix(level Level, colorFunc func(string) string) {
	_, file, line, ok := runtime.Caller(DefaultCallerDepth)
	if ok {
		logPrefix = fmt.Sprintf("[%s][%s:%d]", levelFlags[level], filepath.Base(file), line)
	} else {
		logPrefix = fmt.Sprintf("[%s]", levelFlags[level])
	}
	logger.SetPrefix(colorFunc(logPrefix))
}

func red(s string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", colorRed, s)
}

func green(s string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", colorGreen, s)
}

func yellow(s string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", colorYellow, s)
}

func blue(s string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", colorBlue, s)
}

func magenta(s string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", colorMagenta, s)
}
