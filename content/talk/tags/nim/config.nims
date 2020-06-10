switch("gcc.exe", "C:/llvm-mingw/bin/gcc.exe")
switch("gcc.linkerexe", "C:/llvm-mingw/bin/gcc.exe")

import os
task r, "compile to temp":
   setCommand("compile")
   switch("run")
   switch("outdir", getTempDir())
