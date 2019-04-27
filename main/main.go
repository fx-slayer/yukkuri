package main

import (
	"ascii"
	"command"
	"flag"
	"log"
	"logger"
)

func main(){
	logger.InitLog()
	c := command.ParseCmd()
	if c.Help{
		flag.Usage()
		return
	}
	if c.ImgPath != ""{
		log.Printf("load file [%s]" ,c.ImgPath)
		if err := ascii.TransFile(c.ImgPath ,c);err != nil{
			log.Printf("failed to converted the file into ascii %v\n" ,err)
		}
	}else{
		log.Println("no file is specified")
		return
	}
}