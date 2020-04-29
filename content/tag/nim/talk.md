---
title: Nim
---

## Setup

- <https://github.com/mstorsjo/llvm-mingw>
- <https://nim-lang.org/install_windows.html>

## Run

~~~nim
# %APPDATA%\nim\config.nims
import os
task run, "run":
   setCommand("compile")
   switch("run")
   switch("outdir", getTempDir())
~~~

- <https://nim-lang.github.io/Nim/nims.html>
- <https://nim-lang.org/docs/nimscript.html#nimcacheDir>
