package mylog

import (
	"fmt"
	"time"
)

// Logger log
type Logger struct {
	Leven LogLevel
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

func printLog(lv, msg string, arg ...interface{}) {
	fmt.Println(arg)
	funcName, fileName, lineNo := getLogInfo(3)
	now := time.Now()
	fmt.Printf("[%s][%s][funcName: %s,fileName: %s, line: %d] %s \n",
		now.Format("2006-01-02 15:04:05"), lv, funcName, fileName, lineNo, msg)
}

func (l Logger) Debug(str string, arg ...interface{}) {
	if l.enable(DEBUG) {
		printLog("DEBUG", str, arg)
	}

}

func (l Logger) Info(str string, arg ...interface{}) {
	if l.enable(INFO) {
		printLog("INFO", str, arg...)
	}
}

func (l Logger) Warning(str string) {
	if l.enable(WARING) {
		printLog("WARING", str)
	}
}

func (l Logger) Fatal(str string) {
	if l.enable(FATAL) {
		printLog("FATAL", str)
	}
}
func (l Logger) ERROR(str string) {
	if l.enable(ERROR) {
		printLog("ERROR", str)
	}
}
