### yukkuri

由Golang实现的一个小功能，可将图片转为油库里(Ascii字符画)

#### 依赖
[resize]: https://github.com/nfnt/resize






```
go build main

  -H int
        set ascii max height (default -1)
  -W int
        set ascii max width (default -1)
  -f string
        specify an image file to transcode
  -h    help
  -m string
        save temporary image
  -n string
        set output file name (default "tmp.txt")
  -t int
        set threshold for image grey processing (default 140)
```

