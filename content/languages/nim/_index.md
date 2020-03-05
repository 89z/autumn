---
title: Nim
stars: 8038
---

## Install

<https://nim-lang.org/install_windows.html>

## Run

~~~nim
# %APPDATA%\nim\config.nims
import os
switch("gcc.exe", "x86_64-w64-mingw32-gcc")
switch("gcc.linkerexe", "x86_64-w64-mingw32-gcc")
task run, "run":
   setCommand("compile")
   switch("run")
   switch("outdir", getTempDir())
~~~

- <https://nim-lang.github.io/Nim/nims.html>
- <https://nim-lang.org/docs/nimscript.html#nimcacheDir>

## Stars

<https://github.com/nim-lang/Nim>
