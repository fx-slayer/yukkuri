package ascii

import (
	"image"
)

// 还原风格应当更趋近于油库里风，现在会默认将输入图像压缩3*3倍
// 输出更可能会变得不可辨认.
// 将会使用与Ascii不同的字符集.
type AsciiYukkuri struct {
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

func NewAsciiYukkuri(grey *image.Gray) *AsciiYukkuri{
	a := &AsciiYukkuri{}
	a.GreyImg = grey
	rtg := grey.Bounds()
	a.X = rtg.Max.X / 3
	a.Y = rtg.Max.Y / 3
	a.Ch03 = "ノ"
	a.Ch06 = "│"
	a.Ch36 = "ノ"
	a.Ch90 = "╰"
	a.Ch69 = "╰"
	a.Ch93 = "─"
	a.Ch036 = "ノ"
	a.Ch369 = "─"
	a.Ch690 = "╰"
	a.Ch903 = "─"
	a.Ch0369 = "."
	var m [][]string
	for i:=0 ;i<a.X ;i++{
		mx := make([]string ,0)
		for j:=0 ;j<a.Y ;j++{
			mx = append(mx, " ")
		}
		m = append(m ,mx)
	}
	a.AsciiMap = m
	return a
}


func (asc *AsciiYukkuri)Convert(){
	rtg := asc.GreyImg.Bounds()

	// check image size
	if rtg.Max.X - rtg.Min.X < 4 || rtg.Max.Y - rtg.Min.Y < 4{
		return
	}

	// go through
	for i:=0 ;i<asc.X ;i++{
		for j:=0 ;j<asc.Y ;j++{
			asc.setChar(i ,j)
		}
	}
}

func (asc *AsciiYukkuri)CharMap() [][]string{
	return asc.AsciiMap
}

func (asc *AsciiYukkuri)setChar(x int ,y int){
	img := asc.GreyImg
	if img.GrayAt(3*x ,3*y).Y == 0xff{
		return
	}
	var e ,s ,w ,n bool
	// east
	if img.GrayAt(3*x+1 ,3*y).Y == 0x00{
		e = true
	}
	// south
	if img.GrayAt(3*x ,3*y+1).Y == 0x00{
		s = true
	}
	// west
	if img.GrayAt(3*x-1 ,3*y).Y == 0x00{
		w = true
	}
	// north
	if img.GrayAt(3*x ,3*y-1).Y == 0x00{
		n = true
	}
	asc.AsciiMap[x][y] = asc.getChar(e ,s ,w ,n)
}


func (asc *AsciiYukkuri)getChar(e ,s ,w ,n bool)string{
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
