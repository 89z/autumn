---
title: 'Tag talk:PHP'
---

## Cygwin

This is valid code:

{{< r "a.php" >}}

but it can produce this error in some cases:

~~~
PHP Warning:  symlink(): No such file or directory
~~~

It will happen if you provide a Windows path to Cygwin PHP. Same with these:

{{< r "b.php" >}}
{{< r "c.php" >}}
{{< r "d.php" >}}

Workaround is use Cygwin path:

{{< r "e.php" >}}

## Native

Native PHP works best with Windows 10, because of ANSI escape codes.

New features:

<https://stitcher.io/blog/short-closures-in-php>

Download:

<https://windows.php.net/download>
