---
title: Process
---

I need a process to use with examples. Here is the list:

<https://pubs.opengroup.org/onlinepubs/9699919799/utilities>

Select the correct column, then run this:

{{< r "a.sh" >}}

## Numeric argument

`cal`

## One argument problem

Either one argument is not possible, or it pulls from standard input:

- `awk`
- `cat`
- `chgrp`
- `chmod`
- `chown`
- `cksum`
- `cmp`
- `comm`
- `cp`
- `csplit`
- `cut`
- `dd`
- `diff`
- `dirname`
- `echo`
- `ex`
- `expand`
- `expr`
- `false`
- `file`
- `fold`
- `gencat`
- `getconf`
- `grep`
- `head`
- `iconv`
- `ipcrm`
- `ipcs`
- `join`
- `kill`
- `link`
- `ln`
- `logger`
- `man`
- `mkdir`
- `mkfifo`
- `more`
- `mv`
- `nl`
- `nohup`
- `od`
- `paste`
- `pathchk`
- `pr`
- `printf`
- `renice`
- `rm`
- `rmdir`
- `sed`
- `sh`
- `sleep`
- `sort`
- `split`
- `tabs`
- `tail`
- `tee`
- `test`
- `time`
- `touch`
- `tput`
- `tr`
- `true`
- `tsort`
- `unexpand`
- `uniq`
- `unlink`
- `vi`
- `wc`
- `who`
- `xargs`
- `zcat`

## Single line output

- `basename`
- `date`
- `du`
- `id`
- `logname`
- `ls`
- `nice`
- `pwd`
- `tty`
- `uname`

## Result

Doc size | Tool
---------|-----
14K      | `env`
16K      | `df`
20K      | `locale`
32K      | `ps`
38K      | `find`
44K      | `stty`

Out of these, I am liking `find`. Here are POSIX options:

~~~
-atime  n
-ctime  n
-depth
-exec  utility_name  [argument ...]   {} +
-exec  utility_name  [argument ...] ;
-group  gname
-links  n
-mtime  n
-name  pattern
-newer  file
-nogroup
-nouser
-ok  utility_name  [argument ...] ;
-path  pattern
-perm [-]mode
-perm [-]onum
-print
-prune
-size  n[c]
-type  c
-user  uname
-xdev
~~~
