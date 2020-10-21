import parseopt

var
   s_start = ""
   s_end = ""
   s_in = ""
   o_opt = initOptParser()

for n_kind, s_key, s_val in o_opt.getOpt():
   case o_opt.key
   of "s":
      s_start = s_val
   of "e":
      s_end = s_val
   else:
      s_in = s_key

if s_in == "":
   echo """cat [flags] <input>
-s string
   start
-e string
   end"""
   quit(1)

echo s_start & s_in & s_end
