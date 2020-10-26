package mylog

import "fmt"

// Logger log
type Logger struct {
}

//NewLog
func NewLog() Logger {
	return Logger{}
}

func (l Logger) Debug(str string) {
	fmt.Println(str)
}

func (l Logger) Info(str string) {
	fmt.Println(str)
}

func (l Logger) Warning(str string) {
	fmt.Println(str)
}

func (l Logger) Fatal(str string) {
	fmt.Println(str)
}
