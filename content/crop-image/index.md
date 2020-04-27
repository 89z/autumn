---
title: Crop image
categories: [file-format]
tags: [go, php]
---

## Go

{{< r "a.go" >}}
{{< r "b.go" >}}

- <https://golang.org/pkg/image#YCbCr.SubImage>
- <https://golang.org/pkg/image/draw#Draw>

## PHP

{{< r "a.php" >}}

Warning, this requires `php-gd` package, which pulls in 23 other packages.

<https://php.net/function.imagecrop>

## Shell

~~~sh
$ cygcheck gm.exe | sed '/Found/d;/Windows/d;/90\./d;s/\.\\//' | xargs wc -c
2,389,504 total
~~~

<https://sourceforge.net/projects/graphicsmagick>

~~~sh
$ wc -c magick.exe
14,735,424 magick.exe
~~~

<https://imagemagick.org>

~~~sh
$ cygcheck vips.exe | sed '/Found/d;/Windows/d;s/\.\\//' | xargs wc -c
17,543,219 total
~~~

<https://github.com/libvips/libvips>

~~~sh
$ wc -c ffmpeg.exe
65,644,544 ffmpeg.exe
~~~

<https://ffmpeg.zeranoe.com>
