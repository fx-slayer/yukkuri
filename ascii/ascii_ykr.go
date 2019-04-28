package ascii

import (
	"image"
)

// 还原风格应当更趋近于油库里风，输入的图像大小应当小于50*50.
// 输出更可能会变得不可辨认.
// 将会使用与Ascii不同的字符集.
type AsciiYukkuri struct {
	*Ascii
}

func NewAsciiYukkuri(grey *image.Gray) *AsciiYukkuri{
	a := &AsciiYukkuri{
		NewAscii(grey),
	}
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
	return a
}
