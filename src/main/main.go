package main

import (
	"ascii"
	"flag"
	"log"
	"logger"
)

type Cmd struct {
	imgPath string
	help bool
	threshold int
}

func main(){
	logger.InitLog()
	cmd := parseCmd()
	if cmd.help{
		flag.Usage()
		return
	}
	if cmd.imgPath != ""{
		log.Printf("load file [%s]" ,cmd.imgPath)
		if err := ascii.TransFile(cmd.imgPath ,cmd.threshold);err != nil{
			//log.Panicf("failed to converted the file into ascii %v" ,err)
			log.Printf("failed to converted the file into ascii %v\n" ,err)
		}
	}else{
		log.Println("no file is specified")
		return
	}
}

func parseCmd() *Cmd{
	cmd := &Cmd{}
	flag.StringVar(&cmd.imgPath ,"f" ,"" ,"specify an image file to transcode")
	flag.BoolVar(&cmd.help ,"h" ,false ,"help")
	flag.IntVar(&cmd.threshold ,"t" ,140 ,"set threshold for image grey processing")
	flag.Parse()
	flag.Usage = func() {
		flag.PrintDefaults()
	}
	return cmd
}
