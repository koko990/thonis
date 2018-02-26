package log

import (
	"log"
	"fmt"
	"runtime"
)

func init() {
	log.SetFlags(log.Ltime | log.Ldate )
}

func Infoln(v ...interface{}) {
	_,f,l,_:=runtime.Caller(1)
	fmt.Println("[INFO]",f,l,fmt.Sprint(v...))
}

func Errorln(v ...interface{}) {
	log.Println("[ERROR]",fmt.Sprintln(v...))
}