package mylog

//自定义日志

import (
	"errors"
	"fmt"
	"path"
	"runtime"
	"strings"
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

func getLogInfo(skip int) (funcName, fileName string, line int) {
	// runtime.Caller 获得运行文件的详细信息 行号 文件 参数n表示，向上找几层，0表示获得当前runtime.Caller执行层的行号 文件名 等信息
	pc, file, line, ok := runtime.Caller(skip)
	if !ok {
		fmt.Println("runtime Caller failed, err")
		return
	}
	funcName = runtime.FuncForPC(pc).Name()
	fmt.Println(pc, funcName, file, line)
	fileName = path.Base(file)
	return
}
