#!/bin/dash



ldc2 test.d C:/Path/lib/curl.lib wldap32.lib



rm test.obj


exit
ldc2 -L/OPT:NOICF test.d C:/Path/lib/curl.lib wldap32.lib






