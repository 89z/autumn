import parseopt

var
   o = initOptParser()
   pre_s, suf_s, stem_s: string

while true:
   o.next()
   if o.kind == cmdEnd:
      break
   case o.key
   of "p":
      pre_s = o.val
   of "s":
      suf_s = o.val
   else:
      stem_s = o.key

if stem_s == "":
   echo """add [flags] <stem>
-p string
   prefix
-s string
   suffix"""
   quit 1

echo pre_s & stem_s & suf_s
