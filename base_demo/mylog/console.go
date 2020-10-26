package mylog

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

type LogLevel int16

const (
	UNKOWN LogLevel = iota
	DEBUG
	TRACE
	INFO
	WARING
	ERROR
	FATAL
)

// Logger log
type Logger struct {
	Leven LogLevel
}

func parse_log_level(str string) (LogLevel, error) {
	s := strings.ToLower(str)
	switch s {
	case "debug":
		return DEBUG, nil
	case "trace":
		return TRACE, nil
	case "info":
		return INFO, nil
	case "waring":
		return WARING, nil
	case "error":
		return ERROR, nil
	case "fatal":
		return FATAL, nil
	default:
		err := errors.New("日志级别未知")
		return UNKOWN, err
	}
}

//NewLog
func NewLog(level string) Logger {
	p_level, err := parse_log_level(level)
	if err != nil {
		panic(err)
	}
	return Logger{
		Leven: p_level,
	}
}

func (l Logger) enable(logl LogLevel) bool {
	return logl >= l.Leven
}

func (l Logger) Debug(str string) {
	if l.enable(DEBUG) {
		now := time.Now()
		fmt.Printf("[%s][Debug] %s \n", now.Format("2006-01-02 15:04:05"), str)
	}

}

func (l Logger) Info(str string) {
	if l.enable(INFO) {
		now := time.Now()
		fmt.Printf("[%s][Info] %s \n", now.Format("2006-01-02 15:04:05"), str)
	}
}

func (l Logger) Warning(str string) {
	if l.enable(WARING) {
		now := time.Now()
		fmt.Printf("[%s][Warning] %s \n", now.Format("2006-01-02 15:04:05"), str)
	}
}

func (l Logger) Fatal(str string) {
	if l.enable(FATAL) {
		now := time.Now()
		fmt.Printf("[%s][Fatal] %s \n", now.Format("2006-01-02 15:04:05"), str)
	}
}
func (l Logger) ERROR(str string) {
	if l.enable(ERROR) {
		now := time.Now()
		fmt.Printf("[%s][ERROR] %s \n", now.Format("2006-01-02 15:04:05"), str)
	}
}
