import parseopt, strutils

var
   n_len = 1
   n_start = 0
   o_opt = initOptParser()
   s_in = ""

for n_kind, s_key, s_val in o_opt.getOpt():
   case o_opt.key
   of "start":
      n_start = s_val.parseInt
   of "len":
      n_len = s_val.parseInt
   else:
      s_in = s_key

if s_in == "":
   echo """slice [flags] <string>
--start int
   starting position
--len int
   substring length (default 1)"""
   quit(1)

echo s_in[n_start ..< n_start + n_len]
