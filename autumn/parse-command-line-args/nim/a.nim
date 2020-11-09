import parseopt

var
   s_pre = ""
   s_suf = ""
   s_stem = ""
   o_opt = initOptParser()

for n_kind, s_key, s_val in o_opt.getOpt():
   case o_opt.key
   of "p":
      s_pre = s_val
   of "s":
      s_suf = s_val
   else:
      s_stem = s_key

if s_stem == "":
   echo """add [flags] <stem>
-p string
   prefix
-s string
   suffix"""
   quit(1)

echo s_pre & s_stem & s_suf
