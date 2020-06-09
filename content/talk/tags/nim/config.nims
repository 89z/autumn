import os
task r, "compile to temp":
   setCommand("compile")
   switch("gcc.exe", "C:/llvm-mingw/bin/gcc.exe")
   switch("gcc.linkerexe", "C:/llvm-mingw/bin/gcc.exe")
   switch("outdir", getTempDir())
   switch("run")
