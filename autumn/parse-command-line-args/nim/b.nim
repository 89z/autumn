import parseopt

var
   pre_s = ""
   suf_s = ""
   stem_s = ""
   opt_o = initOptParser()

while true:
   opt_o.next()
   if opt_o.kind == cmdEnd:
      break
   case opt_o.key
   of "p":
      pre_s = opt_o.val
   of "s":
      suf_s = opt_o.val
   else:
      stem_s = opt_o.key

if stem_s == "":
   echo """add [flags] <stem>
-p string
   prefix
-s string
   suffix"""
   quit(1)

echo pre_s & stem_s & suf_s
