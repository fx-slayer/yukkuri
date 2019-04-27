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

type Yukkuri struct {
	Img 	image.Image
	GrayImg *image.Gray
	File	*os.File
	Greyer	GreyFunction
	Cmd 	command.Cmd
}

func NewYukkuri(cmd command.Cmd) *Yukkuri{
	ykr := &Yukkuri{
		Cmd:cmd,
		Greyer:NewGreyHandler(),
	}
	return ykr
}

func (ykr *Yukkuri)TransImgToGrey() error{
	if f ,err := os.OpenFile(ykr.Cmd.ImgPath ,os.O_RDONLY ,os.ModeTemporary);err != nil{
		return err
	}else{
		defer f.Close()

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
		//ykr.HandleImage()
		rtg := ykr.Img.Bounds()
		log.Printf("image size %d*%d" ,ykr.Img.Bounds().Max.X ,ykr.Img.Bounds().Max.Y)
		// resize
		if ykr.Cmd.AscWidth > 0 && ykr.Cmd.AscHeight > 0{
			ykr.Img = resize.Resize(uint(ykr.Cmd.AscWidth) ,uint(ykr.Cmd.AscHeight) ,ykr.Img ,resize.Lanczos3)
			log.Printf("resize image to %d*%d" ,ykr.Cmd.AscWidth ,ykr.Cmd.AscHeight)
		}
		// gray
		grey := image.NewGray(ykr.Img.Bounds())
		for i:=rtg.Min.X ;i<rtg.Max.X ;i++  {
			for j:=rtg.Min.Y ;j<rtg.Max.Y ;j++  {
				grey.SetGray(ykr.Greyer.GreyFunc(i ,j ,ykr.Img ,ykr.Cmd.Threshold))
			}
		}
		ykr.GrayImg = grey
		return nil
	}
}

func (ykr *Yukkuri)TransImgToAsc(converter Asc11Converter){

	converter.Convert()

	w := NewImgWriter()
	if ykr.Cmd.TmpImgName != ""{
		log.Printf("output tmp image file [%s]" ,ykr.Cmd.TmpImgName)
		if err := w.writeImg(ykr.Cmd.TmpImgName ,ykr.GrayImg);err != nil{
			log.Panic(err)
		}
	}
	if err := w.writeAscii(ykr.Cmd.Filename ,converter.CharMap());err != nil{
		log.Panic(err)
	}
	log.Printf("output is complete [%s]" ,ykr.Cmd.Filename)
}


