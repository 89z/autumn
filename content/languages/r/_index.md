---
title: R
stars: 4027
---

## Examples

{{< r "a.r" >}}
{{< r "a2.r" >}}
{{< r "a3.r" >}}

## Issues

R_base
: <https://cygwin.com/ml/cygwin/2019-12/msg00117.html>

Setting LC_CTYPE=en_US.UTF-8 failed
: <https://hypatia.math.ethz.ch/pipermail/r-devel/2019-June/077955.html>

## Setup

Mingw:

~~~
innoextract R-*-win.exe
~~~

<https://cran.cnr.berkeley.edu/bin/windows/base>

Cygwin:

~~~
libgomp1
libicu65
liblapack0
libopenblas
libtirpc3
R
~~~

~~~
PATH=$PATH:/usr/lib/lapack
~~~

## Stars

<https://github.com/tidyverse/ggplot2>
