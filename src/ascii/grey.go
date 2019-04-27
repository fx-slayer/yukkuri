package ascii

import (
	"fmt"
	"image"
	"image/color"
	"sort"
	"strconv"
	"strings"
)

type GreyHandler struct {
	defaultRatio map[GreyColor]float32
}

type GreyColor string

const (
	CRed  GreyColor = "red"
	CGreen GreyColor = "green"
	CBlue GreyColor = "blue"
	CYellow GreyColor = "yellow"
	CCyan GreyColor = "cyan"
	CMagenta GreyColor = "magenta"
)

func NewGreyHandler() GreyHandler{
	g := GreyHandler{}
	g.defaultRatio = map[GreyColor]float32{
		CRed:0.4,
		CGreen:0.4,
		CBlue:0.2,
		CYellow:0.6,
		CCyan:0.6,
		CMagenta:0.8,
	}
	return g
}

/*
	function:gray= (max - mid) * ratio_max + (mid - min) * ratio_max_mid + min.

	max,mid,min: point [max,mid,min] value in R.G.B.
	ratio_max: max color ratio.	ratio_max_mid: max and mid color ratio

	if threshold > -1 ,means grey == 0(<t) or (>t)255 ,decided by threshold

	return current x ,y and grey value
*/
func (gh GreyHandler)GreyFunc(x int ,y int ,img image.Image ,threshold int)(int, int, color.Gray){

	var maxRatio float32 = 0
	var maxMidRatio float32 = 0
	var max = 0x000
	var mid = 0x000
	var min = 0x000

	rgba := fmt.Sprint(img.At(x ,y))
	tmp := strings.Split(rgba[1:len(rgba) - 1] ," ")
	tr ,tg ,tb := tmp[0] ,tmp[1] ,tmp[2]
	r ,_ := strconv.Atoi(tr)
	g ,_ := strconv.Atoi(tg)
	b ,_ := strconv.Atoi(tb)
	rgb := []int{r ,g ,b}
	sort.Ints(rgb)

	switch rgb[2] {
	case r:
		max = r
		maxRatio = gh.defaultRatio[CRed]
	case g:
		max = g
		maxRatio = gh.defaultRatio[CGreen]
	case b:
		max = b
		maxRatio = gh.defaultRatio[CBlue]
	}

	switch rgb[1] {
	case r:
		mid = r
	case g:
		mid = g
	case b:
		mid = b
	}

	if b == rgb[0]{
		min = b
		maxMidRatio = gh.defaultRatio[CYellow]
	}
	if r == rgb[0]{
		min = r
		maxMidRatio = gh.defaultRatio[CCyan]
	}
	if g == rgb[0]{
		min = g
		maxMidRatio = gh.defaultRatio[CMagenta]
	}

	val := uint8((float32(max) - float32(mid)) * maxRatio + (float32(mid) - float32(min)) * maxMidRatio + float32(min))
	if threshold > -1{
		if int(val) < threshold{
			val = 0x00
		}else{
			val = 0xff
		}
	}

	grey := color.Gray{
		Y:val,
	}
	//log.Println(val)
	return x ,y ,grey
}


