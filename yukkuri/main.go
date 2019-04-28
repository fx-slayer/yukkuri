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
		if gray ,err := ykr.TransImgToGrey();err != nil{
			log.Printf("failed to converted the file into ascii %v\n" ,err)
		}else{
			converter := ascii.NewAscii(gray)
			ykr.TransImgToAsc(converter)
		}
	}else{
		log.Println("no file is specified")
		return
	}
}