import parseopt

var
   o = initOptParser()
   pre_s, stem_s, suf_s: string

for m in o.getOpt():
   case m.key
   of "p":
      pre_s = m.val
   of "s":
      suf_s = m.val
   else:
      stem_s = m.key

if stem_s == "":
   echo """add [flags] <stem>
-p string
   prefix
-s string
   suffix"""
   quit 1

echo pre_s & stem_s & suf_s
