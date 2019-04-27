package ascii

import (
	"image"
)

type Asc11Converter interface {
	Convert()
	CharMap() [][]string
}

type Ascii struct {
	X int
	Y int
	GreyImg *image.Gray
	AsciiMap [][]string

	Ch0369 string

	Ch903 string
	Ch036 string
	Ch369 string
	Ch690 string

	Ch93 string
	Ch06 string
	Ch03 string
	Ch36 string
	Ch69 string
	Ch90 string
}

func NewAscii(grey *image.Gray) *Ascii{
	x := grey.Bounds().Max.X - grey.Bounds().Min.X
	y := grey.Bounds().Max.Y - grey.Bounds().Min.Y
	a := &Ascii{
		X:x,
		Y:y,
		GreyImg:grey,
		Ch0369:"┼",
		Ch903:"┴",
		Ch690:"┤",
		Ch036:"├",
		Ch369:"┬",
		Ch06:"│",
		Ch93:"─",
		Ch03:"└",
		Ch36:"┌",
		Ch69:"┐",
		Ch90:"┘",
	}
	var m [][]string
	for i:=0 ;i<x ;i++{
		mx := make([]string ,0)
		for j:=0 ;j<y ;j++{
			mx = append(mx, " ")
		}
		m = append(m ,mx)
	}
	a.AsciiMap = m
	return a
}

func (asc *Ascii)Convert(){
	rtg := asc.GreyImg.Bounds()

	// check image size
	if rtg.Max.X - rtg.Min.X < 2 || rtg.Max.Y - rtg.Min.Y < 2{
		return
	}

	// go through
	for i:=rtg.Min.X+1 ;i<rtg.Max.X-1 ;i++{
		for j:=rtg.Min.Y+1 ;j<rtg.Max.Y-1 ;j++{
			asc.setChar(i ,j)
		}
	}
}

func (asc *Ascii)CharMap() [][]string{
	return asc.AsciiMap
}

func (asc *Ascii)setChar(x int ,y int){
	img := asc.GreyImg
	if img.GrayAt(x ,y).Y == 0xff{
		return
	}
	var e ,s ,w ,n bool
	// east
	if img.GrayAt(x+1 ,y).Y == 0x00{
		e = true
	}
	// south
	if img.GrayAt(x ,y+1).Y == 0x00{
		s = true
	}
	// west
	if img.GrayAt(x-1 ,y).Y == 0x00{
		w = true
	}
	// north
	if img.GrayAt(x ,y-1).Y == 0x00{
		n = true
	}
	asc.AsciiMap[x][y] = asc.getChar(e ,s ,w ,n)
}


func (asc *Ascii)getChar(e ,s ,w ,n bool)string{
	// 0369
	if e && s && w && n{
		return asc.Ch0369
	}
	// xxx
	if e && s && n{
		return asc.Ch036
	}
	if e && s && w{
		return asc.Ch369
	}
	if n && s && w{
		return asc.Ch690
	}
	if n && e && w{
		return asc.Ch903
	}
	// xx
	if e && s{
		return asc.Ch36
	}
	if s && w{
		return asc.Ch69
	}
	if w && n{
		return asc.Ch90
	}
	if e && n{
		return asc.Ch03
	}
	if e && w{
		return asc.Ch93
	}
	if n && s{
		return asc.Ch06
	}
	return " "
}


