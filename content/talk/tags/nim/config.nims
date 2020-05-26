import os
task r, "compile to temp":
   setCommand("compile")
   switch("run")
   switch("outdir", getTempDir())
