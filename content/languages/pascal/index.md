---
title: Pascal language
categories: [language]
---

## Free Pascal

Old

<https://sourceforge.net/projects/freepascal/files/Win32>

## XD Pascal

First, binary files should be moved out of version control:

<https://github.com/vtereshkov/xdpw>

Second, with a file like this:

~~~
program Hello;
uses fphttpclient;
begin
   WriteLn( languagehello world!');
end.
~~~

I get this result:

~~~
$ xdpw a.pas
XD Pascal for Windows 0.11
Copyright (c) 2009-2010, 2019-2020, Vasiliy Tereshkov
Compiling system.pas
a.pas (3) Error: Unable to open source file FPHTTPCLIENT.pas
~~~

but this is confusing as that file doesnt exist:

~~~
units\i386-win32\fcl-web\fphttpclient.o
units\i386-win32\fcl-web\fphttpclient.ppu
units\i386-win32\fcl-web\fphttpclient.rsj
~~~

<https://sourceforge.net/projects/freepascal>

## Setup

- <https://freepascal.org/docs-html/current/user/usersu7.html>
- <https://wiki.lazarus.freepascal.org/Unit_not_found_-_How_to_find_units>
