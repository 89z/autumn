import parseopt

var
   s_pre = ""
   s_suf = ""
   s_stem = ""
   o_opt = initOptParser()

while true:
   o_opt.next()
   if o_opt.kind == cmdEnd:
      break
   case o_opt.key
   of "p":
      s_pre = o_opt.val
   of "s":
      s_suf = o_opt.val
   else:
      s_stem = o_opt.key

if s_stem == "":
   echo """add [flags] <stem>
-p string
   prefix
-s string
   suffix"""
   quit(1)

echo s_pre & s_stem & s_suf
