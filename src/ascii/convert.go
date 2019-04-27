package ascii

import (
	"errors"
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
	Threshold int
}

func TransFile(path string ,threshold int) error{
	ykr := &Yukurri{
		Threshold:threshold,
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
	log.Println(ykr.File.Name() ,ykr.Img.Bounds())
	grey := image.NewGray(ykr.Img.Bounds())
	gh := NewGreyHandler()
	for i:=rtg.Min.X ;i<rtg.Max.X ;i++  {
		for j:=rtg.Min.Y ;j<rtg.Max.Y ;j++  {
			grey.SetGray(gh.GreyFunc(i ,j ,ykr.Img ,ykr.Threshold))
		}
	}
	asc := NewAscii(grey)
	asc.Convert()
	w := NewImgWriter()
	if err := w.writeImg("grey_tmp.jpg" ,grey);err != nil{
		log.Panic(err)
	}
	if err := w.writeAscii("tmp.txt" ,asc.AsciiMap);err != nil{
		log.Panic(err)
	}
}


