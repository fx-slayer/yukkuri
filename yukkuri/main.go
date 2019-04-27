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
		ykr := ascii.NewYukkuri(c)
		if err := ykr.TransImgToGrey();err != nil{
			log.Printf("failed to converted the file into ascii %v\n" ,err)
		}
		ykr.TransImgToAsc(ascii.NewAscii(ykr.GrayImg))
	}else{
		log.Println("no file is specified")
		return
	}
}