---
title: 'Talk:Console color'
---

Cygwin PHP is bad:

<https://sourceware.org/pipermail/cygwin/2020-May/244793.html>

so I am now using native PHP. Since PHP cannot do `SetConsoleTextAttribute`,
that means I need to use Mintty short term, and Windows 10 long term.

Since I am now using Mintty, I no longer have to worry about
`SetConsoleTextAttribute`, as Mintty will just accept ANSI escape codes. So I
can be more relaxed about what package I choose.

<https://github.com/mintty/mintty/issues/994>

## aec

2 imports

{{< r "aec.go" >}}

<https://pkg.go.dev/github.com/morikuni/aec>

## colorstring

5 imports

<https://pkg.go.dev/github.com/mitchellh/colorstring>

## line

6 imports

<https://pkg.go.dev/github.com/dollarshaveclub/line>
