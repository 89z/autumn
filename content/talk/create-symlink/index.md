---
title: 'Talk:Create symlink'
---

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
