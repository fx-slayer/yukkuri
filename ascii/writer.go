package ascii

import (
	"bytes"
	"image"
	"image/jpeg"
	"os"
	"sync"
)

type ImgWriter struct {}

var imgWriter ImgWriter
var once sync.Once

func NewImgWriter() ImgWriter{
	once.Do(func() {
		imgWriter = ImgWriter{}
	})
	return imgWriter
}

type ImageType int

func (w ImgWriter)writeImg(filePath string ,img image.Image) error{
	if f ,err := os.OpenFile(filePath ,os.O_WRONLY | os.O_CREATE ,os.ModePerm);err != nil{
		return err
	}else{
		defer f.Close()

		return jpeg.Encode(f ,img ,&jpeg.Options{Quality:100})
	}
}

func (w ImgWriter)writeAscii(filepath string ,asc [][]string) error{
	if f ,err := os.OpenFile(filepath ,os.O_WRONLY | os.O_CREATE ,os.ModePerm);err != nil{
		return err
	}else{
		defer f.Close()

		buf := bytes.Buffer{}
		for i:=0 ;i<len(asc[0]) ;i++{
			for j:=0 ;j<len(asc) ;j++{
				b := []byte(asc[j][i])
				if _ ,err := buf.Write(b);err != nil{
					return err
				}
			}
			buf.Write([]byte("\r\n"))
		}
		_ ,err := f.Write(buf.Bytes())
		return err
	}
}
