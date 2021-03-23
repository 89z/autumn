import parseopt

var
   opt = initOptParser()
   pre, stem, suf: string

for m in opt.getOpt():
   case m.key
   of "p":
      pre = m.val
   of "s":
      suf = m.val
   else:
      stem = m.key

if stem == "":
   echo """add [flags] <stem>
-p string
   prefix
-s string
   suffix"""
   quit 1

echo pre & stem & suf
