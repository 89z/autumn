#!/bin/sh
set -u

awk '
{
   sub("\r", "")
   sub(".html", "")
   print
}
' /dev/clipboard |
while read each
do
   type -a "$each"
done 2>&1 |
grep -v found |
grep -v builtin |
grep -v Root |
awk '{print $1}'
