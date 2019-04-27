### yukkuri

将图片转为油库里(Ascii字符画)


```
go build main

main -h
  -f string
        specify an image file to transcode
  -h    help
  -t int
        set threshold for image grey processing (default 140)
```

执行`main -f test.jpg` 将在同一目录下生成`tmp.txt`和`grey_tmp.jpg`