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

func (ykr *Yukkuri)TransImgToGrey() (*image.Gray ,error){
	if f ,err := os.OpenFile(ykr.Cmd.ImgPath ,os.O_RDONLY ,os.ModeTemporary);err != nil{
		return nil ,err
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
			return nil ,errors.New("unsupported image type")
		}
		if decodeErr != nil{
			return nil ,decodeErr
		}

		rtg := ykr.Img.Bounds()
		log.Printf("image size %d*%d" ,ykr.Img.Bounds().Max.X ,ykr.Img.Bounds().Max.Y)

		// resize
		if ykr.Cmd.AscWidth > 0 && ykr.Cmd.AscHeight > 0{
			ykr.Img = resize.Resize(uint(ykr.Cmd.AscWidth) ,uint(ykr.Cmd.AscHeight) ,ykr.Img ,resize.Lanczos3)
			log.Printf("resize image to %d*%d" ,ykr.Cmd.AscWidth ,ykr.Cmd.AscHeight)
		}

		// gray handle
		grey := image.NewGray(ykr.Img.Bounds())
		for i:=rtg.Min.X ;i<rtg.Max.X ;i++  {
			for j:=rtg.Min.Y ;j<rtg.Max.Y ;j++  {
				grey.SetGray(ykr.Greyer.GreyFunc(i ,j ,ykr.Img ,ykr.Cmd.Threshold))
			}
		}

		// write tmp gray img
		if ykr.Cmd.TmpImgName != ""{
			w := NewImgWriter()
			log.Printf("output tmp image file [%s]" ,ykr.Cmd.TmpImgName)
			if err := w.writeImg(ykr.Cmd.TmpImgName ,grey);err != nil{
				log.Panic(err)
			}
		}

		return grey ,nil
	}
}

func (ykr *Yukkuri)TransImgToAsc(converter Asc11Converter){

	converter.Convert()

	w := NewImgWriter()
	if err := w.writeAscii(ykr.Cmd.Filename ,converter.CharMap());err != nil{
		log.Panic(err)
	}
	log.Printf("output is complete [%s]" ,ykr.Cmd.Filename)
}


