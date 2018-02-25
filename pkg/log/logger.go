package log

import (
	"log"
	"fmt"
)

func init() {
	log.SetFlags(log.Ltime | log.Ldate | log.Lshortfile)
}

func Infoln(v ...interface{}) {
	log.Println("[INFO]",fmt.Sprintln(v...))
}

func Errorln(v ...interface{}) {
	log.Println("[ERROR]",fmt.Sprintln(v...))
}