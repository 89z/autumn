import parseopt

var
   opt = initOptParser()
   pre, suf, stem: string

while true:
   opt.next()
   if opt.kind == cmdEnd:
      break
   case opt.key
   of "p":
      pre = opt.val
   of "s":
      suf = opt.val
   else:
      stem = opt.key

if stem == "":
   echo """add [flags] <stem>
-p string
   prefix
-s string
   suffix"""
   quit 1

echo pre & stem & suf
