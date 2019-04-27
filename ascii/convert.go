package ascii

import (
	"command"
	"errors"
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"strings"
)

type Yukurri struct {
	Img 	image.Image
	File	*os.File
	Cmd 	command.Cmd
}

func TransFile(path string ,cmd command.Cmd) error{
	ykr := &Yukurri{
		Cmd:cmd,
	}
	if f ,err := os.OpenFile(path ,os.O_RDONLY ,os.ModeTemporary);err != nil{
		return err
	}else{
		ykr.File = f
		tmp := strings.Split(f.Name() ,".")
		suffix := strings.ToLower(tmp[len(tmp) - 1])
		var decodeErr error
		switch suffix {
		case "jpg":
			ykr.Img ,decodeErr = jpeg.Decode(f)
		case "jpeg":
			ykr.Img ,decodeErr = jpeg.Decode(f)
		case "png":
			ykr.Img ,decodeErr = png.Decode(f)
		default:
			return errors.New("unsupported image type")
		}
		if decodeErr != nil{
			return decodeErr
		}
		handleImage(ykr)
		return nil
	}
}

func handleImage(ykr *Yukurri){
	rtg := ykr.Img.Bounds()
	log.Printf("image size %d*%d" ,ykr.Img.Bounds().Max.X ,ykr.Img.Bounds().Max.Y)
	if ykr.Cmd.AscWidth > 0 && ykr.Cmd.AscHeight > 0{
		ykr.Img = resize.Resize(uint(ykr.Cmd.AscWidth) ,uint(ykr.Cmd.AscHeight) ,ykr.Img ,resize.Lanczos3)
		log.Printf("resize image to %d*%d" ,ykr.Cmd.AscWidth ,ykr.Cmd.AscHeight)
	}
	grey := image.NewGray(ykr.Img.Bounds())
	gh := NewGreyHandler()
	for i:=rtg.Min.X ;i<rtg.Max.X ;i++  {
		for j:=rtg.Min.Y ;j<rtg.Max.Y ;j++  {
			grey.SetGray(gh.GreyFunc(i ,j ,ykr.Img ,ykr.Cmd.Threshold))
		}
	}
	asc := NewAscii(grey)
	asc.Convert()
	w := NewImgWriter()
	if ykr.Cmd.TmpImgName != ""{
		log.Printf("output tmp image file [%s]" ,ykr.Cmd.TmpImgName)
		if err := w.writeImg(ykr.Cmd.TmpImgName ,grey);err != nil{
			log.Panic(err)
		}
	}
	if err := w.writeAscii(ykr.Cmd.Filename ,asc.AsciiMap);err != nil{
		log.Panic(err)
	}
	log.Printf("output is complete [%s]" ,ykr.Cmd.Filename)
}


