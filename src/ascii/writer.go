package ascii

import (
	"bytes"
	"image"
	"image/jpeg"
	"os"
)

type ImgWriter struct {

}

func NewImgWriter() ImgWriter{
	return ImgWriter{}
}

type ImageType int

func (w ImgWriter)writeImg(filePath string ,img image.Image) error{
	if f ,err := os.OpenFile(filePath ,os.O_WRONLY | os.O_CREATE ,os.ModePerm);err != nil{
		return err
	}else{
		return jpeg.Encode(f ,img ,&jpeg.Options{Quality:100})
	}
}

func (w ImgWriter)writeAscii(filepath string ,asc [][]string) error{
	if f ,err := os.OpenFile(filepath ,os.O_WRONLY | os.O_CREATE | os.O_APPEND,os.ModePerm);err != nil{
		return err
	}else{
		buf := bytes.Buffer{}
		for i:=0 ;i<len(asc[0]) ;i++{
			for j:=0 ;j<len(asc) ;j++{
				b := []byte(asc[j][i])
				if _ ,err := buf.Write(b);err != nil{
					return err
				}
				//if _ ,err := f.Write(b);err != nil{
				//	return err
				//}
			}
			//f.Write([]byte("\r\n"))
			buf.Write([]byte("\r\n"))
		}
		_ ,err := f.Write(buf.Bytes())
		return err
	}
}
