package logger

import "log"

func InitLog(){
	log.SetFlags(0)
	log.SetPrefix("yukkuri > ")
}
