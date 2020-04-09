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

Out of these, I am liking `ps`, as we should be able to see the program itself
in the list, or interpreter. Here are POSIX options that can be used without an
argument:

~~~
-A
-a
-d
-e
-f
-l
~~~

Remove options not available with my version, result:

~~~
-a
-e
-f
-l
~~~

Now compare output:

~~~
$ ps
      PID    PPID    PGID     WINPID   TTY         UID    STIME COMMAND
     1374       1    1374       4648  cons0     197608 21:24:27 /usr/bin/sh
     1375    1374    1375       2204  cons0     197608 21:24:28 /usr/bin/ps

$ ps -a
      PID    PPID    PGID     WINPID   TTY         UID    STIME COMMAND
     1376    1374    1376       6304  cons0     197608 21:24:29 /usr/bin/ps
     1374       1    1374       4648  cons0     197608 21:24:27 /usr/bin/sh

$ ps -e
      PID    PPID    PGID     WINPID   TTY         UID    STIME COMMAND
     1377    1374    1377       2980  cons0     197608 21:24:31 /usr/bin/ps
     1374       1    1374       4648  cons0     197608 21:24:27 /usr/bin/sh

$ ps -f
     UID     PID    PPID  TTY        STIME COMMAND
  Steven    1374       1 cons0    21:24:27 /usr/bin/sh
  Steven    1378    1374 cons0    21:24:32 /usr/bin/ps

$ ps -l
      PID    PPID    PGID     WINPID   TTY         UID    STIME COMMAND
     1374       1    1374       4648  cons0     197608 21:24:27 /usr/bin/sh
     1379    1374    1379      11236  cons0     197608 21:24:35 /usr/bin/ps
~~~
