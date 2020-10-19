import parseopt, strutils

var
   n_len = 1
   n_start = 0
   o_opt = initOptParser()
   s_in = ""

for n_kind, s_key, s_val in o_opt.getOpt():
   case o_opt.key
   of "a":
      n_start = s_val.parseInt
   of "b":
      n_len = s_val.parseInt
   else:
      s_in = s_key

if s_in == "":
   echo """slice [flags] <input>
-a int
   start
-b int
   length (default 1)"""
   quit(1)

echo s_in[n_start ..< n_start + n_len]
